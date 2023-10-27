package contract

import "github.com/theluckiesthuman/yakshop/internal/entities"

type YakStore interface {
	Store(entities.Herd)
	Reset()
	Read() entities.Herd
}
