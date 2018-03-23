package ledapps

import (
	"time"

	. "github.com/alexellis/blinkt_go"
	"github.com/lucasb-eyer/go-colorful"
)

type WakeUp struct {
	StartBrightness float64
	EndBrightness   float64
	Duration        time.Duration
	StartTime       time.Time
	MBlinkt         Blinkt
	Step            int
}

func (s *WakeUp) Setup() {
	s.Step = 0
	s.StartBrightness = 0.1
	s.EndBrightness = 1.0
	s.Duration, _ = time.ParseDuration("30m")
	s.StartTime = time.Now()
	s.MBlinkt = NewBlinkt(s.StartBrightness)
	s.MBlinkt.SetClearOnExit(true)
	s.MBlinkt.Setup()
	Delay(20)
}

func (s *WakeUp) Loop() {
	// TODO maybe prevent updating the brightness to often
	elapsedTime := time.Since(s.StartTime)

	c1, _ := colorful.Hex("#cc0000")
	c2, _ := colorful.Hex("#cc3000")
	c3, _ := colorful.Hex("#ffffff")

	if elapsedTime.Seconds() <= s.Duration.Seconds() {

		// linear interpolation
		timeT := elapsedTime.Seconds() / s.Duration.Seconds()
		currentBrightness := s.StartBrightness + timeT*(s.EndBrightness-s.StartBrightness)

		s.MBlinkt.SetBrightness(currentBrightness)
		c, _ := colorful.Hex("#000000")
		if timeT < 0.75 {
			c = c1.BlendHcl(c2, timeT*1.33)
		} else {
			c = c2.BlendHcl(c3, (timeT-0.75)*4.0)
		}
		r, g, b := c.Clamped().RGB255()
		//fmt.Printf("RGB: %d, %d, %d, Brightness: %f\n", r, g, b, currentBrightness)

		s.MBlinkt.SetAll(int(r), int(g), int(b))

		s.MBlinkt.Show()
	} else {
		// show its done
		s.Step = s.Step % 2
		switch s.Step {
		case 0:
			s.MBlinkt.SetAll(255, 0, 0)
		case 1:
			s.MBlinkt.SetAll(0, 255, 0)
		}
		s.MBlinkt.SetBrightness(s.EndBrightness)
		s.MBlinkt.Show()
	}

}

func (s *WakeUp) Cleanup() {
	s.MBlinkt.Clear()
	s.MBlinkt.Show()
}
