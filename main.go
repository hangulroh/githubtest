package main

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	fmt.Println("hello")
	_, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://254.0.0.1:12345"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		panic("panic")
	}

}
