package objs

import "go-snake-go/internal/common"

type CellType int

const (
	CellEmpty CellType = iota
	CellSnake
	CellApple
)

type Field struct {
	height, width int
	cellsMap      map[common.ObjectPosition]CellType
}

func NewField(height, width int) *Field {
	return &Field{
		height:   height,
		width:    width,
		cellsMap: make(map[common.ObjectPosition]CellType),
	}
}

func (f *Field) Height() int {
	return f.height
}

func (f *Field) Width() int {
	return f.width
}

func (f *Field) GetEmptyPosition() common.ObjectPosition {
	position := common.GetRandomPosition(f.height, f.width)
	for f.cellsMap[position] != CellEmpty {
		position = common.GetRandomPosition(f.height, f.width)
	}
	return position
}

func (f *Field) CellsMap() map[common.ObjectPosition]CellType {
	return f.cellsMap
}

func (f *Field) CheckCellType(position common.ObjectPosition) CellType {
	return f.cellsMap[position]
}

func (f *Field) SetCellType(position common.ObjectPosition, cellType CellType) {
	f.cellsMap[position] = cellType
}

func (f *Field) ClearCell(position common.ObjectPosition) {
	delete(f.cellsMap, position)
}

func (f *Field) SpawnApple() {
	f.SetCellType(f.GetEmptyPosition(), CellApple)
}

func (f *Field) SpawnSnake() *Snake {
	snake := NewSnake(f)
	return snake
}
