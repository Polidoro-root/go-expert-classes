package usecase

import (
	"github.com/Polidoro-root/go-expert-classes/18_clean_architecture/internal/entity"
	"github.com/Polidoro-root/go-expert-classes/18_clean_architecture/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,

		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrdersUseCase) Execute() ([]entity.Order, error) {

	orders, err := l.OrderRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return orders, nil
}
