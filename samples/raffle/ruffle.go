// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var prizes = []string{"Jetbrains Goland Ide #1",
	"Jetbrains Goland Ide #2",
	"Jetbrains Goland Ide #3",
	"Secure coding in go course",
	"Go recipes - practical projects course",
	"Go-essential-training course"}

var cReset = "\033[0m"
var cYellow = "\033[33m"
var cBlue = "\033[34m"

func main() {

	if data, err := ioutil.ReadFile("RSVP-yes"); err != nil {
		fmt.Println("cant get data from RSVP-yes file")
	} else {
		names := strings.Split(string(data), "\n")

		reader := bufio.NewReader(os.Stdin)

		rand.Seed(time.Now().Unix())
		for i := 0; i < 6 && i < len(names); i++ {
			n := rand.Int() % len(names)
			for cont := false; cont == false; {
				fmt.Printf("%s%s%s won %s'%s'%s !\n        (A)Accept,(R)Reshuffle : ", cBlue, names[n], cReset, cYellow, prizes[i], cReset)
				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1) //remove CRLF
				switch strings.ToLower(text) {
				case "a":
					cont = true
				case "r":
					i -= 1
					cont = true
				default:
					fmt.Println("please choose A to Accept or R to Reshuffle a new winner")
				}
			}
			names = append(names[:n], names[n+1:]...) //dont allow same name twice
		}
	}
}
