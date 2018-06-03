package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// Args for RPC arg
type Args struct{}

// TimeServer a RPC Server
type TimeServer int64

// GiveServerTime return server time
func (t *TimeServer) GiveServerTime(Args *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// Create a new RPC server
	timeserver := new(TimeServer)
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	// Listen for requests on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}
	http.Serve(l, nil)
}
