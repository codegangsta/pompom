package main

import (
	"github.com/codegangsta/cli"
	"github.com/nsf/termbox-go"

	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

var (
	Duration       = 25 * time.Minute
	ExitOnComplete bool
	Label          string
	Current        time.Duration
	Paused         bool
)

func main() {
	app := cli.NewApp()
	app.Name = "pompom"
	app.Usage = "A simple pomodoro timer for the command line"
	app.Action = mainAction
	app.Flags = []cli.Flag{
		cli.IntFlag{Name: "duration,d", Value: 25, Usage: "Duration in minutes"},
		cli.BoolFlag{Name: "e", Usage: "Exit when the timer finishes"},
	}

	app.Run(os.Args)
}

func mainAction(c *cli.Context) {
	// Duration flag
	Duration = time.Duration(c.GlobalInt("duration")) * time.Minute

	// Exit On Complete flag
	ExitOnComplete = c.GlobalBool("e")

	termbox.Init()
	defer termbox.Close()

	Label = strings.Join(c.Args(), " ")
	ticker := time.NewTicker(1 * time.Second)

	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

loop:
	for {
		select {
		case ev := <-events:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeySpace {
				Paused = !Paused
			}

		case <-ticker.C:
			if !Paused {
				Current += time.Second
			}

			if Current > Duration && ExitOnComplete {
				break loop
			}

		default:
			if Paused {
				draw(Current, "[Paused] "+Label)
			} else {
				draw(Current, Label)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func draw(current time.Duration, label string) {
	w, h := termbox.Size()

	t := time.Duration(math.Max(0, float64(Duration-current)))
	timeLeft := fmt.Sprintf("%02d:%02d", (t / time.Minute), ((t % time.Minute) / time.Second))
	color := termbox.ColorGreen

	if t <= 5*time.Minute {
		color = termbox.ColorRed
	} else if t <= 10*time.Minute {
		color = termbox.ColorYellow
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// draw digits
	cw := DigitWidth + 1
	for i, r := range timeLeft {
		x := w/2 + cw*i - cw*len(timeLeft)/2
		y := h/2 - DigitWidth/2 - 2
		drawDigit(x, y, Digits[r], color)
	}

	// draw label
	for i, c := range label {
		x := w/2 + i - len(label)/2
		y := h/2 + 2
		termbox.SetCell(x, y, c, color, 0)
	}

	termbox.Flush()
}

func drawDigit(x, y int, digit []int, color termbox.Attribute) {
	for i, v := range digit {
		x1 := x + i%DigitWidth
		y1 := y + i/DigitWidth

		if v == 1 {
			termbox.SetCell(x1, y1, '█', color, color)
		} else {
			termbox.SetCell(x1, y1, ' ', color, 0)
		}

	}
}
