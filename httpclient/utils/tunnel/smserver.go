package tunnel

import (
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"
	"io/ioutil"
	"log"
	"net"
	"io"
)

func StartSTLServer(src string, dest string) {
	// 信任的根证书
	certPool := x509.NewCertPool()
	cacert, err := ioutil.ReadFile("root.cer")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(cacert)
	sigCert, _ := gmtls.LoadX509KeyPair("sm2_sign_cert.cer", "sm2_sign_key.pem")
	encCert, _ := gmtls.LoadX509KeyPair("sm2_enc_cert.cer", "sm2_enc_key.pem")
	
	config := &gmtls.Config{
		GMSupport:    &gmtls.GMSupport{},
		RootCAs:      certPool,
		Certificates: []gmtls.Certificate{sigCert, encCert},
	}

	l, err  := gmtls.Listen("tcp", "localhost:50052", config)
	if err != nil {
		log.Panic(err)
	}
	
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
 
		go handleClientRequestServer(dest, client)
	}
}

func handleClientRequestServer(dest string, client net.Conn) {
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
