package main

import (
	"flag"
	"fmt"
	"github.com/funkygao/fae/servant/gen-go/fun/rpc"
	"github.com/funkygao/fae/servant/proxy"
)

const (
	REASON = "pingfae"
	RID    = "1" // request id
)

var (
	host string
	port string
)

func init() {
	flag.StringVar(&host, "host", "", "host name of faed")
	flag.StringVar(&port, "port", "", "fae port")
	flag.Parse()
}

func main() {
	proxy := proxy.NewWithDefaultConfig()
	if host == "" {
		pingCluster(proxy)
		return
	}

	// ping a single faed
	client, err := proxy.Servant(host + ":" + port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Recycle()

	ctx := rpc.NewContext()
	ctx.Reason = REASON
	ctx.Rid = RID
	pong, err := client.Ping(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pong)
	}

}

func pingCluster(proxy *proxy.Proxy) {
	go proxy.StartMonitorCluster()
	this.client.AwaitClusterTopologyReady()

	for peerAddr := range proxy.ClusterPeers() {
		client, err := proxy.Servant(peerAddr)
		if err != nil {
			fmt.Printf("%s: %s\n", peerAddr, err)
			continue
		}

		ctx := rpc.NewContext()
		ctx.Reason = REASON
		ctx.Rid = RID
		pong, err := client.Ping(ctx)
		if err != nil {
			fmt.Printf("%16s: %s\n", peerAddr, err)
		} else {
			fmt.Printf("%16s: %s\n", peerAddr, pong)
		}

		client.Recyle()
	}

}
