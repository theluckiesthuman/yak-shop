package contract

import (
	"github.com/theluckiesthuman/yakshop/internal/dto"
	"github.com/theluckiesthuman/yakshop/internal/entities"
)

type YakManager interface {
	Store(entities.Herd)
	Reset()
	ViewStock(elapsedDays int) (*dto.Stock, error)
	ViewHerd(elapsedDays int) (*dto.Herd, error)
	Order(T int, co dto.CustomerOrder) (*dto.OrderResponse, dto.OrderStatus)
}
