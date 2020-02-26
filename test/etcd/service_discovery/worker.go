package service_discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/client"
)

var TTL = time.Second * 3
var HeartbeatInterval = time.Second * 2

type Worker struct {
	kapi   client.KeysAPI
	key    string
	info   string
	active bool
	stop   bool
}

func NewWorker(service string, addrList []string, node string, info string) (*Worker, error) {
	cfg := client.Config{
		Endpoints:               addrList,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second * 3,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	worker := &Worker{
		kapi:   client.NewKeysAPI(c),
		key:    fmt.Sprintf("%s/%s/%s", KEY, service, node),
		info:   info,
		active: false,
		stop:   false,
	}
	return worker, nil
}

func (w *Worker) heartbeat() error {
	_, err := w.kapi.Set(context.Background(), w.key, w.info, &client.SetOptions{
		TTL: TTL,
	})
	w.active = err == nil
	return err
}

func (w *Worker) heartbeatPeriod() {
	for !w.stop {
		w.heartbeat()
		time.Sleep(HeartbeatInterval)
	}
}

func (w *Worker) Register() {
	go w.heartbeatPeriod()
}

func (w *Worker) UnRegister() {
	w.stop = true
}

func (w *Worker) IsActive() bool {
	return w.active
}

func (w *Worker) IsStop() bool {
	return w.stop
}
