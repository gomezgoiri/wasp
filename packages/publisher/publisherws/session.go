package publisherws

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"nhooyr.io/websocket"
)

type WebSocketSession struct {
	OnMessage    func(connectionId string, message string)
	OnConnect    func(connectionId string)
	OnDisconnect func(connectionId string)

	readChannel  chan string
	writeChannel chan string

	connectionId string
	connection   *websocket.Conn
	context      context.Context
	shouldClose  bool
}

func NewWebSocketSession(connection *websocket.Conn, context context.Context) *WebSocketSession {
	return &WebSocketSession{
		OnConnect:    func(connectionId string) {},
		OnDisconnect: func(connectionId string) {},
		OnMessage:    func(connectionId string, message string) {},
		readChannel:  make(chan string, 10),
		writeChannel: make(chan string, 10),
		connectionId: uuid.NewString(),
		context:      context,
		connection:   connection,
	}
}

func InitializeWebSocketSessionFromRequest(responseWriter http.ResponseWriter, request *http.Request) (*WebSocketSession, error) {
	connection, err := websocket.Accept(responseWriter, request, &websocket.AcceptOptions{
		InsecureSkipVerify: true, // TODO: make accept origin configurable
	})
	if err != nil {
		return nil, err
	}

	return NewWebSocketSession(connection, request.Context()), nil
}

func (s *WebSocketSession) handleWrite(connectionContext context.Context) {
	for {
		if s.shouldClose {
			break
		}

		select {
		case msg := <-s.writeChannel:
			err := s.connection.Write(connectionContext, websocket.MessageText, []byte(msg))

			if err != nil {
				s.connection.Close(websocket.StatusInternalError, err.Error())
				return
			}
		}
	}
}

func (s *WebSocketSession) handleRead(connectionContext context.Context) {
	for {
		if s.shouldClose {
			break
		}

		msgType, msg, err := s.connection.Read(connectionContext)

		if err != nil {
			s.connection.Close(websocket.StatusInternalError, err.Error())
			return
		}

		if msgType != websocket.MessageText {
			continue
		}

		s.OnMessage(s.connectionId, string(msg))
	}
}

func (s *WebSocketSession) Close() {
	s.shouldClose = true
}

func (s *WebSocketSession) Listen() error {
	defer s.connection.Close(websocket.StatusInternalError, "something went wrong")
	connectionContext := s.connection.CloseRead(s.context)

	s.OnConnect(s.connectionId)

	go s.handleWrite(connectionContext)
	s.handleRead(connectionContext)

	s.shouldClose = true
	s.OnDisconnect(s.connectionId)

	return nil
}

func (s *WebSocketSession) WriteMessage(message string) error {
	select {
	case s.writeChannel <- message:
	default:
		return errors.New("message buffer is full")
	}

	return nil
}
