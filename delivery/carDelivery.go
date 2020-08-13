package delivery

import "github.com/edwardsuwirya/carCollection/useCase"

type CarDelivery interface {
	init(uc useCase.CarUseCase) error
}
