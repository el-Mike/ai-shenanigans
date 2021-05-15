package game

type Player interface {
	Move()
	GetSign() Sign
}
