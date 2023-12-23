package repository

import (
	"math/rand"
	"wisdom/internal/domain/repository"
)

type WisdomRepository struct {
}

var _ repository.WisdomRepository = &WisdomRepository{}

func NewWisdomRepository() *WisdomRepository {
	return &WisdomRepository{}
}

func (w WisdomRepository) GetRandomWisdom() (string, error) {
	arr := []string{
		"Be yourself; everyone else is already taken.",
		"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
		"So many books, so little time.",
		"A room without books is like a body without a soul.",
	}
	randNum := rand.Intn(len(arr))
	return arr[randNum], nil
}
