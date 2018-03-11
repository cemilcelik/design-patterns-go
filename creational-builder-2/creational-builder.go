package main

import (
	"fmt"
)

type PC struct {
	cpu string
	ram int
	gpu string
	hdd int
	os  string
}

type PCBuilder interface {
	setCPU() PCBuilder
	setRAM() PCBuilder
	setGPU() PCBuilder
	setHDD() PCBuilder
	setOS() PCBuilder
	getPC() PC
}

type HomePCBuilder struct {
	pc PC
}

func (h *HomePCBuilder) setCPU() PCBuilder {
	h.pc.cpu = "i3"
	return h
}
func (h *HomePCBuilder) setRAM() PCBuilder {
	h.pc.ram = 4
	return h
}
func (h *HomePCBuilder) setGPU() PCBuilder {
	h.pc.gpu = "Intel Graphic Card"
	return h
}
func (h *HomePCBuilder) setHDD() PCBuilder {
	h.pc.hdd = 500
	return h
}
func (h *HomePCBuilder) setOS() PCBuilder {
	h.pc.os = "Windows 7 Home"
	return h
}
func (h *HomePCBuilder) getPC() PC {
	return h.pc
}

type GamePCBuilder struct {
	pc PC
}

func (g *GamePCBuilder) setCPU() PCBuilder {
	g.pc.cpu = "i7"
	return g
}
func (g *GamePCBuilder) setRAM() PCBuilder {
	g.pc.ram = 4
	return g
}
func (g *GamePCBuilder) setGPU() PCBuilder {
	g.pc.gpu = "AMD Radeon X80"
	return g
}
func (g *GamePCBuilder) setHDD() PCBuilder {
	g.pc.hdd = 500
	return g
}
func (g *GamePCBuilder) setOS() PCBuilder {
	g.pc.os = "Windows 7 Ultimate"
	return g
}
func (g *GamePCBuilder) getPC() PC {
	return g.pc
}

func (g *GamePCBuilder) build() PC {
	return PC{}
}

type Manufacturer struct {
	builder PCBuilder
}

func (m *Manufacturer) setBuilder(builder PCBuilder) {
	m.builder = builder
}

func (m *Manufacturer) construct() {
	m.builder.setCPU().setRAM().setGPU().setHDD().setOS()
}

func main() {
	manufacturer := Manufacturer{}

	homePCBuilder := &HomePCBuilder{}
	manufacturer.setBuilder(homePCBuilder)
	manufacturer.construct()
	homePC := homePCBuilder.getPC()

	gamePCBuilder := &GamePCBuilder{}
	manufacturer.setBuilder(gamePCBuilder)
	manufacturer.construct()
	gamePC := gamePCBuilder.getPC()

	fmt.Println(homePC) // {i3 4 Intel Graphic Card 500 Windows 7 Home}
	fmt.Println(gamePC) // {i7 4 AMD Radeon X80 500 Windows 7 Ultimate}
}
