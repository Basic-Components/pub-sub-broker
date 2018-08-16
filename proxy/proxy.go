//
// 发布订阅模式的代理组件,沟通发布者和订阅者,降低发布者负载
//
package proxy

import (
	"github.com/Basic-Components/pub-sub-broker/consts"
	loadconfig "github.com/Basic-Components/pub-sub-broker/loadconfig"

	zmq "github.com/pebbe/zmq4"
	log "github.com/sirupsen/logrus"
)

// 代理本体
func Run(config loadconfig.Config) {
	//  Prepare our sockets
	collector, _ := zmq.NewSocket(zmq.PULL)
	defer collector.Close()

	backend, _ := zmq.NewSocket(zmq.PUB)
	defer backend.Close()

	if config.Conflate {
		collector.SetConflate(true)
		backend.SetConflate(true)
	} else {
		if config.RCVHWM >= 0 {
			collector.SetRcvhwm(config.RCVHWM)
		}
		if config.SNDHWM >= 0 {
			backend.SetSndhwm(config.SNDHWM)
		}
	}
	backend.Bind(config.BackendURL)
	collector.Bind(config.FrontendURL)
	for {
		msg, _ := collector.Recv(0)
		log.WithFields(log.Fields{
			consts.TYPE: consts.NAME,
			"Direction": "pull"}).Debug("pulled message!")
		backend.Send(msg, 0)
		log.WithFields(log.Fields{
			consts.TYPE: consts.NAME,
			"Direction": "publish"}).Debug("publish message to subscriber!")
	}
}
