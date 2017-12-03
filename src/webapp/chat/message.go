package chat

import (
	"strconv"
)

type Message struct {
	Uid        int    `json:"uid"`
	Content    string `json:"content"`
	Createtime string `json:"createtime"`
}

func (self *Message) String() string {
	return strconv.Itoa(self.Uid) + " says " + self.Content
}
