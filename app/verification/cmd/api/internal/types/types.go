// Code generated by goctl. DO NOT EDIT.
package types

type VerifyEmailReq struct {
	Email string `json:"email"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type VerifyEmailResp struct {
}

type VerifyImageReq struct {
}

type VerifyImageResp struct {
	ImageUrl string `json:"imageurl"`
}
