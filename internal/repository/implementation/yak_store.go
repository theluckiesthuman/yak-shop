package implementation

import (
	"sort"
	"sync"

	"github.com/theluckiesthuman/yakshop/internal/entities"
	"github.com/theluckiesthuman/yakshop/internal/repository/contract"
)

type yakStore struct {
	yakStore map[int]entities.Yak
	mx       sync.RWMutex
}

func NewYakStore() contract.YakStore {
	return &yakStore{
		yakStore: make(map[int]entities.Yak),
	}
}

func (y *yakStore) Store(herd entities.Herd) {
	y.Reset()
	y.mx.Lock()
	defer y.mx.Unlock()
	for i, yak := range herd.Yaks {
		yak.ID = i + 1
		y.yakStore[yak.ID] = yak
	}
}

func (y *yakStore) Reset() {
	y.mx.Lock()
	defer y.mx.Unlock()
	y.yakStore = make(map[int]entities.Yak)
}

func (y *yakStore) Read() entities.Herd {
	y.mx.RLock()
	defer y.mx.RUnlock()
	keys := make([]int, 0, len(y.yakStore))
	for k := range y.yakStore {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var herd entities.Herd
	for _, k := range keys {
		herd.Yaks = append(herd.Yaks, y.yakStore[k])
	}
	return herd
}
