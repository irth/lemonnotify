package main

import (
	"fmt"
	"github.com/pocke/oshirase"
	"io"
	"os"
	"os/exec"
)

type Notification struct {
	AppName string `json:"appName"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
}

func Notify(n Notification) {
	bar := exec.Command("lemonbar")
	stdin, err := bar.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}

	bar.Stdout = os.Stdout
	bar.Stderr = os.Stderr

	if err := bar.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}

	io.WriteString(stdin, "notification\n")

}

func main() {
	server, err := oshirase.NewServer("LemonNotify", "irth", "1.0")

	if err != nil {
		panic(err)
	}

	server.OnNotify(func(notification *oshirase.Notify) {
		var n = Notification{
			AppName: notification.AppName,
			Summary: notification.Summary,
			Body:    notification.Body,
		}
		Notify(n)
	})
	select {}
}
