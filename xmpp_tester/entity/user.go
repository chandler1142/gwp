package entity

import (
	"fmt"
	"github.com/sausheong/gwp/xmpp_tester/xmpp"
	"log"
)

type User struct {
	Name   string
	Client *xmpp.Client
}

//登录
func NewUserClient(options *xmpp.Options) *User {
	client, err := options.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return &User{
		Name:   options.User,
		Client: client,
	}
}

func (user *User) Start() {
	//监听消息
	go func() {
		for {
			chat, err := user.Client.Recv()
			if err != nil {
				log.Fatal(err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				fmt.Println(v.Remote+": ", v.Text)
			case xmpp.Presence:
				fmt.Println(v.From+": ", v.Show)
			}
		}
	}()
}

//发送消息
func (user *User) SendMessage(targetJID, message string) {
	user.Client.Send(xmpp.Chat{Remote: targetJID, Type: "chat", Text: message})
}

//注销

//获取花名册
func (user *User) GetRoster() {
	err := user.Client.Roster()
	if err != nil {
		fmt.Println("获取花名册失败")
	}
}