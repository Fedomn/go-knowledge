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

// 3. optional use gops
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
// gops
// gops trace pid

// reference:
// https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/
// https://blog.golang.org/pprof

// useful commands:
// GOGC=2000 GODEBUG='gctrace=1' go run main.go 2>&1 | cat > gc.log
// go build -gcflags='-m -m'  2>&1 | grep 'escapes to heap'

// https://wudaijun.com/2020/01/go-gc-keypoint-and-monitor/
// https://golang.org/pkg/runtime/
