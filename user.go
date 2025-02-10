package roles

type UserRole string

const (
	RoleUserSuperAdmin UserRole = "super_admin"
	RoleUserAdmin      UserRole = "admin"
	RoleUserVerify     UserRole = "verify"
	RoleUserSimple     UserRole = "user"
)

func ValidateRoleUser(r UserRole) bool {
	switch r {
	case RoleUserSuperAdmin, RoleUserAdmin, RoleUserVerify, RoleUserSimple:
		return true
	default:
		return false
	}
}

func StringToRoleUser(role string) (UserRole, error) {
	switch role {
	case "super_admin":
		return RoleUserSuperAdmin, nil
	case "admin":
		return RoleUserAdmin, nil
	case "verify":
		return RoleUserVerify, nil
	case "user":
		return RoleUserSimple, nil
	default:
		return "", ErrorRole
	}
}

//	1, if first role is higher priority
//
// -1, if second role is higher priority
//
//	0, if roles are equal
func CompareRolesUser(role1, role2 UserRole) int {
	priority := map[UserRole]int{
		RoleUserSuperAdmin: 5,
		RoleUserAdmin:      4,
		RoleUserVerify:     2,
		RoleUserSimple:     1,
	}

	p1, ok1 := priority[role1]
	p2, ok2 := priority[role2]

	if !ok1 || !ok2 {
		return -1
	}

	if p1 > p2 {
		return 1
	} else if p1 < p2 {
		return -1
	}
	return 0
}
