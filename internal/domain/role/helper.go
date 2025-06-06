package role

// roleType is a type for role type.
type roleType string

const (
	// RoleTypeAdmin is the role type for admin.
	RoleTypeAdmin roleType = "admin"

	// RoleTypeStoreManager is the role type for store manager.
	RoleTypeStoreManager roleType = "store_manager"

	// RoleTypeDepartmentManager is the role type for department manager.
	RoleTypeDepartmentManager roleType = "department_manager"

	// RoleTypeStockManager is the role type for stock manager.
	RoleTypeStockManager roleType = "stock_manager"

	// RoleTypeSeller is the role type for seller.
	RoleTypeSeller roleType = "seller"

	// RoleTypeHumanResources is the role type for human resources.
	RoleTypeHumanResources roleType = "human_resources"
)

// RoleTypeList is a list of role types.
var RoleTypeList = []roleType{
	RoleTypeAdmin,
	RoleTypeStoreManager,
	RoleTypeDepartmentManager,
	RoleTypeStockManager,
	RoleTypeSeller,
	RoleTypeHumanResources,
}

// String returns the string representation of the role type.
func (r roleType) String() string {
	return string(r)
}
