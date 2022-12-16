package domain

import (
	"errors"
	"fmt"
)

func (fb *FunctionBlock) GetFunction(name string) (BlockFunc, error) {
	switch name {
	case "SetArgsFunc":
		return SetArgsFunc, nil
	default:
		return NullResFunc, errors.New("Function not found")
	}
}

var SetArgsFunc = func(args ...any) (string, error) {
	text, ok := args[0].(string)
	if !ok {
		return "", errors.New("Invalid argument type")
	}

	args = args[1:]

	res := fmt.Sprintf(text, args...)
	return res, nil
}

var NullResFunc = func(_ ...any) (string, error) {
	return "", nil
}
