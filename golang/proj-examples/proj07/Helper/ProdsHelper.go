package Helper

type ProdsRequest struct {
	// Size int `form:"size"` // form表单形式提交
	Size int `json:"size"` // json形式提交
}
