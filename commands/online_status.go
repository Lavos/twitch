package main

import (
	"fmt"
	"os"
	"time"
	"github.com/Lavos/twitch"
	"github.com/gdamore/tcell"
	"github.com/kelseyhightower/envconfig"
)

var (
	c Configuration
	s tcell.Screen
	t *twitch.TwitchClient

	tickerStyle tcell.Style
	nameStyle tcell.Style
	gameStyle tcell.Style
)

const (

)

type Configuration struct {
	Username string
	ClientID string
}

func ticker(count int) {
	s.SetContent(count, 0, tcell.RuneDiamond, nil, tickerStyle)
	s.Show()
}

func draw() {
	s.Clear()

	streams, err := t.Online()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	var y int = 1

	for row, b := range streams {
		var x int
		var r rune

		for x, r = range b.Channel.Name {
			s.SetContent(x, y + row, r, nil, nameStyle)
		}

		x += 2

		for gx, r := range b.Channel.Game {
			s.SetContent(x + gx, y + row, r, nil, gameStyle)
		}
	}

	s.Show()
}

func main() {
	var err error
	envconfig.MustProcess("TWITCH", &c)

	if c.Username == "" || c.ClientID == "" {
		fmt.Fprintf(os.Stderr, "Missing required ENV variables: %#v", c)
		os.Exit(1)
	}

	// init tcell
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err = tcell.NewScreen()

	tickerStyle = gameStyle.Foreground(tcell.ColorRed)
	nameStyle = nameStyle.Foreground(tcell.ColorWhite)
	gameStyle = gameStyle.Foreground(tcell.ColorGreen)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err = s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// create Twitch Client
	t = &twitch.TwitchClient{c.Username, c.ClientID}

	// create ticker
	tc := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			ev := s.PollEvent()

			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return

				case tcell.KeyCtrlL:
					s.Sync()
				}

			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	draw()

	var counter int

	for {
		select {
		case <-tc.C:
			if counter > 0 && counter % 20 == 0 {
				draw()
				counter = 0
			} else {
				ticker(counter)
				counter++
			}

		case <-quit:
			s.Fini()
			return
		}
	}
}
