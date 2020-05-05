package main

import (
    //"github.com/bairn/props"
    //"fmt"
    //"time"
    "github.com/bairn/props/kvs"
    "fmt"
)

func mainx() {

    //root := "config/app1/dev"
    //address := "127.0.0.1:8500"
    //
    //p := props.NewConsulPropsConfigSourceByName("consul-props", address, root, 10*time.Second)
    //fmt.Println(p.Get(""))
    command := "java"
    params := []string{"-jar", "zookeeper/mock.jar"}
    started := kvs.ExecCommand(command, params...)
    fmt.Println(started)
}
