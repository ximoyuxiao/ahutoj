package request

type PermissionReq struct {
	UID             string `json:"UID"`
	Administrator   bool   `json:"SuperAdmin"`
	Problem_edit    bool   `json:"ProbelmAdmin"`
	Source_browser  bool   `json:"SourceAdmin"`
	Contest_creator bool   `json:"ContestAdmin"`
}
type EditPermissionReq struct {
	PermissionReq
}
type AddPermissionReq struct {
	PermissionReq
}
type PermissionListReq GetListReq

type DeletePermissionReq struct {
	UIDs []string `json:"UIDs"`
}
