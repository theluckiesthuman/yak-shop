package mapper

import (
	"encoding/xml"

	"github.com/theluckiesthuman/yakshop/internal/dto"
	"github.com/theluckiesthuman/yakshop/internal/entities"
)

func MapReqBodyToHerd(data []byte) (*entities.Herd, error) {
	var herd dto.Herd

	if err := xml.Unmarshal(data, &herd); err != nil {
		return nil, err
	}

	var yaks []entities.Yak
	for _, y := range herd.Yaks {
		yaks = append(yaks, entities.Yak{
			Name:  y.Name,
			Age:   y.Age,
			Sex:   y.Sex,
			Tribe: entities.LabYaks,
		})
	}

	return &entities.Herd{Yaks: yaks}, nil
}

func MapToHerdDto(herd entities.Herd) dto.Herd {
	h := dto.Herd{}
	for _, yak := range herd.Yaks {
		h.Yaks = append(h.Yaks, dto.Yak{
			Name:          yak.Name,
			Age:           yak.Age,
			Sex:           yak.Sex,
			AgeLastShaved: yak.AgeLastShaved,
		})
	}
	return h
}

func MapToHerdEntity(herd dto.Herd, tribe entities.YakTribe) entities.Herd {
	h := entities.Herd{}
	for _, yak := range herd.Yaks {
		h.Yaks = append(h.Yaks, entities.Yak{
			Name:  yak.Name,
			Age:   yak.Age,
			Sex:   yak.Sex,
			Tribe: tribe,
		})
	}
	return h
}
