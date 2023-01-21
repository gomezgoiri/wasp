package publisherws

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"net/http"

	"nhooyr.io/websocket"
)

type WebSocketSession struct {
	OnMessage    func(connectionId string, message []byte)
	OnConnect    func(connectionId string)
	OnDisconnect func(connectionId string)
	OnError      func(connectionId string, err error)

	readChannel  chan []byte
	writeChannel chan []byte

	connectionID  string
	connection    *websocket.Conn
	context       context.Context
	cancelContext context.CancelFunc
}

func NewWebSocketSession(ctx context.Context, connection *websocket.Conn) *WebSocketSession {
	ctx, cancel := context.WithCancel(ctx)

	return &WebSocketSession{
		OnConnect:     func(connectionId string) {},
		OnDisconnect:  func(connectionId string) {},
		OnMessage:     func(connectionId string, message []byte) {},
		OnError:       func(connectionId string, err error) {},
		readChannel:   make(chan []byte, 10),
		writeChannel:  make(chan []byte, 10),
		connectionID:  uuid.NewString(),
		context:       ctx,
		cancelContext: cancel,
		connection:    connection,
	}
}

func AcceptWebSocketSession(responseWriter http.ResponseWriter, request *http.Request) (*WebSocketSession, error) {
	connection, err := websocket.Accept(responseWriter, request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}

	return NewWebSocketSession(request.Context(), connection), nil
}

func (s *WebSocketSession) handleWrite(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case msg := <-s.writeChannel:
			if err := s.connection.Write(ctx, websocket.MessageText, msg); err != nil {
				s.OnError(s.connectionID, err)
				return
			}
		}
	}
}

func (s *WebSocketSession) handleRead(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msgType, msg, err := s.connection.Read(ctx)
			if err != nil {
				s.OnError(s.connectionID, err)
				return
			}

			if msgType != websocket.MessageText {
				continue
			}

			s.OnMessage(s.connectionID, msg)
		}
	}
}

func (s *WebSocketSession) Close() {
	s.cancelContext()
}

func (s *WebSocketSession) Listen() error {
	defer s.connection.Close(websocket.StatusInternalError, "something went wrong")
	defer s.OnDisconnect(s.connectionID)

	s.OnConnect(s.connectionID)

	go s.handleWrite(s.context)
	s.handleRead(s.context)

	return s.connection.Close(websocket.StatusNormalClosure, "")
}

func (s *WebSocketSession) WriteMessage(message []byte) error {
	select {
	case s.writeChannel <- message:
	default:
		return errors.New("message buffer is full")
	}

	return nil
}
