## 服务发现

解决分布式集群中如何找到对方。而服务发现需要解决三大问题：

* 强一致性 高可用的 服务存储目录
* 注册健康检测
* 提供目录查询服务

下面通过ETCD的watcher机制来实现简单的[服务发现](https://github.com/Fedomn/go-knowledge/blob/master/test/etcd/service_discovery/master.go)

主要逻辑 master通过sync.RWMutex和watch监听set/update和expire/delete 来进行添加和删除节点。worker通过循环注册来告诉master 自己的健康状态。