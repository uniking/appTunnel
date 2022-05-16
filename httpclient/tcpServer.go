package main
 
import (
	"fmt"
	"log"
	"net"
	"flag"
)

func main() {
    src := flag.String("src", "0.0.0.0:8821", "要监听的地址")
    flag.Parse()
    fmt.Println(*src)
	
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
 
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
    var b [1024]byte
    if addr, ok := client.RemoteAddr().(*net.TCPAddr); ok {
        fmt.Println("client ip:", addr.IP.String())
    }
    for{
        n, err := client.Read(b[:])
        if err != nil {
            log.Println(err)
            break
        }
        fmt.Println(string(b[:n]))
    }
}