package usecase

import (
	"redemption/internal/domain"
)

type PointUsecase struct {
	pointRepository domain.PointRepository
}

func NewPointUsecase(pointRepository domain.PointRepository) *PointUsecase {
	return &PointUsecase{
		pointRepository: pointRepository,
	}
}

func (u *PointUsecase) DeletePointByUser(id string) error {
	return u.pointRepository.DeletePointByUser(id)
}
