# kafka常用命令
```sh
查询topic
kafka-topics.sh --list --zookeeper 192.168.1.130:2181

创建topic
kafka-topics.sh --bootstrap-server 192.168.1.130:9092 --create --topic audit_log --partitions 4 --replication-factor 1

查询topic详情
kafka-topics.sh --bootstrap-server 192.168.1.130:9092 --describe --topic audit_log

修改topic分区
kafka-topics.sh --bootstrap-server 192.168.1.130:9092 --alter --topic audit_log --partitions 2

插入数据到topic中
kafka-console-producer.sh --bootstrap-server 192.168.1.130:9092 --topic audit_log

查询topic内容
kafka-console-consumer.sh --bootstrap-server 192.168.1.130:9092 --topic audit_log --from-beginning

删除topic
kafka-topics.sh  kafka-console-consumer.sh --bootstrap-server 192.168.1.130:9092 --delete --topic audit_log
```
