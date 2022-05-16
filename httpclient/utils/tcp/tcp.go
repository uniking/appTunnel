package tcp
 
import (
	"fmt"
	"io"
	"log"
	"net"
)

type TcpProxy struct{
	vertion string
}

func (https *TcpProxy) HandleClientRequest(client net.Conn, target_host string, appendB []byte) {
	//解析host, port
	var host string
	fmt.Sscanf(target_host, "TARGET_HOST:%s:%d\n", &host)
	
	//建立tcp tunnel
	server, err := net.Dial("tcp", host)
	if err != nil {
		log.Println(err)
		return
	}
	
	//如果有附加数据，发送
	if len(appendB) > 0 {
		server.Write(appendB[:])
	}
	
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}
