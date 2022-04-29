package https
 
import (
	"fmt"
	"io"
	"log"
	"net"
)

type HttpsProxy struct{
	vertion string
}

func (https *HttpsProxy) HandleClientRequest(client net.Conn, server net.Conn, target_host string, appendB []byte) {
	
	var b [1024]byte
	
	//解析host, port
	var host string
	fmt.Sscanf(target_host, "TARGET_HOST:%s:%d\n", &host)
	
	//建立http tunnel
	sf := "CONNECT %s HTTP/1.1\r\nHost: %s\r\nProxy-Connection: keep-alive\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36\r\n\r\n"
	connectmsg := fmt.Sprintf(sf, host, host)
	fmt.Fprint(server, connectmsg)
	
	//等待proxy确认tunnel成功
	_, err := server.Read(b[:])
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
