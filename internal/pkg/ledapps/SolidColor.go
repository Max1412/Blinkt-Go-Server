package ledapps

import . "github.com/alexellis/blinkt_go"

type SolidColor struct {
	Brightness float64
	Step       int
	MBlinkt    Blinkt
}

func (s *SolidColor) Setup() {
	s.Brightness = 0.5
	s.MBlinkt = NewBlinkt(s.Brightness)
	s.MBlinkt.SetClearOnExit(true)
	s.MBlinkt.Setup()
	Delay(100)
	s.Step = 0
}

func (s *SolidColor) Loop() {
	s.Step = s.Step % 3
	switch s.Step {
	case 0:
		s.MBlinkt.SetAll(128, 0, 0)
	case 1:
		s.MBlinkt.SetAll(0, 128, 0)
	case 2:
		s.MBlinkt.SetAll(0, 0, 128)
	}
	s.Step++
	s.MBlinkt.Show()
	Delay(500)
}

func (s *SolidColor) Cleanup() {
	s.MBlinkt.Clear()
	s.MBlinkt.Show()
}
