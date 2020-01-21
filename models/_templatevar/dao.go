package templatevar

const (
	DaoTemplate = `
package dao

import (
    "github.com/jinzhu/gorm"
    "golang.org/x/sync/singleflight"
    "gopkg.in/redis.v5"

    _ "github.com/jinzhu/gorm/dialects/postgres"

    "{{.Git}}/{{.ProjectName}}/models"
    "{{.Git}}/util/icache"
    "{{.Git}}/util/idb"
)

type Dao struct {
    db       *gorm.DB
    redisCli *redis.Client
    sf       *singleflight.Group
}

func New(cfg *models.Config) (*Dao, error) {

    return &Dao{
        db:       idb.New(cfg.Db),
        redisCli: icache.NewClient(cfg.Redis),
        sf:       &singleflight.Group{},
    }, nil
}

func (d *Dao) Stop() {
    if d.db != nil {
        _ = d.db.Close()
    }
}

`
)
