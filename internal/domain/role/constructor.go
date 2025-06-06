package role

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// NewRole creates a new role instance with the provided name and type.
// It initializes the role with the given name and type,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the role before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - Name: Name of the role.
//   - Type: Type of the role.
//
// Returns:
//   - *role: A pointer to the newly created and validated role.
//   - error: An error if the role validation fails.
func NewRole(Name string, Type roleType) (*role, error) {
	role := &role{
		ID:        entity.NewID(),
		Name:      Name,
		Type:      Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	role, err := role.validate()
	if err != nil {
		return nil, err
	}

	return role, nil
}

// validate validates the role instance.
// It checks if the role ID is required and valid,
// and if the role name and type are required.
//
// Parameters:
//   - r: The role instance to validate.
//
// Returns:
//   - *role: A pointer to the validated role instance.
//   - error: An error if the role validation fails.
func (r *role) validate() (*role, error) {
	if r.ID.String() == "" {
		return nil, ErrRoleIDisRequired
	}

	_, err := entity.ParseID(r.ID.String())
	if err != nil {
		return nil, ErrRoleInvalidID
	}

	if r.Name == "" {
		return nil, ErrRoleNameRequired
	}

	if r.Type == "" {
		return nil, ErrRoleTypeRequired
	}

	if !isValidRoleType(r.Type) {
		return nil, ErrRoleInvalidType
	}

	return r, nil
}

// isValidRoleType checks if the role type is valid.
// It checks if the role type is required and valid.
//
// Parameters:
//   - roleType: The role type to validate.
//
// Returns:
//   - bool: True if the role type is valid, false otherwise.
func isValidRoleType(roleType roleType) bool {
	for _, role := range RoleTypeList {
		if role == roleType {
			return true
		}
	}
	return false
}
