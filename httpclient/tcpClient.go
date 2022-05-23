package main
 
import (
	"fmt"
	"log"
	"net"
	"flag"
    "time"
    "sync"
)

func main() {
    dest := flag.String("dest", "127.0.0.1:8082", "服务端的地址")
    gnum := flag.Int("gnum", 10, "连接数")
    flag.Parse()
    fmt.Println(*dest, *gnum)

    var wg sync.WaitGroup
    wg.Add(*gnum)

    for i:=0; i<*gnum; i++{
        fmt.Println("new goroutine")
        go clientRequest(*dest, &wg)
    }

    wg.Wait()
}

func clientRequest(dest string, waitGroup *sync.WaitGroup) {
    defer waitGroup.Done()

	//建立tcp tunnel
	server, err := net.Dial("tcp", dest)
	if err != nil {
		log.Println(err)
		return
	}

    var b [1024]byte

    for{
        server.Write([]byte("hello Server"))
		_, err := server.Read(b[:])
		if err != nil {
			log.Println(err)
			break
		}
        time.Sleep(1 * time.Second)
    }

    fmt.Println("goroutine exit")
}
