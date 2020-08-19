package main

import (
	"fmt"
	//_ "github.com/micro/go-plugins/registry/consul"
	"go-rpc-happysooner-file/internal"
)

func main() {
	fmt.Println("cmd/app  go init~")
	internal.Init()

}
