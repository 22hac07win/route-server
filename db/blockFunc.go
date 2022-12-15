package db

import (
	"errors"
	"github.com/22hac07win/route-server.git/domain"
)

func (b *DBBlock) GetFunction(name string) (domain.BlockFunc, error) {
	switch name {
	case "NullResFunc":
		return NullResFunc, nil
	default:
		return nil, errors.New("function not found")
	}
}

var NullResFunc = func() (string, error) {
	return "", nil
}
