package ledApps

import . "github.com/alexellis/blinkt_go"

type Progress struct {
	Brightness float64
	MBlinkt    Blinkt
	r, g, b    int
}

func (s *Progress) Setup() {
	s.Brightness = 0.5
	s.MBlinkt = NewBlinkt(s.Brightness)
	s.MBlinkt.SetClearOnExit(true)
	s.MBlinkt.Setup()
	Delay(100)
	s.r = 150
	s.g = 0
	s.b = 0
}

func (s *Progress) Loop() {
	for pixel := 0; pixel < 8; pixel++ {
		s.MBlinkt.Clear()
		s.MBlinkt.SetPixel(pixel, s.r, s.g, s.b)
		s.MBlinkt.Show()
		Delay(100)
	}
	for pixel := 7; pixel > 0; pixel-- {
		s.MBlinkt.Clear()
		s.MBlinkt.SetPixel(pixel, s.r, s.g, s.b)
		s.MBlinkt.Show()
		Delay(100)
	}
}

func (s *Progress) Cleanup() {
	s.MBlinkt.Clear()
	s.MBlinkt.Show()
}
