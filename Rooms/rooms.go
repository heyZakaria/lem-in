package rooms

import(
	//"lem-in/graphs"
)
type Room struct {
	name string
	x,y int
	isOpen,isVisited bool
}

func NewRoom(name_ string,x_,y_ int) *Room{
	return &Room{
		name :name_,
		x: x_,
		y: y_,
		isOpen: false,
		isVisited: false,
	}
}

func (R *Room) IsVisited() bool{
	return R.isVisited
}

func (R *Room) IsOped() bool{
	return R.isOpen
}

func (R *Room) GetName() string{
	return R.name
}

func (R *Room) SetVisited(b bool) {
	R.isVisited = b
}