package response

import "ahutoj/web/dao"

type SolutionPublish struct {
	Response
}
type SolutionList struct {
	Response
	SolutionList []dao.Solution `json:"solution_list"`
}
