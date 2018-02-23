package distributed_lock

import (
	"fmt"
	"testing"
	"time"
)

type TestServer struct {
	Name    string
	Id      string
	Runaway bool
	Status  string
}

func (s TestServer) OnMaster(msg string) error {
	fmt.Printf("OnMaster i'm master server\n")
	return nil
}
func (s TestServer) OnSlave(msg string) error {
	fmt.Printf("OnSlave i'm salve server\n")
	return nil
}
func (s TestServer) ServerName() string {
	return s.Name
}
func (s TestServer) GroupId() string {
	return s.Id
}
func (s TestServer) IsRunaway() (bool, string) {
	return s.Runaway, "i'm alive"
}

var etcdAddr = []string{"http://127.0.0.1:2379"}

func TestHaServer1(t *testing.T) {
	var server1 = &TestServer{"TestServer", "group_1", false, "slave"}
	if err := RunAsHAServer(server1, etcdAddr); err != nil {
		fmt.Printf("start server fail err:%+v", err)
	}
	time.Sleep(30 * time.Second)
	server1.Runaway = true
	select {}
}

func TestHaSever2(t *testing.T) {
	var server2 = &TestServer{"TestServer", "group_1", false, "slave"}
	if err := RunAsHAServer(server2, etcdAddr); err != nil {
		fmt.Printf("start server fail err:%+v", err)
	}
	time.Sleep(40 * time.Second)
	server2.Runaway = true
	select {}
}
