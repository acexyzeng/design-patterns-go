package chainofresponsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	checkinProgress := BuildCheckinProgress()
	tenant := &tenant{
		Name:              "test01",
		Authenticate:      false,
		ContractSigned:    false,
		PaymentDone:       false,
		IsCompleteCheckIn: false,
	}
	checkinProgress.Execute(tenant)
}

func BuildCheckinProgress() CheckInProgress {
	completeNode := &completeCheckInProgress{}

	payNode := &payProgress{}
	payNode.SetNextProgress(completeNode)

	contractSignNode := &signContractProgress{}
	contractSignNode.SetNextProgress(payNode)

	authNode := &authenticateProgress{}
	authNode.SetNextProgress(contractSignNode)
	return authNode
}
