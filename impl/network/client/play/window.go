package play

import (
	"github.com/minelc/go-server/api/network"
)

/*
Called on window close
*/
type PacketPlayInCloseWindow struct {
	Id byte
}

func (p *PacketPlayInCloseWindow) UUID() int32 {
	return 13
}

func (p *PacketPlayInCloseWindow) Pull(reader network.Buffer) {
	p.Id = reader.PullByt()
}

func (p *PacketPlayInCloseWindow) Handle(c *network.Connection) {

}

/*
Called on window close
*/
type PacketPlayInWindowClick struct {
	Id     byte
	Slot   int16
	Button byte
	D      int16
	Shift  byte
	// Todo Add itemstack item
}

func (p *PacketPlayInWindowClick) UUID() int32 {
	return 14
}

func (p *PacketPlayInWindowClick) Pull(reader network.Buffer) {
	p.Id = reader.PullByt()
	p.Slot = reader.PullI16()
	p.Button = reader.PullByt()
	p.D = reader.PullI16()
	p.Shift = reader.PullByt()
	// Todo Parse to itemstack
}

func (p *PacketPlayInWindowClick) Handle(c *network.Connection) {

}

/*
Called on trade
*/
type PacketPlayInTransaction struct {
	Id   byte
	Slot int16
	C    bool
}

func (p *PacketPlayInTransaction) UUID() int32 {
	return 15
}

func (p *PacketPlayInTransaction) Pull(reader network.Buffer) {
	p.Id = reader.PullByt()
	p.Slot = reader.PullI16()
	p.C = reader.PullBit()
}
func (p *PacketPlayInTransaction) Handle(c *network.Connection) {

}

/*
Called on interact with slot on gamemode creative
*/
type PacketPlayInSetCreativeSlot struct {
	Slot int16
	// TODO: Add itemstack
}

func (p *PacketPlayInSetCreativeSlot) UUID() int32 {
	return 16
}

func (p *PacketPlayInSetCreativeSlot) Pull(reader network.Buffer) {
	p.Slot = reader.PullI16()
	// TODO: Add itemstack
}
func (p *PacketPlayInSetCreativeSlot) Handle(c *network.Connection) {

}

/*
Called on enchant item
*/
type PacketPlayInEnchantItem struct {
	Id byte
	A  byte
}

func (p *PacketPlayInEnchantItem) UUID() int32 {
	return 17
}

func (p *PacketPlayInEnchantItem) Pull(reader network.Buffer) {
	p.Id = reader.PullByt()
	p.A = reader.PullByt()
}

func (p *PacketPlayInEnchantItem) Handle(c *network.Connection) {

}

/*
On change item hand slot
*/
type PacketPlayInHeldItemSlot struct {
	HandIndex int16
}

func (p *PacketPlayInHeldItemSlot) UUID() int32 {
	return 9
}

func (p *PacketPlayInHeldItemSlot) Pull(reader network.Buffer) {
	p.HandIndex = reader.PullI16()
}

func (p *PacketPlayInHeldItemSlot) Handle(c *network.Connection) {

}
