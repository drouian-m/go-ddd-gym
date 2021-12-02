package adapter

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"ziggornif.xyz/ddd-gym/domains/formules/model"
	"ziggornif.xyz/ddd-gym/domains/formules/port"
)

type jetstreamDispatcher struct {
	js nats.JetStreamContext
}

const (
	streamName     = "FORMULE"
	streamSubjects = "FORMULE.*"
)

func NewJetStreamDispatcher() port.Dispatcher {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream() // Returns JetStreamContext
	_, err := js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{streamSubjects},
	})

	if err != nil {
		log.Println(err)
	}

	return &jetstreamDispatcher{js}
}

func (jsd *jetstreamDispatcher) Emit(event model.Event) {
	var subject string
	switch event.Name {
	case "formule-created":
		subject = "FORMULE.created"
	}

	jsonEvent, _ := json.Marshal(event)
	ack, err := jsd.js.Publish(subject, jsonEvent)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Event has been published %v\n", ack.Sequence)
}
