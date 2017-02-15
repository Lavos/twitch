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
	statusStyle tcell.Style
	viewerStyle tcell.Style
)

const (

)

type Configuration struct {
	Username string
	ClientID string
}

func ticker(count int) {
	s.SetContent(count % 20, 0, tcell.RuneDiamond, nil, tickerStyle)
	s.Show()
}

func draw() {
	s.Clear()

	streams, err := t.Online()

	if err != nil {
		for x, r := range err.Error() {
			s.SetContent(x + 21, 0, r, nil, tickerStyle)
		}

		s.Show()
		return
	}

	var viewers_column_end, name_column_end, game_column_end int

	// get column widths
	for _, b := range streams {
		if len(b.Channel.Name) > name_column_end {
			name_column_end = len(b.Channel.Name)
		}

		characters := len(fmt.Sprintf("%7.f", b.Viewers))

		if viewers_column_end < characters {
			viewers_column_end = characters
		}

		if len(b.Channel.Game) > game_column_end {
			game_column_end = len(b.Channel.Game)
		}
	}

	// draw columns
	for y, b := range streams {
		var x int
		var r rune

		for x, r = range fmt.Sprintf("%7.f", b.Viewers) {
			s.SetContent(x, y + 1, r, nil, viewerStyle)
		}

		for x, r = range b.Channel.Name {
			s.SetContent(x + viewers_column_end + 1, y + 1, r, nil, nameStyle)
		}

		for x, r = range b.Channel.Game {
			s.SetContent(x + viewers_column_end + name_column_end + 2, y + 1, r, nil, gameStyle)
		}

		for x, r = range b.Channel.Status {
			s.SetContent(x + viewers_column_end + game_column_end + name_column_end + 3, y + 1, r, nil, statusStyle)
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

	tickerStyle = tickerStyle.Foreground(tcell.ColorRed)
	nameStyle = nameStyle.Foreground(tcell.ColorWhite)
	gameStyle = gameStyle.Foreground(tcell.ColorGreen)
	viewerStyle = viewerStyle.Foreground(tcell.ColorFuchsia)
	statusStyle = statusStyle.Foreground(tcell.ColorLightSlateGray)

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
				case tcell.KeyEscape:
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

	var counter int

	for {
		select {
		case <-tc.C:
			if counter % 20 == 0 {
				draw()
			}

			ticker(counter)
			counter++

		case <-quit:
			s.Fini()
			return
		}
	}
}
