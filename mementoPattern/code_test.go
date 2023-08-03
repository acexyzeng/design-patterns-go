package mementopattern

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	dungeonManger := NewDungeonManager()

	dungeon := NewDungeon("ICC", "Rogue")
	fmt.Println(dungeon)

	dungeonManger.Put(dungeon.Save("Boss No.1"))
	dungeon.SetRoleStatus([]string{"Rogue", "failed"})
	dungeon.SetScenario("Boss 1 failed")
	fmt.Println(dungeon)

	dungeonManger.Reload("Boss No.1")
	fmt.Println(dungeon)
}
