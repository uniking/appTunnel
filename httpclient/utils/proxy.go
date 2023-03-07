package utils

import (
    "log"
    "strings"
    "bytes"
	"net"
    "suninfo.com/utils/http"
    "suninfo.com/utils/https"
    "suninfo.com/utils/tcp"
)

type Proxy struct{
	vertion string
}

func (proxy *Proxy) ProxyClientHttpRequest(dest string, client net.Conn){
	if client == nil {
		return
	}
	defer client.Close()
	
	var b []byte
	
	server, err := net.Dial("tcp", dest)
	if err != nil {
		log.Println(err)
		return
	}

    var n int = 0
	//读取目标地址, 
    var buffer bytes.Buffer
    for n < 12 {
        var bl [256]byte
	    nl, errl := client.Read(bl[:])
	    if errl != nil {
		    log.Println(errl)
		    return
	    }
        buffer.Write(bl[:nl])
        n = n+nl
    }
    b = buffer.Bytes()

	var firstmsg string
	if n > 12 {
		firstmsg = string(b[1:64])
	}else{
		firstmsg = string(b[1:n])
	}

	
	if strings.HasPrefix(firstmsg, "TARGET_HOST"){
		index := bytes.IndexByte(b[:], '\n')
		target_host := string(b[:index])
		appendB := b[index+1:n]
        
        s := strings.Split(target_host, ":")
		//TARGET_HOST:uuid:appid:cn.bing.com:443
        if len(s) == 5 {
            port := s[2]
            if port == "80" {
                //http
                var httpProxy http.HttpProxy
                httpProxy.HandleClientRequest(client, server, appendB)
            }else{
                //https
                var httpsProxy https.HttpsProxy
                httpsProxy.HandleClientRequest(client, server, target_host, appendB)
            }
        }else{
			log.Println("TARGET_HOST form error")
		}
	}else{
		log.Println("proxy not start with TARGET_HOST")
	}
}

func (proxy *Proxy) ProxyClientTcpRequest(dest string, client net.Conn){
	if client == nil {
		return
	}
	defer client.Close()
	var b []byte
    var n int = 0
	//读取目标地址, 
    var buffer bytes.Buffer
    for n < 12 {
        var bl [256]byte
	    nl, errl := client.Read(bl[:])
	    if errl != nil {
		    log.Println(errl)
		    return
	    }
        buffer.Write(bl[:nl])
        n = n+nl
    }
    b = buffer.Bytes()

	var firstmsg string
	if n > 12 {
		firstmsg = string(b[1:64])
	}else{
		firstmsg = string(b[1:n])
	}

	////TARGET_HOST:uuid:appid:cn.bing.com:443
	if strings.HasPrefix(firstmsg, "TARGET_HOST"){
		index := bytes.IndexByte(b[:], '\n')
		target_host := string(b[1:index])
		appendB := b[index+1:n]
        
        s := strings.Split(target_host, ":")
        if len(s) == 5 {
			log.Println("get TARGET_HOST", target_host)
            var tcpProxy tcp.TcpProxy
            tcpProxy.HandleClientRequest(client, target_host, appendB)
        }else{
			log.Println("TARGET_HOST form error ", target_host)
		}
	}else{
		log.Println("proxy not start with TARGET_HOST")
	}
}
