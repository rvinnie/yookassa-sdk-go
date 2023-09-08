// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Agent type is specified in the agent_type parameter of the items array in the request for creating a receipt,
// if you send data for creating receipts using the Receipt after payment scenario.
type AgentType string

const (
	// Banking payment agent
	BankingPaymentAgent AgentType = "banking_payment_agent"

	// Banking payment subagent
	BankingPaymentSubagent AgentType = "banking_payment_subagent"

	// Payment agent
	PaymentAgent AgentType = "payment_agent"

	// Payment subagent
	PaymentSubagent AgentType = "payment_subagent"

	// Attorney
	Attorney AgentType = "attorney"

	// Broker
	Commissioner AgentType = "commissioner"

	// Agent
	Agent AgentType = "agent"
)
