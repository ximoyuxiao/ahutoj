package response

type Permission struct {
	Uid           string `json:"uid"`
	PermissionMap int    `json:"permission_map"`
}

type PermissionResp struct {
	Response
	Permission
}

type PermissionListResp struct {
	Response
	Conut int          `json:"count"`
	Data  []Permission `json:"data"`
}
