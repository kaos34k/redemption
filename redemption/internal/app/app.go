package app

import (
	"redemption/internal/usecase"
)

type MyApp struct {
	pointUsecase usecase.PointUsecase
}

func NewMyApp(pointUsecase usecase.PointUsecase) *MyApp {
	return &MyApp{
		pointUsecase: pointUsecase,
	}
}

func (a *MyApp) HandleRequest(id string) error {
	err := a.pointUsecase.DeletePointByUser(id)
	if err != nil {
		return err
	}

	return nil
}
