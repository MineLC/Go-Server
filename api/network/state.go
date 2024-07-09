package network

type PacketState int32

const (
	SHAKE PacketState = iota
	STATUS
	LOGIN
	PLAY
	INVALID
)

func StateValueOf(i int) PacketState {
	switch i {
	case 0:
		return SHAKE
	case 1:
		return STATUS
	case 2:
		return LOGIN
	case 3:
		return PLAY
	}
	return INVALID
}
