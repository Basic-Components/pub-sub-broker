//
// 发布订阅模式的代理组件,沟通发布者和订阅者,降低发布者负载
//
package proxy

import (
	"log"

	loadconfig "github.com/Basic-Components/pub-sub-broker/loadconfig"

	zmq "github.com/pebbe/zmq4"
)

// 代理本体
func Run(config loadconfig.Config) {
	//  Prepare our sockets
	frontend, _ := zmq.NewSocket(zmq.XSUB)
	defer frontend.Close()
	backend, _ := zmq.NewSocket(zmq.XPUB)
	defer backend.Close()
	frontend.Bind(config.FrontendURL)
	backend.Bind(config.BackendURL)

	//  Initialize poll set
	err := zmq.Proxy(frontend, backend, nil)
	log.Fatalln("Proxy interrupted:", err)
}
