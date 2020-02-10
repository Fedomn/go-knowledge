## 分布式锁

ETCD提供一套实现CAS(CompareAndSwap)的API，通过设置prevExist来保证多节点同时创建，只有一个节点成功，即获得了锁。可以理解PrevExist:更新请求，PrevNoExist:创建请求。

CAS语义：我认为Val的值应该是1，如果是，就将Val的值更新为2，否则不修改并返回Val的实际值。

所以当多个线程使用CAS同时修改一个变量时，只能一个线程修改成功，其他的都会失败，但是可以再次尝试。

下面通过分布式锁来实现服务的高可用，即存在一个master与多个slave。抢占到锁的为master，锁都有TTL来保证健康。

而slave会每隔一段时间 去尝试获取锁。若获取成功，把自己状态改成master。

**这种情况适合，只能跑一个服务，但是又要保证该服务的高可用。如[定时服务](https://github.com/Fedomn/go-knowledge/blob/master/test/etcd/distributed_lock/ha.go)**