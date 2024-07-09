package data

type CREATURE int32

const (
	MOB         CREATURE = 48
	MONSTER     CREATURE = 49
	CREEPER     CREATURE = 50
	SKELETON    CREATURE = 51
	SPIDER      CREATURE = 52
	GIANT       CREATURE = 53
	ZOMBIE      CREATURE = 54
	SLIME       CREATURE = 55
	GHAST       CREATURE = 56
	PIGZOMBIE   CREATURE = 57
	ENDERMAN    CREATURE = 58
	CAVESPIDER  CREATURE = 59
	SILVERFISH  CREATURE = 60
	BLAZE       CREATURE = 61
	LAVASLIME   CREATURE = 62
	ENDERDRAGON CREATURE = 63
	WITHERBOSS  CREATURE = 64
	BAT         CREATURE = 65
	WITCH       CREATURE = 66
	ENDERMITE   CREATURE = 67
	GUARDIAN    CREATURE = 68
	// Pasive creatures
	PIG           CREATURE = 90
	SHEEP         CREATURE = 91
	COW           CREATURE = 92
	CHICKEN       CREATURE = 93
	SQUID         CREATURE = 94
	WOLF          CREATURE = 95
	MUSHROOMCOW   CREATURE = 96
	SNOWMAN       CREATURE = 97
	OZELOT        CREATURE = 98
	VILLAGERGOLEM CREATURE = 99
	ENTITYHORSE   CREATURE = 100
	RABBIT        CREATURE = 101
	VILLAGER      CREATURE = 120
)

func GetCreature(name string) CREATURE {
	switch name {
	case "MOB":
		return 48
	case "MONSTER":
		return 49
	case "CREEPER":
		return 50
	case "SKELETON":
		return 51
	case "SPIDER":
		return 52
	case "GIANT":
		return 53
	case "ZOMBIE":
		return 54
	case "SLIME":
		return 55
	case "GHAST":
		return 56
	case "PIGZOMBIE":
		return 57
	case "ENDERMAN":
		return 58
	case "CAVESPIDER":
		return 59
	case "SILVERFISH":
		return 60
	case "BLAZE":
		return 61
	case "LAVASLIME":
		return 62
	case "ENDERDRAGON":
		return 63
	case "WITHERBOSS":
		return 64
	case "BAT":
		return 65
	case "WITCH":
		return 66
	case "ENDERMITE":
		return 67
	case "GUARDIAN":
		return 68
	case "PIG":
		return 90
	case "SHEEP":
		return 91
	case "COW":
		return 92
	case "CHICKEN":
		return 93
	case "SQUID":
		return 94
	case "WOLF":
		return 95
	case "MUSHROOMCOW":
		return 96
	case "SNOWMAN":
		return 97
	case "OZELOT":
		return 98
	case "VILLAGERGOLEM":
		return 99
	case "ENTITYHORSE":
		return 100
	case "RABBIT":
		return 101
	case "VILLAGER":
		return 120
	}
	return MOB
}
