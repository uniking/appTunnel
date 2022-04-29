package main
 
import (
	"fmt"
	"flag"
	"suninfo.com/utils/tunnel"
)

func main() {
    src := flag.String("src", "127.0.0.1:8082", "要监听的地址")
    dest := flag.String("dest", "192.168.199.112:3128", "服务端的地址")
	crypt := flag.Bool("crypt", false, "对隧道进行加密，当启用加密后，需指定-server")
	servertype := flag.Bool("server", false, "服务端，不指定会作为客户端，客户端需要指定证书和私钥，服务端需要指定加密证书和私钥，签名证书和私钥，以及根证书")
	rootcer := flag.String("rootcer", "/home/web/cer/root.cer", "根证书")
	
	
    flag.Parse()
    fmt.Printf("%s %s crypt=%d server=%d %s\n", *src, *dest, *crypt, *servertype, *rootcer)
	
	if *crypt {
		if *servertype {
			//server port
			tunnel.StartSTLServer(*src, *dest)
		}else{
			//client port
			tunnel.StartSTLClient(*src, *dest)
		}
	}else{
		tunnel.StartTcpTunnel(*src, *dest)
	}
}

