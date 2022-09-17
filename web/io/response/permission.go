package response

type Permission struct {
	UID           string `json:"UID"`
	PermissionMap int    `json:"PermissionMap"`
}

type PermissionResp struct {
	Response
	Permission
}

type PermissionListResp struct {
	Response
	Conut int          `json:"Count"`
	Data  []Permission `json:"Data"`
}
