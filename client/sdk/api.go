package sdk

type Chat struct {
	Nickname   string
	UserId     string
	SessionId  string
	connection *connection
}

type Message struct {
	Type       string
	Name       string
	FromUserID string
	ToUserID   string
	Content    string
	Session    string
}

func NewChat(serverAddr, nickname, userId, sessionId string) *Chat {
	return &Chat{
		Nickname:   nickname,
		UserId:     userId,
		SessionId:  sessionId,
		connection: newConnection(serverAddr),
	}
}

func (chat *Chat) Send(msg *Message) {
	chat.connection.send(msg)
}

func (chat *Chat) Recv() <-chan *Message {
	return chat.connection.recv()
}

func (chat *Chat) Close() {
	//todo
}
