package main
 
import (
	"flag"
	_ "net/http/pprof"
	"strconv"
	"time"
    //"fmt"
	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/log"
	"github.com/Allenxuxu/toolkit/sync/atomic"
)

type example struct {
	Count atomic.Int64
}

func (s *example) OnConnect(c *gev.Connection) {
	s.Count.Add(1)
	//log.Println(" OnConnect ： ", c.PeerAddr())
}
func (s *example) OnMessage(c *gev.Connection, ctx interface{}, data []byte) (out interface{}) {
	//fmt.Println("OnMessage")
	out = data
	return
}

func (s *example) OnClose(c *gev.Connection) {
	s.Count.Add(-1)
	//log.Println("OnClose")
}


func main() {
	handler := new(example)
	var port int
	var loops int

	flag.IntVar(&port, "port", 8082, "server port")
	flag.IntVar(&loops, "loops", -1, "num loops")
	flag.Parse()

	s, err := gev.NewServer(handler,
		gev.Network("tcp"),
		gev.Address(":"+strconv.Itoa(port)),
		gev.NumLoops(loops),
		gev.MetricsServer("", ":9091"),
	)
	if err != nil {
		panic(err)
	}

	s.RunEvery(time.Second*2, func() {
		log.Info("connections :", handler.Count.Get())
	})

	s.Start()
}
