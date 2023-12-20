package repository

import (
	"wisdom/internal/domain/repository"
)

type WisdomRepository struct {
}

var _ repository.WisdomRepository = &WisdomRepository{}

func NewWisdomRepository() *WisdomRepository {
	return &WisdomRepository{}
}

func (w WisdomRepository) GetRandomWisdom() (string, error) {
	return "It is wisdom", nil
}
