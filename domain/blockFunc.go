package domain

import (
	"fmt"
)

func (fb *FunctionBlock) GenerateText() (string, error) {
	switch fb.Function {
	case "SetArgsFunc":
		return fb.SetArgsFunc()
	default:
		return fb.NullResFunc()
	}
}

func (fb *FunctionBlock) SetArgsFunc() (string, error) {

	var args []any
	for _, arg := range fb.Args {
		args = append(args, arg)
	}
	res := fmt.Sprintf(fb.Text, args...)
	return res, nil
}

func (fb *FunctionBlock) NullResFunc() (string, error) {
	return "", ErrFuncitonNotFound
}
