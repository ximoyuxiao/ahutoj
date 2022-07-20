package response

type Permission struct {
	Administrator   bool `json:"administrator"`
	Problem_edit    bool `json:"problem_edit"`
	Source_browser  bool `json:"source_browser"`
	Contest_creator bool `json:"contest_creator"`
}
