package heka_redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/mozilla-services/heka/pipeline"
)

type RedisOutputConfig struct {
	Address string `toml:"address"`
	Key string `toml:"key"`
}

type RedisOutput struct {
	conf *RedisOutputConfig
	conn redis.Conn
}

func (rop *RedisOutput) ConfigStruct() interface{} {
	return &RedisOutputConfig{
		Address: "127.0.0.1:6379",
		Key: "heka",
	}
}

func (rop *RedisOutput) Init(config interface{}) error {
	rop.conf = config.(*RedisOutputConfig)
	var err error
	rop.conn, err = redis.Dial("tcp", rop.conf.Address)
	if err != nil {
		return fmt.Errorf("connecting to - %s", err.Error())
	}
	return nil
}

func (rop *RedisOutput) Run(or pipeline.OutputRunner, h pipeline.PluginHelper) error {
	inChan := or.InChan()
	for pack := range inChan {
		payload := pack.Message.GetPayload()
		_,err := rop.conn.Do("LPUSH", rop.conf.Key, payload)
		if err != nil {
			or.LogError(err)
			continue
		}
		pack.Recycle()
	}
	return nil
}

func (rop *RedisOutput) Stop() {
	rop.conn.Close()
}

func init() {
	pipeline.RegisterPlugin("RedisOutput", func() interface{} {
		return new(RedisOutput)
	})
}
