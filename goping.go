package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var (
		//prot    string
		address string
		port    string
		url     string
		prot    string
	)
	if len(os.Args) < 3 {
		fmt.Println(`
使用方法1: goping  www.baidu.com 80	  默认使用TCP协议
使用方法2: goping [TCP/UDP]  www.baidu.com 80  指定使用协议
		`)
		return

	} else if len(os.Args) < 4 {
		address = os.Args[1]
		port = os.Args[2]
		url = address + ":" + port
		prot = "tcp"
	} else if len(os.Args) == 4 {
		prot = os.Args[1]
		address = os.Args[2]
		port = os.Args[3]
		url = address + ":" + port
	}

	for {
		star_time := time.Now()
		conn, err := net.Dial(prot, url)
		if err != nil {
			fmt.Printf("与%v建立连接失败,失败原因:%v", address, err)
		}
		defer conn.Close()
		fmt.Printf("连接地址:%v ~~~  延迟:%v\n", url, time.Now().Sub(star_time))
		time.Sleep(time.Second)
		fmt.Println("test")
	}

}
