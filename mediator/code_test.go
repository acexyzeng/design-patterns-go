package mediator

import "testing"

func TestMediator(t *testing.T) {
	workMediator := &DevTeam{hasFreeTime: true}

	deviceRequirement := NewRequirement("deviceAlarm", workMediator)

	paymentRequirement := NewRequirement("payment platform", workMediator)

	deviceRequirement.Review()
	paymentRequirement.Review()
	deviceRequirement.DevelopmentCompleted()
	paymentRequirement.DevelopmentCompleted()
}
