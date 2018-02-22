package service_discovery

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var EtcdAddrList = []string{"http://127.0.0.1:2379"}
var Service = "test"

func TestMaster(t *testing.T) {
	master, err := NewMaster(Service, EtcdAddrList)
	if err != nil {
		log.Fatal(err)
	}
	master.Fetch()
	go master.Watch()
	for {
		log.Println("nodes: ", master.GetNodes())
		time.Sleep(time.Second * 2)
	}
}

func TestWorker(t *testing.T) {
	w1, err := NewWorker(Service, EtcdAddrList, "worker1", "i am worker1")
	if err != nil {
		log.Fatal(fmt.Sprintf("new worker 1 err: %+v", err))
	}
	w1.Register()

	go func() {
		time.Sleep(time.Second * 10)
		w1.UnRegister()
	}()

	for {
		log.Println("w1 isActive: ", w1.IsActive())
		log.Println("w1 isStop: ", w1.IsStop())
		time.Sleep(time.Second * 2)
		if w1.stop {
			return
		}
	}
}
