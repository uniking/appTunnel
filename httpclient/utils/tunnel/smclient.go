package tunnel

import (
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"
	"io/ioutil"
	"log"
	"net"
	"io"
)

func StartSTLClient(src string, dest string) {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	l, err := net.Listen("tcp", src)
	if err != nil {
		log.Panic(err)
	}
 
    initCer()
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
 
		go handleClientRequestClient(dest, client)
	}
}


var config *gmtls.Config

func initCer(){
	// 信任的根证书
	certPool := x509.NewCertPool()
	cacert, err := ioutil.ReadFile("root.cer")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(cacert)
	cert, err := gmtls.LoadX509KeyPair("sm2_cli.cer", "sm2_cli.pem")

	config = &gmtls.Config{
		GMSupport:    &gmtls.GMSupport{},
		RootCAs:      certPool,
		Certificates: []gmtls.Certificate{cert},
	}
}

func handleClientRequestClient(dest string, client net.Conn) {
	server, err := gmtls.Dial("tcp", "localhost:50052", config)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	
	go io.Copy(server, client)
	io.Copy(client, server)
}
