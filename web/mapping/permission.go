package mapping

import "ahutoj/web/dao"

type PermissionBit int

const (
	UNLOGINBit       PermissionBit = 0x01 // B0000 0001
	CommomUserBit    PermissionBit = 0x02 // B0000 0010
	AdministratorBit PermissionBit = 0x04 // B0000 0100
	ProblemAdminBit  PermissionBit = 0x08 // B0000 1000
	ContestAdminBit  PermissionBit = 0x10 // B0001 0000
	SourceBorwserBit PermissionBit = 0x20 // B0010 0000
	ListAdminBit     PermissionBit = 0x40 // B0100 0000
	SuperAdminBit    PermissionBit = 0x80 // B1000 0000
)

func AddPermissionBit(permissionMap *PermissionBit, bit PermissionBit) {
	*permissionMap = *permissionMap | bit
}
func PermissionToBitMap(permission dao.Permission) int {
	ans := PermissionBit(UNLOGINBit)
	if permission.Administrator == "Y" {
		AddPermissionBit(&ans, SuperAdminBit)
	}
	if permission.Contest_creator == "Y" {
		AddPermissionBit(&ans, AdministratorBit)
	}
	if permission.Problem_edit == "Y" {
		AddPermissionBit(&ans, ProblemAdminBit)
	}
	if permission.Source_browser == "Y" {
		AddPermissionBit(&ans, SourceBorwserBit)
	}
	return int(ans)
}
