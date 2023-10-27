package implementation

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/theluckiesthuman/yakshop/internal/dto"
	"github.com/theluckiesthuman/yakshop/internal/entities"
	"github.com/theluckiesthuman/yakshop/internal/mapper"
	"github.com/theluckiesthuman/yakshop/internal/usecase/contract"
)

type yakQuery struct {
	herd *entities.Herd
}

func NewFileQuery(f *os.File) (contract.YakQuery, error) {
	var herd dto.Herd
	decoder := xml.NewDecoder(f)
	err := decoder.Decode(&herd)
	if err != nil {
		return nil, fmt.Errorf("error parsing XML: %w", err)
	}
	mherd := mapper.MapToHerdEntity(herd, entities.LabYaks)
	return &yakQuery{
		herd: &mherd,
	}, nil
}

// Deep copy herd before modifying
func copyHerd(original *entities.Herd) entities.Herd {
	copied := entities.Herd{}
	for _, yak := range original.Yaks {
		copiedYak := &entities.Yak{
			Name: yak.Name,
			Age:  yak.Age,
			Sex:  yak.Sex,
		}
		copied.Yaks = append(copied.Yaks, *copiedYak)
	}
	return copied
}

func (yq *yakQuery) CalculateStock(T int) (*dto.Stock, error) {
	herdCopy := copyHerd(yq.herd)
	var totalMilk float64
	var totalSkin int

	for _, yak := range herdCopy.Yaks {
		totalMilk += yak.CalculateMilk(T)
		totalSkin += yak.GetSkinStockTillDay(T)
	}

	return &dto.Stock{Milk: totalMilk, Skins: totalSkin}, nil
}

func (fq *yakQuery) CalculateAge(T int) (*dto.Herd, error) {
	herdCopy := copyHerd(fq.herd)
	for i, yak := range herdCopy.Yaks {
		yak.Age = yak.GetCurrentYakAgeInYears(T)
		herdCopy.Yaks[i] = yak
	}
	dtoHerd := mapper.MapToHerdDto(herdCopy)
	return &dtoHerd, nil
}
