package distributed_lock

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/client"
)

type EtcdMutex struct {
	Key  string
	Val  string
	Ttl  time.Duration
	KApi client.KeysAPI
}

// PrevExist 	更新请求
// PrevNoExist	创建请求
func (em EtcdMutex) Lock() error {
	opt := &client.SetOptions{PrevExist: client.PrevExist, PrevValue: em.Val, TTL: em.Ttl}
	_, err := em.KApi.Set(context.TODO(), em.Key, em.Val, opt)
	return err
}

func (em EtcdMutex) TryLock() error {
	opt := &client.SetOptions{PrevExist: client.PrevNoExist, TTL: em.Ttl}
	_, err := em.KApi.Set(context.TODO(), em.Key, em.Val, opt)
	return err
}

func (em EtcdMutex) UnLock() error {
	opt := &client.SetOptions{PrevExist: client.PrevExist, PrevValue: em.Val, TTL: time.Nanosecond}
	_, err := em.KApi.Set(context.TODO(), em.Key, em.Val, opt)
	return err
}

type KeysAPI struct {
	Addr []string
	Api  client.KeysAPI
}

func NewEtcdClient(addrList []string) (error, *KeysAPI) {
	if len(addrList) == 0 {
		return fmt.Errorf("addrList empty"), nil
	}

	cfg := client.Config{
		Endpoints:               addrList,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: TTL,
	}

	c, err := client.New(cfg)
	if err != nil {
		fmt.Printf("new etcd client [cfg:%+v] err: %+v", cfg, err)
		return err, nil
	}

	kAPi := &KeysAPI{
		Addr: addrList,
		Api:  client.NewKeysAPI(c),
	}
	return nil, kAPi
}
