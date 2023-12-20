package repository

type WisdomRepository interface {
	GetRandomWisdom() (string, error)
}
