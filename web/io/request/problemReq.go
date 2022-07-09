package request

type Permission struct {
	Uid             string `form:"uid"`
	Administrator   string `form:"administrator"`
	Problem_edit    string `form:"problem_edit"`
	Source_browser  string `form:"source_browser"`
	Contest_creator string `form:"contest_creator"`
}
