package chat

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
	"labix.org/v2/mgo/bson"
	"log"
	"yhl/help"
)

const pastSize = 5

type Server struct {
	pattern   string
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *Message
	doneCh    chan bool
	errCh     chan error
}

func NewServer(pattern string) *Server {
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) SendAll(msg *Message) {
	s.sendAllCh <- msg
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) sendPastMessages(c *Client) {
	pasMsg := []*Message{}
	help.MongoDb.C("messages").Find(bson.M{"type": "message", "gid": bson.M{"$in": c.follow}}).Sort("-createtime").Limit(3).All(&pasMsg)

	for i := len(pasMsg) - 1; i >= 0; i-- {
		c.Write(pasMsg[i])
	}
}

func (s *Server) sendAll(msg *Message) {
	for _, c := range s.clients {
		for _, f := range c.follow {
			if msg.Gid == f {
				c.Write(msg)
				break
			}
		}
	}
}

func (s *Server) Listen() {

	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		//s.Add(client)
		client.Listen()
	}
	beego.Handler(s.pattern, websocket.Handler(onConnected))

	for {
		select {
		case c := <-s.addCh:
			s.clients[c.id] = c
			s.sendPastMessages(c)
		case c := <-s.delCh:
			delete(s.clients, c.id)
		case msg := <-s.sendAllCh:
			s.sendAll(msg)
		case err := <-s.errCh:
			log.Println("Error:", err.Error())
		case <-s.doneCh:
			return
		}
	}
}
