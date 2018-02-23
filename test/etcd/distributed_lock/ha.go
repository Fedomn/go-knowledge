package distributed_lock

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"
)

type Node struct {
	ServerName string
	GroupId    string
}

type NodeStatus string

const (
	Unknown NodeStatus = "unknown"
	Master  NodeStatus = "master"
	Slave   NodeStatus = "slave"
)

type Server interface {
	OnMaster(msg string) error
	OnSlave(msg string) error
	ServerName() string
	GroupId() string
	IsRunaway() (bool, string)
}

type HAStat struct {
	AsMasterCheckFailedCnt int64
	AsSlaveCheckFailedCnt  int64
	SwitchToMasterCnt      int64
	SwitchToSlaveCnt       int64
}

type HAServer struct {
	node       Node
	server     Server
	ServerId   string
	NodeStatus NodeStatus
	HAStat     HAStat
	EtcdMutex  EtcdMutex
}

func (hs *HAServer) String() string {
	return fmt.Sprintf("%s:%s ServerId: %s Status: %s", hs.node.ServerName, hs.node.GroupId, hs.ServerId, hs.NodeStatus)
}

func (hs *HAServer) getKey() string {
	env := os.Getenv("GOENV")
	if env == "" {
		env = "online"
	} else {
		env = strings.ToLower(env)
	}
	return fmt.Sprintf("/lock/%s/%s/%s", env, hs.node.ServerName, hs.node.GroupId)
}

func (hs *HAServer) switchToMaster(msg string) {
	log.Printf("switch to master msg: %+v, %+v", msg, hs)
	if err := hs.server.OnMaster(msg); err != nil {
		log.Printf("switch to master err: %+v, %+v", err, hs)
		return
	}
	hs.HAStat.SwitchToMasterCnt++
	hs.NodeStatus = Master
}

func (hs *HAServer) switchToSlave(msg string) {
	log.Printf("switch to slave msg: %s, %+v", msg, hs)
	if err := hs.server.OnSlave(msg); err != nil {
		log.Printf("switch to slave err: %+v, %+v", err, hs)
	}
	hs.HAStat.SwitchToSlaveCnt++
	hs.NodeStatus = Slave
	return
}

func (hs *HAServer) runaway() (bool, string) {
	return hs.server.IsRunaway()
}

var TTL = 10 * time.Second

const (
	MsgNoMaster       = "no master server"
	MsgMasterCheckErr = "master check err"
)

var slaveCheckReg = regexp.MustCompile("Key already exists")

func (hs *HAServer) run() {
	log.Printf("start run HAServer: %+v", hs)

	interval := TTL / 3
	if interval < time.Second {
		interval = 3 * time.Second
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	sig := listenSignal()
	masterLockFailed := 0
	for {
		select {
		case <-ticker.C:
			if hs.NodeStatus == Master {
				if err := hs.EtcdMutex.Lock(); err != nil {
					masterLockFailed++
					log.Printf("server check failed, failed_cnt:%v, err:%+v, %+v", masterLockFailed, err, hs)
					if masterLockFailed == 3 {
						hs.EtcdMutex.UnLock()
						hs.switchToSlave(fmt.Sprintf("%v:%v", MsgMasterCheckErr, err))
						masterLockFailed = 0
					}
					hs.HAStat.AsMasterCheckFailedCnt++
				} else {
					if isRunaway, msg := hs.server.IsRunaway(); isRunaway {
						log.Printf("server runaway msg: %s, %+v", msg, hs)
						hs.switchToSlave(msg)
						hs.EtcdMutex.UnLock()
						masterLockFailed = 0
						time.Sleep(TTL)
					}
					hs.HAStat.AsMasterCheckFailedCnt = 0
					log.Printf("server check success, %+v", hs)
				}
			} else {
				if err := hs.EtcdMutex.TryLock(); err == nil {
					hs.switchToMaster(MsgNoMaster)
				} else {
					if !slaveCheckReg.MatchString(err.Error()) {
						hs.HAStat.AsSlaveCheckFailedCnt++
						log.Printf("server check failed, err:%+v, %+v", err, hs)
					} else {
						hs.HAStat.AsSlaveCheckFailedCnt = 0
						log.Printf("server check success, %+v", hs)
					}
				}
			}
		case s := <-sig:
			msg := fmt.Sprintf("server get signal:%v, exit", s.String())
			log.Printf(msg)
			hs.switchToSlave(msg)
			os.Exit(0)
		}
	}
	log.Printf("stop HAServer: %+v", hs)
}

func RunAsHAServer(server Server, etcd []string) error {
	var err error

	node := Node{ServerName: server.ServerName(), GroupId: server.GroupId()}
	if len(node.ServerName) == 0 || len(node.GroupId) == 0 {
		err = fmt.Errorf("node info err: %+v", node)
	}
	if server == nil {
		err = fmt.Errorf("server is nil")
	}

	var keysAPI *KeysAPI
	err, keysAPI = NewEtcdClient(etcd)

	if err != nil {
		log.Printf("RunAsHAServer Failed: %+v", err)
		return err
	}

	ha := HAServer{node: node, server: server, ServerId: serverId(), NodeStatus: Slave}
	ha.EtcdMutex = EtcdMutex{Key: ha.getKey(), Val: ha.ServerId, Ttl: TTL, KApi: keysAPI.Api}

	go ha.run()
	return nil
}

func serverId() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func listenSignal(signals ...os.Signal) <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	if len(signals) == 0 {
		signals = append(signals, os.Kill, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR2)
	}
	signal.Notify(sig, signals...)
	return sig
}
