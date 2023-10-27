package implementation

import (
	"github.com/theluckiesthuman/yakshop/internal/dto"
	"github.com/theluckiesthuman/yakshop/internal/entities"
	"github.com/theluckiesthuman/yakshop/internal/mapper"
	ys "github.com/theluckiesthuman/yakshop/internal/repository/contract"
	"github.com/theluckiesthuman/yakshop/internal/usecase/contract"
)

type yakManager struct {
	store ys.YakStore
}

func NewYakManager(store ys.YakStore) contract.YakManager {
	return &yakManager{
		store: store,
	}
}

func (y *yakManager) Store(herd entities.Herd) {
	y.store.Store(herd)
}

func (y *yakManager) Reset() {
	y.store.Reset()
}

func (y *yakManager) ViewStock(elapsedDays int) (*dto.Stock, error) {
	herd := y.store.Read()
	var totalMilk float64
	var totalWool int

	for _, yak := range herd.Yaks {
		totalMilk += yak.CalculateMilk(elapsedDays)
		totalWool += yak.GetSkinStockTillDay(elapsedDays)
	}

	return &dto.Stock{Milk: totalMilk, Skins: totalWool}, nil
}

func (y *yakManager) ViewHerd(elapsedDays int) (*dto.Herd, error) {
	herd := y.store.Read()
	for i, yak := range herd.Yaks {
		yak.AgeLastShaved = yak.CalculateLastShavedAge(elapsedDays)
		yak.Age = yak.GetCurrentYakAgeInYears(elapsedDays)
		yak.Sex = ""
		herd.Yaks[i] = yak
	}
	dtoHerd := mapper.MapToHerdDto(herd)
	return &dtoHerd, nil
}

func (y *yakManager) Order(T int, co dto.CustomerOrder) (*dto.OrderResponse, dto.OrderStatus) {
	herd := y.store.Read()
	var totalMilk float64
	var totalSkins int
	order := co.Order
	for _, yak := range herd.Yaks {
		totalMilk += yak.CalculateMilk(T)
		totalSkins += yak.GetSkinStockTillDay(T)
	}

	response := dto.OrderResponse{}
	status := dto.Unfulfilled
	if order.Milk <= totalMilk && order.Skins <= totalSkins {
		response.Milk = order.Milk
		response.Skins = order.Skins
		status = dto.Fulfilled
	} else if order.Milk > totalMilk && order.Skins <= totalSkins {
		response.Skins = order.Skins
		status = dto.Partial
	} else if order.Milk <= totalMilk && order.Skins > totalSkins {
		response.Milk = totalMilk
		status = dto.Partial
	}
	return &response, status
}
