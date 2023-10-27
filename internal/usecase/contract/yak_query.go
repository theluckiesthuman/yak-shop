package contract

import (
	"github.com/theluckiesthuman/yakshop/internal/dto"
)

type YakQuery interface {
	CalculateStock(T int) (*dto.Stock, error)
	CalculateAge(T int) (*dto.Herd, error)
}
