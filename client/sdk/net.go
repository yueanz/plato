package sdk

const (
	MsgTypeText = "text"
)

type connection struct {
	serverAddr         string
	sendChan, recvChan chan *Message
}

func newConnection(serverAddr string) *connection {
	return &connection{
		serverAddr: serverAddr,
		sendChan:   make(chan *Message),
		recvChan:   make(chan *Message),
	}
}

func (c *connection) send(msg *Message) {
	c.recvChan <- msg
}

func (c *connection) recv() <-chan *Message {
	return c.recvChan
}
