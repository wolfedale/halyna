package mrad

import "math"

var (
	BC          = 0.45
	Speed       = 800
	Win308Wind0 = map[int]int{
		300: 15,
		250: 10,
		240: 9,
		230: 9,
		220: 8,
		210: 7,
		200: 6,
		190: 5,
		180: 5,
		170: 4,
		160: 3,
		150: 3,
		140: 2,
		130: 1,
		120: 1,
		100: 0,
	}
)

type Clicker interface {
	Clicks(missedInCentimeters int) float32
}

type Mrad struct {
	distanceInMeters   int
	windSpeedKmProHour int
	mrad               float64
}

func New(distanceInMeters int, windSpeedKmProHour int, mrad float64) *Mrad {
	if mrad == 0 {
		mrad = 0.1
	}

	return &Mrad{
		distanceInMeters:   distanceInMeters,
		windSpeedKmProHour: windSpeedKmProHour,
		mrad:               mrad,
	}
}

func (m *Mrad) Clicks(missedInCentimeters int) int {
	return customRound(float64(missedInCentimeters) / m.correction(float64(m.distanceInMeters)))
}

func (m *Mrad) correction(distance float64) float64 {
	return distance / 100.0
}

func customRound(value float64) int {
	if value-math.Floor(value) >= 0.5 {
		return int(math.Ceil(value))
	}

	return int(math.Floor(value))
}
