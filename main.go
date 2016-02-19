package main

import (
	"fmt"
	"github.com/pocke/oshirase"
	"io"
	"os"
	"os/exec"
	"time"
)

func Notify(n *oshirase.Notify) {
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
	io.WriteString(stdin, fmt.Sprintf("%%{F#ff0000}[%s]%%{F} %%{F#FF8C00}<%s>%%{F} %s\n", n.AppName, n.Summary, n.Body))
	if n.ExpireTimeout < 0 {
		n.ExpireTimeout = 2000
	}
	time.Sleep(time.Duration(n.ExpireTimeout) * time.Millisecond)
	stdin.Close()
}

func main() {
	server, err := oshirase.NewServer("LemonNotify", "irth", "1.0")

	if err != nil {
		panic(err)
	}

	server.OnNotify(func(notification *oshirase.Notify) {
		Notify(notification)
	})
	select {}
}
