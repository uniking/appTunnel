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
    flag.Parse()
    fmt.Println(*src, *dest)
	
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
 
		go handleClientRequest(*dest, client)
	}
}

func handleClientRequest(dest string, client net.Conn) {
	var proxy utils.Proxy
	proxy.ProxyClientRequest(dest, client)
}
