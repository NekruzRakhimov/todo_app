package broker

import (
	"github.com/nats-io/nats.go"
	"log"
)

func NewNatsConnection(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		log.Println("Couldn't connect to nats: ", err.Error())
		return nil, err
	}

	return nc, nil
}

func CloseNatsConnection(conn *nats.Conn) {
	conn.Close()
}
