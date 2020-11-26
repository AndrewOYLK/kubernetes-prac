package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/client"
)

func main() {
	cfg := client.Config{
		Endpoints:               []string{"http://10.1.0.15:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)
	log.Print("Setting!")
	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Set is done. \n")
	}

	resp, err = kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
		log.Println(resp.Node.Key, "===", resp.Node.Value)
	}
}
