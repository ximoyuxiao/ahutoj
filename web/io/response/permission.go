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
	Conut int64        `json:"Count"`
	Data  []Permission `json:"Data"`
}
