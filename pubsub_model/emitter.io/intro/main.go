package main

import (
	"fmt"
	"log"

	emitter "github.com/emitter-io/go/v2"
)

func main() {

	// Create the client and connect to the broker
	c, err := emitter.Connect("tcp://localhost:8073", func(_ *emitter.Client, msg emitter.Message) {
		fmt.Printf("[emitter] -> [B] received: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	})
	if err != nil {
		log.Fatal(err)
	}

	key, chanName, err := c.GenerateKey("PtW4_f4m38p0H49Ywy_sHKcjhVt3XeFR", "sdk-integration-test/", "rwls", 600)
	if err != nil {
		log.Fatal(err)
	}

	_ = key
	_ = chanName
	log.Println("Key: ", key)
	log.Println("ChanName: ", chanName)

	// Set the presence handler
	c.OnPresence(func(_ *emitter.Client, ev emitter.PresenceEvent) {
		log.Printf("[emitter] -> [B] presence event: %d subscriber(s) at topic: '%s'\n", len(ev.Who), ev.Channel)
	})

	log.Println("1: [emitter] <- [B] querying own name")
	id := c.ID()
	log.Println("2: [emitter] -> [B] my name is " + id)

	// Subscribe to sdk-integration-test channel
	log.Printf("3: [emitter] <- [B] subscribing to '%s'\n", chanName)
	c.Subscribe(key, "sdk-integration-test/", func(_ *emitter.Client, msg emitter.Message) {
		log.Printf("SUB [emitter] -> [B] received on specific handler: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	})

	// Ask for presence
	log.Printf("4: [emitter] <- [B] asking for presence on '%s'\n", chanName)
	c.Presence(key, "sdk-integration-test/", true, false)

	// Publish to the channel
	log.Printf("5: [emitter] <- [B] publishing to '%s'\n", chanName)
	c.Publish(key, "sdk-integration-test/", "{\"msg\": \"hello\"}")

	select {} // wait forever without busy loop - Ctrl-c to stop

}
