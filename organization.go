package roles

type OrgRole string

const (
	RoleOrgOwner  OrgRole = "owner"
	RoleOrgAdmin  OrgRole = "admin"
	RoleOrgModer  OrgRole = "moderator"
	RoleOrgMember OrgRole = "member"
)

func ValidateRoleOrg(r OrgRole) bool {
	switch r {
	case RoleOrgOwner, RoleOrgAdmin, RoleOrgModer, RoleOrgMember:
		return true
	default:
		return false
	}
}

func ParseOrgRole(role string) (OrgRole, error) {
	switch role {
	case "owner":
		return RoleOrgOwner, nil
	case "admin":
		return RoleOrgAdmin, nil
	case "moderator":
		return RoleOrgModer, nil
	case "member":
		return RoleOrgMember, nil
	default:
		return "", ErrorRole
	}
}

//	1, if first role is higher priority
//
// -1, if second role is higher priority
//
//	0, if roles are equal
func CompareRolesOrg(role1, role2 OrgRole) int {
	priority := map[OrgRole]int{
		RoleOrgOwner:  4,
		RoleOrgAdmin:  3,
		RoleOrgModer:  2,
		RoleOrgMember: 1,
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
