package mementopattern

import "fmt"

type Originator interface {
	Save(tag string) Memento
}

type Dungeon struct {
	name       string
	roleStatus []string
	scenario   string
}

func NewDungeon(name string, roleName string) *Dungeon {
	return &Dungeon{
		name:       name,
		roleStatus: []string{roleName, "Boss down number: 0"},
		scenario:   "Start to raid",
	}
}

func (d *Dungeon) Save(tag string) Memento {
	return NewPlayerRecord(tag, d.roleStatus, d.scenario, d)
}

func (d *Dungeon) SetRoleStatus(roleStatus []string) {
	d.roleStatus = roleStatus
}

func (d *Dungeon) SetScenario(scenario string) {
	d.scenario = scenario
}

func (d *Dungeon) String() string {
	return fmt.Sprintf("In dungeon %s, %s %s", d.name, d.roleStatus[0], d.scenario)
}

type Memento interface {
	Tag() string
	Restore()
}

type playerRecord struct {
	tag        string
	roleStatus []string
	scenario   string
	dungeon    *Dungeon
}

func NewPlayerRecord(tag string, roleStatus []string, scenario string, dungeon *Dungeon) *playerRecord {
	return &playerRecord{
		tag:        tag,
		roleStatus: roleStatus,
		scenario:   scenario,
		dungeon:    dungeon,
	}
}

func (p *playerRecord) Tag() string {
	return p.tag
}

func (p *playerRecord) Restore() {
	p.dungeon.SetRoleStatus(p.roleStatus)
	p.dungeon.SetScenario(p.scenario)
}

type DungeonManager struct {
	archives map[string]Memento
}

func NewDungeonManager() *DungeonManager {
	return &DungeonManager{archives: make(map[string]Memento)}
}

func (d *DungeonManager) Reload(tag string) {
	if achive, ok := d.archives[tag]; ok {
		fmt.Printf("Reload Dungeon %s \n", tag)
		achive.Restore()
	}
}

func (d *DungeonManager) Put(memento Memento) {
	d.archives[memento.Tag()] = memento
}
