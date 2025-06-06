package role

// RoleRepository is an interface for role repository.
type RoleRepository interface {
	// Create creates a new role.
	Create(role *role) error
}
