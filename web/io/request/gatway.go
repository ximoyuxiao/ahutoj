package request

import "ahutoj/web/io/constanct"

type AddRouterReq struct {
	FromURL     string                `json:"FromURL"`
	Method      string                `json:"Method"`
	ToHost      string                `json:"Host"`
	Weight      int64                 `json:"Weight"`
	VerfiyLevel constanct.VerfiyLevel `json:"VerfiyLevel"`
}

type DelRouterReq struct {
	FromURL string `json:"FromURL"`
	Method  string `json:"Method"`
	ToHost  string `json:"Host"`
}
