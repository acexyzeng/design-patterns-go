package mediator

import "fmt"

type Developer interface {
	Review()
	DevelopmentCompleted()
}

type WorkMediator interface {
	CanDevelop(developer Developer) bool
	NotifyWaitingRequirement()
}

type productRequirement struct {
	name         string
	workMediator WorkMediator
}

func NewRequirement(name string, mediator WorkMediator) *productRequirement {
	return &productRequirement{
		name:         name,
		workMediator: mediator,
	}
}

func (p *productRequirement) Review() {
	if !p.workMediator.CanDevelop(p) {
		fmt.Printf("developer are busy, product %s is waiting for review \n", p.name)
		return
	}
	fmt.Printf("productRequirement %s, can be reviewed \n", p.name)
}

func (p *productRequirement) DevelopmentCompleted() {
	fmt.Printf("productRequirement %s is completed \n", p.name)
	p.workMediator.NotifyWaitingRequirement()
}

type DevTeam struct {
	hasFreeTime bool
	waitingTask []Developer
}

func (d *DevTeam) CanDevelop(dev Developer) bool {
	if d.hasFreeTime {
		d.hasFreeTime = false
		return true
	}
	d.waitingTask = append(d.waitingTask, dev)
	return false
}

func (d *DevTeam) NotifyWaitingRequirement() {
	if !d.hasFreeTime {
		d.hasFreeTime = true
	}
	if len(d.waitingTask) > 0 {
		first := d.waitingTask[0]
		d.waitingTask = d.waitingTask[1:]
		first.Review()
	}
}
