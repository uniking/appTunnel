package tcp
 
import (
	"fmt"
	"log"
	"net"
	"io"
	"strings"
)

type TcpProxy struct{
	vertion string
}

func selfCopy(server net.Conn, client net.Conn) {
	defer client.Close()
	defer server.Close()

	//var n int = 0
	//var buffer bytes.Buffer
	var bl [10240]byte
	for {
		nl, errl := client.Read(bl[:])
		if errl != nil {
			//if n != 0 {
			//	server.Write(buffer.Bytes()[:n])
			//}
			if errl == io.EOF {
				//io.EOF 是用来表示连接正常关闭的值
				//测试不能调用Close, Close是双向关闭而不是单向关闭
				fmt.Println("close")
			}else{
				fmt.Println("selfCopy read error", errl.Error())
			}
			return
		}
		//fmt.Println("selfCopy read ok")
		//buffer.Write(bl[:nl])
		//n = n + nl
		//if n == 1 {
		//	continue
		//}
		//fmt.Println("to server", string(buffer.Bytes()[:n]))
		//server.Write(buffer.Bytes()[:n])
		//n = 0
		//buffer.Reset()
		nl, errl = server.Write(bl[:nl])
		if errl != nil {
			//io.EOF 是用来表示连接正常关闭的值
			fmt.Println("selfCopy write error", errl.Error())
			//if n != 0 {
			//	server.Write(buffer.Bytes()[:n])
			//}
			return
		}
	}

}

func (https *TcpProxy) HandleClientRequest(client net.Conn, target_host string, appendB []byte) {
	//解析host, port TARGET_HOST:uuid:appid:cn.bing.com:443
	s := strings.Split(target_host, ":")
	
	//建立tcp tunnel
	server, err := net.Dial("tcp", s[3]+":"+s[4])
	if err != nil {
		log.Println(err)
		return
	}
	
	//如果有附加数据，发送
	if len(appendB) > 0 {
		server.Write(appendB[:])
	}
	
	//进行转发
	//go io.Copy(server, client)
	//io.Copy(client, server)
	
	go selfCopy(server, client)
	selfCopy(client, server)
}
