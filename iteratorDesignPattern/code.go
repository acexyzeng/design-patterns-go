package iteratordesignpattern

import "fmt"

type Backpack interface {
	Desc() string // describe prop or player info
}

type Player struct {
	name string
}

func NewPlayer(name string) *Player {
	return &Player{name: name}
}

func (p *Player) Desc() string {
	return fmt.Sprintf("Player name: %s \n", p.name)
}

type Prop struct {
	name  string
	value int
}

func NewProp(name string, value int) *Prop {
	return &Prop{name: name, value: value}
}

func (p *Prop) Desc() string {
	return fmt.Sprintf("Prop: %s, Value: %d", p.name, p.value)
}

type Iterator interface {
	Next() Backpack
	HasMore() bool
}

type PlayerBackpack struct {
	number int
	player *Player
	props  []*Prop
}

type backpackIterator struct {
	playerBackpack *PlayerBackpack
	index          int
}

func (b *backpackIterator) Next() Backpack {
	if b.index == -1 {
		b.index++
		return b.playerBackpack.player
	}
	prop := b.playerBackpack.props[b.index]
	b.index++
	return prop
}

func (b *backpackIterator) HasMore() bool {
	return b.index < len(b.playerBackpack.props)
}

type Iterable interface {
	CreateIterator() Iterator
}

func NewPlayerBackpack(number int, playerName string) *PlayerBackpack {
	return &PlayerBackpack{
		number: number,
		player: NewPlayer(playerName),
	}
}

// Create Iterator
func (p *PlayerBackpack) CreateIterator() Iterator {
	return &backpackIterator{
		playerBackpack: p,
		index:          -1,
	}
}

func (p *PlayerBackpack) Number() int {
	return p.number
}

func (p *PlayerBackpack) AddProp(props ...*Prop) {
	p.props = append(p.props, props...)
}
