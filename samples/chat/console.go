package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var cReset = "\033[0m"
var cRed = "\033[31m"
var cGreen = "\033[32m"
var cYellow = "\033[33m"
var cBlue = "\033[34m"
var cPurple = "\033[35m"
var cCyan = "\033[36m"
var cGray = "\033[37m"
var cWhite = "\033[97m"

func (m msg) ToString() string {
	if bt, err := json.Marshal(m); err != nil {
		log.Printf("cant marshal msg %v", m)
		return ""
	} else {
		return string(bt)
	}
}

func printHelp(user string) {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	fmt.Println(cCyan + "hi " + cBlue + user)
	fmt.Println(cYellow + "   +group " + cCyan + "- to be added to a group. e.g. +family")
	fmt.Println(cYellow + "   -group " + cCyan + "- to be removed from a group. e.g. -covid")
	fmt.Println(cYellow + "   @group " + cCyan + "- to select current group. e.g. @meetup")
	fmt.Println(cYellow + "   --help " + cCyan + "-to display this menu")
	fmt.Println(cYellow + "   anything else" + cCyan + " will be published on you current group")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
}

func wrapSystemMsg(m string) string {
	return cCyan + m + cReset
}

func getUser(reader *bufio.Reader) string {
	fmt.Println(wrapSystemMsg("welcome to the meetup chat"))
	fmt.Println(wrapSystemMsg("*****type your name please****"))
	user, _ := reader.ReadString('\n')
	return strings.Replace(user, "\n", "", -1) //remove CRLF
}

func showChatOnConsole(nc *nats.Conn) {

	reader := bufio.NewReader(os.Stdin)
	user := getUser(reader)
	printHelp(user)

	var curGroup string
	subs := make(map[string]*nats.Subscription)

	prompt := func() {
		fmt.Print(cReset + curGroup + "$ ")
	}

	printMsg := func(s []byte, group string) {
		var m msg
		if err := json.Unmarshal([]byte(s), &m); err != nil {
			log.Printf("cant unmarshal msg %s", s)
		} else if m.User != user { //dont display my own messages
			fmt.Printf(cPurple+"\n%s:(%s):%s\n"+cReset, group, m.User, m.Message) //display time ??
			prompt()
		}
	}

	for {
		prompt()
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1) //remove CRLF

		if len(text) < 2 {
			continue
		}

		if text == "--help" {
			printHelp(user)
			continue
		}

		content := text[1:]

		switch text[:1] {
		case "+":
			if ok, err := regexp.MatchString("^[a-zA-Z0-9_.-]*$", content); !ok || err != nil {
				fmt.Printf("group name can only use letters, numbers and dash\n")
			} else if s, err := subscribe(nc, content, printMsg); err != nil {
				fmt.Printf("cant join group %s : %s\n", content, err.Error())
			} else {
				subs[content] = s
			}
		case "-":
			if s, ok := subs[content]; !ok {
				fmt.Printf("not registered to group %s\n", content)
			} else {
				s.Unsubscribe()
				delete(subs, content)
				fmt.Printf("I'm not on the %s group anymore :(\n", content)
				if content == curGroup {
					curGroup = ""
				}
			}
		case "@":
			if _, ok := subs[content]; !ok {
				fmt.Printf("you are not registered to group %s\n", content)
			} else {
				curGroup = content
			}
		default:
			if curGroup == "" {
				fmt.Println("you must first select a group, type @family for example")
			} else {
				m := msg{User: user, Message: text, Time: time.Now().Format("15:06")}
				publish(nc, curGroup, m.ToString())
			}
		}
	}
}
