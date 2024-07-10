package impl

import "time"

type Mspt struct {
	max               int64
	promedium         int64
	min               int64
	elapseTicks       int64
	nextTwentySeconds int64
}

func (m Mspt) GetPromedium() float32 {
	return float32(m.promedium / m.elapseTicks)
}

func (m Mspt) GetMax() int64 {
	return m.max
}

func (m Mspt) GetMin() int64 {
	return m.min
}

func (m *Mspt) Handle(startTime int64) int64 {

	endTime := time.Now().UnixMilli()
	mspt := endTime - startTime

	// Restart mspt data every 20s
	if endTime >= m.nextTwentySeconds {
		m.nextTwentySeconds = endTime + 20_000
		m.min = 0
		m.elapseTicks = 0
		m.promedium = 0
		m.max = 0
		m.min = 0
	}

	m.promedium += mspt

	if mspt < m.min {
		m.min = mspt
	} else if mspt > m.max {
		m.max = mspt
	}
	return mspt
}
