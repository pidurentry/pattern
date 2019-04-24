package action

import "github.com/pidurentry/pattern"

type Expression struct {
	Expression pattern.Expression `json:"expression"`
	True       []interface{}      `json:"true"`
	False      []interface{}      `json:"false"`
}
