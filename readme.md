## mysql-statsd

收集 MySQL `SHOW GLOBAL STATUS` 中获取到的 metrics, 发送给 StatsD

只收集数字类型的，对`Innodb_buffer_pool_dump_status` 等 string 格式的将自动过滤

## 配置文件

见 `conf_sample.yaml`

* `mysql`: mysql 连接地址
* `statsd`: StatsD 连接地址
* `interval`: 收集数据的时间间隔，单位 s
* `metrics`: 所有需要收集的 metrics 名称， 准确名称见 `show global status`

## 运行 

`go run main.go -c /path-to-file/conf.yaml`
