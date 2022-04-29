package tunnel

import (
	"log"
	"net"
	"io"
)

func StartTcpTunnel(src string, dest string) {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	l, err := net.Listen("tcp", src)
	if err != nil {
		log.Panic(err)
	}
 
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
 
		go handleClientTcpRequest(dest, client)
	}
}

func handleClientTcpRequest(dest string, client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	server, err := net.Dial("tcp", dest)
	if err != nil {
		log.Println(err)
		return
	}
	
	go io.Copy(server, client)
	io.Copy(client, server)
}
