package chainofresponsibility

import (
	"fmt"
)

// 办理入住
type CheckInProgress interface {
	SetNextProgress(processor CheckInProgress)
	Execute(tenant *tenant)
}

type tenant struct {
	Name              string // 姓名
	Authenticate      bool   // 验证身份信息
	ContractSigned    bool   // 签约
	PaymentDone       bool   // 缴押金
	IsCompleteCheckIn bool   // 完成办理入住
}

type baseCheckinProgress struct {
	nextProgress CheckInProgress
}

func (b *baseCheckinProgress) SetNextProgress(processor CheckInProgress) {
	b.nextProgress = processor
}

func (b *baseCheckinProgress) Execute(tenant *tenant) {
	if b.nextProgress != nil {
		b.nextProgress.Execute(tenant)
	}
}

type authenticateProgress struct {
	baseCheckinProgress
}

func (a *authenticateProgress) Execute(tenant *tenant) {
	if !tenant.Authenticate {
		fmt.Printf("租客:%s, 正在校验租客身份信息;\n", tenant.Name)
		tenant.Authenticate = true
	}
	a.baseCheckinProgress.Execute(tenant)
}

type signContractProgress struct {
	baseCheckinProgress
}

func (s *signContractProgress) Execute(tenant *tenant) {
	if !tenant.Authenticate {
		fmt.Printf("租客%s身份校验不通过, 无法签署;\n", tenant.Name)
		return
	}

	if !tenant.ContractSigned {
		fmt.Printf("租客:%s, 正在签署合同;\n", tenant.Name)
		tenant.ContractSigned = true
	}
	s.baseCheckinProgress.Execute(tenant)
}

type payProgress struct {
	baseCheckinProgress
}

func (p *payProgress) Execute(tenant *tenant) {
	if !tenant.ContractSigned {
		fmt.Printf("租客:%s, 未签署合同无法付款;\n", tenant.Name)
		return
	}

	if !tenant.PaymentDone {
		fmt.Printf("租客:%s, 正在支付押金;\n", tenant.Name)
		tenant.PaymentDone = true
	}
	p.baseCheckinProgress.Execute(tenant)
}

type completeCheckInProgress struct {
	baseCheckinProgress
}

func (c *completeCheckInProgress) Execute(tenant *tenant) {
	if !tenant.Authenticate ||
		!tenant.ContractSigned ||
		!tenant.PaymentDone {
		fmt.Printf("租客%s未完成相关手续, 无法办理入住;\n", tenant.Name)
		return
	}
	tenant.IsCompleteCheckIn = true
	fmt.Printf("租客:%s, 成功办理入住;\n", tenant.Name)
}
