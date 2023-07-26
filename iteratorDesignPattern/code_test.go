package iteratordesignpattern

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	playerBackpack := NewPlayerBackpack(1, "Saltedfisher")
	playerBackpack.AddProp(NewProp("Frost Sorrow Sword", 999),
		NewProp("Invincible Monte", 888),
		NewProp("MusicBox", 666),
	)

	fmt.Printf("%s's props info:\n", playerBackpack.player.name)
	backpackIterator := playerBackpack.CreateIterator()
	for backpackIterator.HasMore() {
		backpack := backpackIterator.Next()
		fmt.Println(backpack.Desc())
	}
}
