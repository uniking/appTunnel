package main
 
import (
	"fmt"
	"log"
	"net"
	"flag"
	"suninfo.com/utils"
)

func main() {
    src := flag.String("src", "0.0.0.0:8081", "要监听的地址")
    dest := flag.String("dest", "127.0.0.1:8082", "服务端的地址")
    proto := flag.String("type", "tcp", "隧道类型， tcp, http")
    flag.Parse()
    fmt.Println(*src, *dest, *proto)
	
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	fmt.Println("监听" + *src)
	l, err := net.Listen("tcp", *src)
	if err != nil {
		log.Panic(err)
	}
 
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
 
        go handleClientRequest(*dest, client, *proto)
	}
}

func handleClientRequest(dest string, client net.Conn, proto string) {
    fmt.Println("new client", proto)

    if "tcp" == proto{
	    var proxy utils.Proxy
	    proxy.ProxyClientTcpRequest(dest, client)
    }else if "http" == proto{
	    var proxy utils.Proxy
	    proxy.ProxyClientHttpRequest(dest, client)
    }else{
        log.Println("proto type error")
    }
}
