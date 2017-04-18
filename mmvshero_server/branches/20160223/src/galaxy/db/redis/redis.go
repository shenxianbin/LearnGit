package redis

import (
	"github.com/mediocregopher/radix.v2/cluster"
	"github.com/mediocregopher/radix.v2/redis"
)

type GxRedis struct {
	cluster *cluster.Cluster
	host    string
}

func NewGxRedis() *GxRedis {
	return &GxRedis{}
}

func (this *GxRedis) Cmd(cmd string, args ...interface{}) (*redis.Resp, error) {
	resp := this.cluster.Cmd(cmd, args)
	if resp.Err != nil {
		return nil, resp.Err
	}

	return resp, nil
}

func (this *GxRedis) Run(config string) error {
	var err error
	this.host = config
	this.cluster, err = cluster.New(this.host)
	if err != nil {
		return err
	}
	return nil
}
