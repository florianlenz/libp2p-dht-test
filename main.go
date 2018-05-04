package main

import (
	"context"
	dht "gx/ipfs/QmSBxn1eLMdViZRDGW9rRHRYwtqq5bqUgipqTMPuTim616/go-libp2p-kad-dht"
	bootstrap "gx/ipfs/QmVRQBf4hnofDzDZ7oFKSb8GchwVBK2ojuZw1Biwbxvget/go-libp2p-bootstrap"
	libp2p "gx/ipfs/QmWsV6kzPaYGBDVyuUfWBvyQygEc9Qrv9vzo8vZ7X4mdLN/go-libp2p"
	"gx/ipfs/QmeiCcJfDW1GJnWUArudsv5rQsihpi4oyddPhdqo3CfX6i/go-datastore"
	"time"
	"fmt"
)

var BootstrapPeers = []string{
	"/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	"/ip4/104.236.76.40/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64",
	"/ip4/104.236.176.52/tcp/4001/ipfs/QmSoLnSGccFuZQJzRadHn95W2CrSFmZuTdDWP8HXaHca9z",
	"/ip4/104.236.179.241/tcp/4001/ipfs/QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM",
	"/ip4/162.243.248.213/tcp/4001/ipfs/QmSoLueR4xBeUbY9WZ9xGUUxunbKWcrNFTDAadQJmocnWm",
	"/ip4/128.199.219.111/tcp/4001/ipfs/QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu",
}


func main() {
	
	//Create host
	h, err := libp2p.New(context.Background(), libp2p.Defaults)
	if err != nil {
		panic(err)
	}
	
	//Create DHT
	d := dht.NewDHTClient(context.Background(), h, datastore.NewMapDatastore())
	
	//Bootstrap object
	err, boot := bootstrap.NewBootstrap(h, bootstrap.Config{
		BootstrapPeers:    BootstrapPeers,
		MinPeers:          4,
		BootstrapInterval: time.Second * 5,
		HardBootstrap:     time.Second * 100,
	})
	if err != nil {
		panic(err)
	}
	
	//Start bootstrap
	if err := boot.Start(context.Background()); err != nil {
		panic(err)
	}
	
	//Exit on DHT bootstrap error
	if err := d.Bootstrap(context.Background()); err != nil {
		panic(err)
	}
	
	fmt.Println(time.Now())
	pi, err := d.FindPeer(context.Background(), "QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd")
	fmt.Println(time.Now())
	if err != nil {
		panic(err)
	}
	
	fmt.Println(pi)
	
}
