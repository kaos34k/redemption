package domain

type PointRepository interface {
	DeletePointByUser(id string) error
}
