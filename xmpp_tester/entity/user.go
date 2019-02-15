package entity

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sausheong/gwp/xmpp_tester/xmpp"
	"log"
	"math/rand"
)

type User struct {
	Name    string
	Client  *xmpp.Client
	Friends []string
}

//登录
func NewUserClient(options *xmpp.Options, friends []string) *User {
	client, err := options.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return &User{
		Name:    options.User,
		Client:  client,
		Friends: friends,
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
				fmt.Printf("%s received message from %s: %s\n", user.Name, v.Remote, v.Text)
			case xmpp.Presence:
				fmt.Println(v.From+": ", v.Show)
			}
		}
	}()
}

//发送消息
func (user *User) SendMessage(targetJID, message string) {
	//go func() {
		chat := xmpp.Chat{Remote: targetJID, Type: "chat", Text: message}
		user.Client.Send(chat)
	//}()
}

//发送消息给随机好友
func (user *User) SendMessageToRandomFriend(times int) {
	size := len(user.Friends)
	for i := 0; i < times; i++ {
		uuid, err := uuid.NewV4()
		if err != nil {
			fmt.Println("generate random message text fail")
			continue
		}
		r := rand.Intn(size)
		user.SendMessage(user.Friends[r], uuid.String())
	}
}

//注销

//获取花名册
func (user *User) GetRoster() {
	err := user.Client.Roster()
	if err != nil {
		fmt.Println("获取花名册失败")
	}
}
