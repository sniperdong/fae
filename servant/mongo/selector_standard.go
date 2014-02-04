package mongo

import (
	"fmt"
	"github.com/funkygao/fae/config"
	"strings"
)

type StandardServerSelector struct {
	ShardBaseNum int
	Servers      map[string]*config.ConfigMongodbServer // key is shardName
}

func NewStandardServerSelector(baseNum int) *StandardServerSelector {
	return &StandardServerSelector{ShardBaseNum: baseNum}
}

func (this *StandardServerSelector) PickServer(shardName string,
	shardId int) (string, error) {
	var bucket string
	if !strings.HasPrefix(shardName, SHARD_DB_PREFIX) {
		bucket = shardName
	} else {
		bucket = fmt.Sprintf("db%s", (shardId/this.ShardBaseNum)+1)
	}

	if server, present := this.Servers[bucket]; present {
		return server.Address(), nil
	}

	return "", ErrServerNotFound
}

func (this *StandardServerSelector) SetServers(servers map[string]*config.ConfigMongodbServer) {
	this.Servers = servers
}