package etcdv3

import (
	"context"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/karlma/soulsbag/source"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 5 * time.Second
)

type Etcdv3 struct {
	Endpoints []string
	Key       string
	User      string
	Password  string
}

func (e Etcdv3) Read() ([]byte, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   e.Endpoints,
		DialTimeout: dialTimeout,
		Username:    e.User,
		Password:    e.Password,
	})
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, e.Key)
	cancel()
	if err != nil {
		return nil, err
	}

	for _, ev := range resp.Kvs {
		return ev.Value, nil
	}

	return nil, nil
}

func (e Etcdv3) String() string {
	return "etcdv3"
}

func NewSource(opts source.Options) (source.Source, error) {

	return Etcdv3{
		Endpoints: strings.Split(opts.Path, ","),
		Key:       opts.Key,
		User:      opts.AuthUser,
		Password:  opts.AuthPassword,
	}, nil
}

// import this module from soulsbag
func init() {
	source.Register("etcdv3", NewSource)
}
