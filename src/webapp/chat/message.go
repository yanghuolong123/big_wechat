package chat

import (
//"time"
)

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
	//Createtime time.Time `json:"createtime"`
	Createtime string `json:"createtime"`
}

func (self *Message) String() string {
	return self.Author + " says " + self.Body
}
