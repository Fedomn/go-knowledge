package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:8082", nil))
	}()

	time.Sleep(time.Minute * 5)
}

// 1. first capture
// open http://localhost:8082/debug/pprof/
// go tool pprof -svg http://localhost:8082/debug/pprof/profile > cpu.svg
// go tool pprof -svg http://localhost:8082/debug/pprof/heap > heap.svg

// 2. second
// open generated pprof files
// go tool pprof -http=:8081 ~/pprof.cpu.pb.gz

/**
package main

import (
"log"
"time"

"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Hour)
}
*/
