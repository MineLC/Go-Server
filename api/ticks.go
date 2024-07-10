package api

type Mspt interface {
	GetPromedium() float32
	GetMax() int64
	GetMin() int64
}
