package http

import (
	"strings"
	"log"
	"net"
	"io"
    "fmt"
)

type HttpProxy struct{
	method string
	path string
	version string
	body map[string]string
	other []string
}

func (http *HttpProxy)Parse(content string){
    http.body = make(map[string]string)
    fmt.Println(content)
    lines := strings.Split(content, "\r\n")
    head := strings.Split(lines[0], " ")
    if len(head) != 3{
            http.path = "http"
            return
    }
    
    http.method = head[0]
    http.path = head[1]
    http.version = head[2]
    
    ln := len(lines)
    for i:=1 ; i<ln ; i++ {
        //": "有空格
        b:=strings.Split(lines[i], ": ")
        if len(b) == 2 {
            http.body[b[0]]=b[1]
        }else{
            //非 ": "分割
            if len(lines[i]) > 2 {//"\r\n"
                http.other = append(http.other, lines[i])
            }
        }
    }
}

func (http *HttpProxy)ToString() string{
    var content string
    content = http.method + " " + http.path + " " + http.version + "\r\n"
    
    for k, v := range http.body {
        //": "有空格
        content = content + k + ": " + v + "\r\n"
    }
    
    for _, v := range http.other{
        content = content + v + "\r\n"
    }
    
    content = content + "\r\n"
    return content
}

func delBlank(content string) string{
    if content[0] == ' ' {
        return content[1:]
    }else{
        return content
    }
}

func (http *HttpProxy)BuildProxyData(content string) string{
    http.Parse(content)
    if strings.HasPrefix(http.path, "http"){
        return content
    }
    
    http.body["Proxy-Connection"] = " keep-alive"
    
    host := http.body["Host"]
    s := strings.Split(host, ":")
    fmt.Println("host, ", host)
    if len(s) >= 2 {
        if s[1] == "80" {
            http.path = "http://" + delBlank(http.body["Host"]) + ":80" + http.path
        }else{
            http.path = "https://" + delBlank(http.body["Host"]) + ":" + s[1] + http.path
        }
    }else{
        http.path = "http://" + delBlank(http.body["Host"]) + ":80" + http.path
    }
    
    newcontent := http.ToString()
    
    return newcontent
}

func (http *HttpProxy)BuildProxyDataForbytes(b []byte) []byte{
    return []byte(http.BuildProxyData(string(b[:])))
}

func (http *HttpProxy)HandleClientRequest(client net.Conn, server net.Conn, appendB []byte) {
    
    var b [10240]byte
    
    if len(appendB) > 0 {
        server.Write(http.BuildProxyDataForbytes(appendB[:]))
    }
	
	//server to client
	go io.Copy(client, server)
	//client to server
	//go io.Copy(server, client)
	for{
		n, err := client.Read(b[:])
		if err != nil {
			log.Println(err)
			break
		}
		server.Write(http.BuildProxyDataForbytes(b[:n]))
	}
}
