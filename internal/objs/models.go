package objs

import "go-snake-go/internal/common"

type FieldModel struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type AppleModel struct {
	Position common.ObjectPosition `json:"position"`
}

type SnakeModel struct {
	Nodes []*SnakeNodeModel `json:"nodes"`
}

type SnakeNodeModel struct {
	Position common.ObjectPosition `json:"position"`
}
