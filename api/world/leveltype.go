package world

type LevelType int

const (
	DEFAULT LevelType = iota
	FLAT
	LARGEBIOMES
	AMPLIFIED
	CUSTOMIZED
	BUFFET
	DEFAULT11
)

func (l LevelType) String() string {
	switch l {
	case DEFAULT:
		return "default"
	case FLAT:
		return "flat"
	case LARGEBIOMES:
		return "largeBiomes"
	case AMPLIFIED:
		return "amplified"
	case CUSTOMIZED:
		return "customized"
	case BUFFET:
		return "buffet"
	case DEFAULT11:
		return "default_1_1"
	}
	return "default"
}
