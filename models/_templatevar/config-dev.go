package templatevar

const (
	ConfigDev = `
addr = ":9001"

# redis 需要哨兵机制 + 集群
[redis]
    host = "redis.dev:6379"

[etcd]
    address = [
    "etcd.dev:2379"
    ]
    dial_timeout = 3
    dial_keep_alive = 3
    heartbeat = 3
    ttl = 10

[instrumentation]
[instrumentation.tracing]
    sampler_type = 1
    sampler_param = 1
    reporter_log_spans = true

[db]
    db_source = "root:123456@tcp(db.dev:5432)/{{.ProjectName}}?charset=utf8mb4&collation=utf8mb4_unicode_ci"
    log_mode = true
`
)
