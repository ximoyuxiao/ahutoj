package request

type PermissionReq struct {
	Uid             string `json:"uid"`
	Administrator   bool   `json:"administrator"`
	Problem_edit    bool   `json:"problem_edit"`
	Source_browser  bool   `json:"source_browser"`
	Contest_creator bool   `json:"contest_creator"`
}
type EditPermissionReq struct {
	PermissionReq
}
type AddPermissionReq struct {
	PermissionReq
}
type PermissionListReq GetListReq

type DeletePermissionReq struct {
	Uids []string `json:"uids"`
}
