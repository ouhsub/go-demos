package state

import "fmt"

type IState interface {
	Approval(*Machine)
	Reject(*Machine)
	GetName() string
}

type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

type leaderApproveState struct{}

func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功")
	m.SetState(GetFinanceApproveState())
}

func (leaderApproveState) GetName() string {
	return "leaderApproveState"
}

func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return leaderApproveState{}
}

type financeApproveState struct{}

func (financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("发出打款操作")
}

func (financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

func (financeApproveState) GetName() string {
	return "FinanceApproveState"
}

func GetFinanceApproveState() IState {
	return financeApproveState{}
}
