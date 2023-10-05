package response

import "ahutoj/web/dao"

type SolutionPublish struct {
	Response
}
type SolutionList struct {
	Response
	SolutionList []dao.Solution `json:"solution_list"`
}
type Solution struct {
	Response
	Solution dao.Solution `json:"solution"`
}
