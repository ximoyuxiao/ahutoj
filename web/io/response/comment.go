package response

import "ahutoj/web/dao"

type CommentPublish struct {
	Response
}

type CommentList struct {
	Response
	CommentList []dao.Comment `json:"comment_list"` // 评论列表
}
