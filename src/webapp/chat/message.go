package chat

import (
	"strconv"
)

type Message struct {
	Uid        int    `json:"uid"`
	Gid        int    `json:"gid"`
	Type       string `json:"type"`
	Follow     []int  `json:"follow"`
	ToUid      int    `json:"toUid"`
	NickName   string `json:"nickName"`
	GroupName  string `json:"groupName"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

func (self *Message) String() string {
	return strconv.Itoa(self.Uid) + " says " + self.Content
}
