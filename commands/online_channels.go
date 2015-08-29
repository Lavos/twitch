package main

import (
	"os"
	"fmt"
	"log"
	"github.com/Lavos/twitch"
	"github.com/daviddengcn/go-colortext"
)

func main () {
	username := os.Getenv("TWITCH_USERNAME")

	if username == "" {
		log.Fatalf("No username found in TWITCH_USERNAME.")
	}

	tc := &twitch.TwitchClient{username}

	streams, err := tc.Online()

	if err != nil {
		log.Fatal(err)
	}

	ct.ChangeColor(ct.Yellow, true, ct.Black, false)
	fmt.Printf("Twitch Channels Online: %d\n", len(streams))


	for _, c := range streams {
		ct.ChangeColor(ct.White, true, ct.Black, true)
		fmt.Printf("%s", c.Channel.Name)
		ct.ChangeColor(ct.Green, false, ct.Black, false)
		fmt.Printf(" %s", c.Channel.Game)
		ct.ChangeColor(ct.Red, false, ct.Black, false)
		fmt.Printf(" [")
		ct.ChangeColor(ct.White, false, ct.Black, false)
		fmt.Printf("%d", int64(c.Viewers))
		ct.ChangeColor(ct.Red, false, ct.Black, false)
		fmt.Printf("<-")
		ct.ChangeColor(ct.White, false, ct.Black, false)
		fmt.Printf("%dh", int64(c.Height))
		ct.ChangeColor(ct.Red, false, ct.Black, false)
		fmt.Printf("@")
		ct.ChangeColor(ct.White, false, ct.Black, false)
		fmt.Printf("%f", c.FPS)
		ct.ChangeColor(ct.Red, false, ct.Black, false)
		fmt.Printf("]\n")
		ct.ChangeColor(ct.Cyan, false, ct.Black, false)
		fmt.Printf("  %s\n", c.Channel.Status)
	}
}
