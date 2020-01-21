package templatevar

const (
	ModelsConfigGoTemplate = `
package models

import (
	"encoding/json"

	"{{.Git}}/util/idb"
	"{{.Git}}/util/instrumenting"
	"{{.Git}}/util/instrumenting/itracing"

	"{{.Git}}/util/icache"
	"{{.Git}}/util/iregister"
)

type Config struct {
	Addr            string
	Etcd            *iregister.EtcdConfig
	Redis           *icache.Options
	Instrumentation *instrumenting.Option
	Db          	*idb.DBConf
}

func (c *Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return "\n" + string(b)
}

var DefaultConfig = &Config{
	Addr: ":9001",
	Etcd: &iregister.EtcdConfig{
		Addresses: []string{":2379"},
	},
	Redis: &icache.Options{
		Address: nil,
		Host:    ":6379",
	},
	Instrumentation: &instrumenting.Option{
		Tracing: &itracing.Options{
			SamplerType:      "const",
			SamplerParam:     1,
			ReporterLogSpans: true,
		},
	},
	Db: &idb.DBConf{
		DbSource: "root:123456@tcp(db.dev:5432)/{{.ProjectName}}?charset=utf8mb4&collation=utf8mb4_unicode_ci",
		LogMode:  true,
	},
}

`
)
