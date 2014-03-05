package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

const DigitWidth int = 5

var Zero = []int{
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 1,
	1, 0, 0, 0, 1,
	1, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
}

var One = []int{
	0, 0, 1, 0, 0,
	0, 1, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 1, 1, 1, 0,
}

var Two = []int{
	1, 1, 1, 1, 1,
	0, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 0,
	1, 1, 1, 1, 1,
}

var Three = []int{
	1, 1, 1, 1, 1,
	0, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
	0, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
}

var Four = []int{
	1, 0, 0, 0, 1,
	1, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
	0, 0, 0, 0, 1,
	0, 0, 0, 0, 1,
}

var Five = []int{
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 0,
	1, 1, 1, 1, 1,
	0, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
}

var Six = []int{
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 0,
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
}

var Seven = []int{
	1, 1, 1, 1, 1,
	0, 0, 0, 0, 1,
	0, 0, 0, 0, 1,
	0, 0, 0, 0, 1,
	0, 0, 0, 0, 1,
}

var Eight = []int{
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
}

var Nine = []int{
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 1,
	1, 1, 1, 1, 1,
	0, 0, 0, 0, 1,
	0, 0, 0, 0, 1,
}

var Colon = []int{
	0, 0, 0, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 0, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 0, 0, 0,
}

var Digits = map[rune][]int{
	'0': Zero,
	'1': One,
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	':': Colon,
}

var End = time.Now().Add(20 * time.Minute)

func main() {

	termbox.Init()
	defer termbox.Close()

	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()

	draw()

loop:
	for {
		select {
		case ev := <-events:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break loop
			}
		default:
			draw()
			time.Sleep(10 * time.Millisecond)
		}
	}

}

func draw() {
	// w, h := termbox.Size()

	now := time.Now()
	m := int(End.Sub(now) / time.Minute)
	s := int((End.Sub(now) % time.Minute) / time.Second)
	timeLeft := fmt.Sprintf("%02d:%02d", m, s)

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for i, r := range timeLeft {
		drawDigit((DigitWidth+2)*i, 0, Digits[r])
	}

	termbox.Flush()
}

func drawDigit(x, y int, digit []int) {
	for i, v := range digit {
		char := ' '
		x1 := x + i%DigitWidth
		y1 := y + i/DigitWidth

		if v == 1 {
			char = '▒'
		}

		termbox.SetCell(x1, y1, char, termbox.ColorGreen, termbox.Attribute(termbox.AttrBold))
	}
}
