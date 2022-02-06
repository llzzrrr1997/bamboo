package id

import (
	"github.com/rs/xid"
)

type BambooIDService struct {
}

func NewBambooIDService(params ...interface{}) (interface{}, error) {
	return &BambooIDService{}, nil
}

func (s *BambooIDService) NewID() string {
	return xid.New().String()
}
