package response

import "ahutoj/web/dao"

type SolutionPublish struct {
	Response
}
type SolutionList struct {
	Response
	Count        int            `json:"count"`
	SolutionList []dao.Solution `json:"solution_list"`
}
type Solution struct {
	Response
	Solution dao.Solution `json:"solution"`
}
