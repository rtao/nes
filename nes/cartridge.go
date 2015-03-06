package nes

import "log"

const (
	MirrorHorizontal = 0
	MirrorVertical   = 1
	MirrorQuad       = 2
)

type Cartridge struct {
	PRG     []byte // PRG-ROM banks
	CHR     []byte // CHR-ROM banks
	SRAM    []byte // Save RAM
	Mapper  int    // mapper type
	Mirror  int    // mirroring mode
	Battery bool   // battery present
}

func (c *Cartridge) Read(address uint16) byte {
	switch {
	case address < 0x2000:
		return c.CHR[address]
	case address >= 0x8000:
		index := (int(address) - 0x8000) % len(c.PRG)
		return c.PRG[index]
	case address >= 0x6000:
		index := int(address) - 0x6000
		return c.SRAM[index]
	default:
		log.Fatalf("unhandled cartridge read at address: 0x%04X", address)
	}
	return 0
}

func (c *Cartridge) Write(address uint16, value byte) {
	switch {
	case address < 0x2000:
		c.CHR[address] = value
	case address >= 0x8000:
		break
	case address >= 0x6000:
		index := int(address) - 0x6000
		c.SRAM[index] = value
	default:
		log.Fatalf("unhandled cartridge write at address: 0x%04X", address)
	}
}
