package ents

type EntityMeta struct {
	OnFire    bool
	Crouched  bool
	Sprinting bool
	Eating    bool
	Invisible bool
}

type MetadataType byte

const (
	Byte   MetadataType = 0
	Short  MetadataType = 1
	Int    MetadataType = 2
	Float  MetadataType = 3
	String MetadataType = 4
	Slot   MetadataType = 5
	Pos    MetadataType = 6
	Head   MetadataType = 7

	True  byte = 0x01
	False byte = 0x00

	// Entity
	EntityBitMask byte = 0 // 0x01 byte = On Fire | 0x02 byte = Crouched | 0x08 byte = Sprinting | 0x10 byte = Eating/Drinking/Blocking | 0x20 byte = Invisible
	EntityShort   byte = 1

	// EntityLiving
	NameTag         byte = 2
	ShowNameTag     byte = 3
	Health          byte = 6
	PotionColor     byte = 7
	IsPotionAmbient byte = 8
	NumberArrows    byte = 9
	HasIA           byte = 15

	// Ageable
	Age byte = 12

	// ArmorStand
	StandBitMask  byte = 10 // 0x01 byte = Small | 0x02 byte = Has Gravity | 0x04 byte = Has Arms | 0x08 byte = Remove BasePlate | 0x16 byte = Marker
	StandHead     byte = 11
	StandPos      byte = 12
	StandLeftArm  byte = 13
	StandRightArm byte = 14
	StandLeftLeg  byte = 15
	StandRightLeg byte = 16

	// Human
	HumanSkin       byte = 10
	HumanCape       byte = 16 // 0x02 byte = Hide cape
	HumanAbsorption byte = 17
	HumanScore      byte = 18
)
