package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sausheong/gwp/xmpp_tester/entity"
	"github.com/sausheong/gwp/xmpp_tester/xmpp"
	"math/rand"
	"os"
	"strings"
	"sync"
)

var server = flag.String("server", "localhost:5222", "server")
var username = flag.String("username", "admin@laptop-d5d42j5u", "username")
var password = flag.String("password", "admin", "password")
var status = flag.String("status", "chat", "status")
var statusMessage = flag.String("status-msg", "I for one welcome our new codebot overlords.", "status message")
var notls = flag.Bool("notls", true, "No TLS")
var debug = flag.Bool("debug", false, "debug output")
var session = flag.Bool("session", false, "use server session")

var options xmpp.Options

func serverName(host string) string {
	return strings.Split(host, ":")[0]
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: example [options]\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	if *username == "" || *password == "" {
		if *debug && *username == "" && *password == "" {
			fmt.Fprintf(os.Stderr, "no username or password were given; attempting ANONYMOUS auth\n")
		} else if *username != "" || *password != "" {
			flag.Usage()
		}
	}

	if !*notls {
		xmpp.DefaultConfig = tls.Config{
			ServerName:         serverName(*server),
			InsecureSkipVerify: false,
		}
	} else {
		xmpp.DefaultConfig = tls.Config{
			ServerName:         serverName(*server),
			InsecureSkipVerify: true,
		}
	}

	options = xmpp.Options{
		Host:          *server,
		User:          *username,
		Password:      *password,
		NoTLS:         *notls,
		Debug:         *debug,
		Session:       *session,
		Status:        *status,
		StatusMessage: *statusMessage,
	}

}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	userA := entity.NewUserClient(&xmpp.Options{
		Host:          *server,
		User:          *username,
		Password:      *password,
		NoTLS:         *notls,
		Debug:         *debug,
		Session:       *session,
		Status:        *status,
		StatusMessage: *statusMessage,
	}, append([]string{}, "test2@laptop-d5d42j5u"))

	userB := entity.NewUserClient(&xmpp.Options{
		Host:          *server,
		User:          "test2@laptop-d5d42j5u",
		Password:      "123123",
		NoTLS:         *notls,
		Debug:         *debug,
		Session:       *session,
		Status:        *status,
		StatusMessage: *statusMessage,
	}, append([]string{}, "admin@laptop-d5d42j5u"))

	userC := entity.NewUserClient(&xmpp.Options{
		Host:          *server,
		User:          "test3@laptop-d5d42j5u",
		Password:      "123123",
		NoTLS:         *notls,
		Debug:         *debug,
		Session:       *session,
		Status:        *status,
		StatusMessage: *statusMessage,
	}, append([]string{}, "admin@laptop-d5d42j5u"))

	userD := entity.NewUserClient(&xmpp.Options{
		Host:          *server,
		User:          "test@laptop-d5d42j5u",
		Password:      "123",
		NoTLS:         *notls,
		Debug:         *debug,
		Session:       *session,
		Status:        *status,
		StatusMessage: *statusMessage,
	}, append([]string{}, "admin@laptop-d5d42j5u"))


	msgCount := 100
	users := []entity.User{*userA, *userB, *userC, *userD}

	var sendWg sync.WaitGroup
	sendWg.Add(msgCount * 4)

	userA.Start()
	userB.Start()
	userC.Start()
	userD.Start()

	for i := 0; i < msgCount; i++ {
		for j := 0; j < len(users); j++ {
			u := users[j]
			size := len(u.Friends)
			uuid, err := uuid.NewV4()
			if err != nil {
				fmt.Println("generate random message text fail")
				continue
			}
			r := rand.Intn(size)
			go func() {
				u.SendMessage(u.Friends[r], uuid.String())
				sendWg.Done()
			}()
		}
	}
	sendWg.Wait()

	wg.Wait()
}
