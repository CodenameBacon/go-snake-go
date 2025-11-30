package objs

import "go-snake-go/internal/common"

type FieldModel struct {
	Height int                                `json:"height"`
	Width  int                                `json:"width"`
	Cells  map[common.ObjectPosition]CellType `json:"cells"`
}
