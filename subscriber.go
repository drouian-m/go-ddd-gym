package main

import (
	"encoding/json"
	"log"
	"runtime"

	"ziggornif.xyz/ddd-gym/domains/formules/model"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create durable consumer monitor
	js.Subscribe("FORMULE.*", func(msg *nats.Msg) {
		msg.Ack()
		var formuleEvent model.Event
		err := json.Unmarshal(msg.Data, &formuleEvent)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("monitor service subscribes from subject:%s\n", msg.Subject)
		log.Printf("ID:%d, Name:%s\n", formuleEvent.ID, formuleEvent.Name)
	}, nats.Durable("monitor"), nats.ManualAck())

	runtime.Goexit()

}
