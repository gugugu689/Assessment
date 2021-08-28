package Models

type ParamCondition struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"form"`
	Order string `json:"order" form:"form"`
}
