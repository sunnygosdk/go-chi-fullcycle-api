package contract

// ContractRepository is an interface for contract repository.
type ContractRepository interface {
	// Create creates a new contract.
	Create(contract *contract) error
}
