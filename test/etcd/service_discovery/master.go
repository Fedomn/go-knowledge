package service_discovery

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/client"
	"sync"
	"time"
)

const KEY = "/service/discovery"

type Master struct {
	sync.RWMutex
	kapi   client.KeysAPI
	key    string
	nodes  map[string]string
	active bool
}

func NewMaster(service string, addrList []string) (*Master, error) {
	cfg := client.Config{
		Endpoints:               addrList,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	master := &Master{
		kapi:   client.NewKeysAPI(c),
		key:    fmt.Sprintf("%s/%s", KEY, service),
		nodes:  make(map[string]string),
		active: true,
	}
	return master, err
}

func (m *Master) GetNodes() map[string]string {
	m.RLock()
	defer m.RUnlock()
	return m.nodes
}

func (m *Master) AddNodes(node, info string) {
	m.Lock()
	defer m.Unlock()
	m.nodes[node] = info
}

func (m *Master) DelNode(node string) {
	m.Lock()
	defer m.Unlock()
	delete(m.nodes, node)
}

func (m *Master) Fetch() error {
	rsp, err := m.kapi.Get(context.Background(), m.key, nil)
	if err != nil {
		return err
	}
	if rsp.Node.Dir {
		for _, v := range rsp.Node.Nodes {
			m.AddNodes(v.Key, v.Value)
		}
	}
	return nil
}

func (m *Master) Watch() {
	watcher := m.kapi.Watcher(m.key, &client.WatcherOptions{Recursive: true})
	for {
		rsp, err := watcher.Next(context.Background())
		if err != nil {
			fmt.Printf("watcher next err: %+v\n", err)
			m.active = false
			continue
		}
		m.active = true
		switch rsp.Action {
		case "set", "update":
			m.AddNodes(rsp.Node.Key, rsp.Node.Value)
		case "expire", "delete":
			m.DelNode(rsp.Node.Key)
		default:
			fmt.Printf("watch something: %+v", rsp)
		}
	}
}
