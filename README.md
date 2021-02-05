# distributedid

基于snowflake的分布式唯一id。

说明： 需要传入ip用于生成节点的唯一标识。

适用场景：
- 分布式服务中同一个服务每个节点所在ip的后两位不能完全相同。常用的k8s集群在规划网段的时候基本都会满足这个网络要求

改进：
- 对于基础设施能直接区分物理机和容器ip的场景下，可以修改ip的获取方式：不传ip。


## 待支持的功能

- [x] 生成分布式唯一id
- [ ] mysql的自增id转换分布式id工具
- [ ] mongodb的自增id转换分布式id工具
