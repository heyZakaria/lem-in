package graphs

import (
	"lem-in/Rooms"
	"lem-in/stack"
)

type Graph struct{
	AdjList map[string][]*rooms.Room
	Paths   [][]string
}

func (my_Graph *Graph) NewGraph(listRooms []string){
	my_Graph.AdjList = make(map[string][]*rooms.Room)
	my_Graph.Paths = make([][]string, 0)
	for _, v := range listRooms {
		my_Graph.AdjList[v] = nil
	}
}

func (my_Graph *Graph) AddRoom(src , destRoom *rooms.Room) bool {
	_, src_room := my_Graph.AdjList[src.GetName()]
	_, dest_room := my_Graph.AdjList[destRoom.GetName()]
	if !src_room || !dest_room {
		return false
	}
	my_Graph.AdjList[src.GetName()] = append(my_Graph.AdjList[src.GetName()], destRoom)
	my_Graph.AdjList[destRoom.GetName()] = append(my_Graph.AdjList[destRoom.GetName()],src)
	return true
}

var MyStack = stack.NewStack()
func (my_Graph *Graph) DFSExplore(root *rooms.Room,end string) {
	root.SetVisited(true)
	MyStack.Push(root.GetName())
	if root.GetName() == end {
		root.SetVisited(false)
		my_Graph.Paths = append(my_Graph.Paths, MyStack.ConvertToArray())
		MyStack.Show()
		MyStack.Pop()
		return
	}
	for _,v := range my_Graph.AdjList[root.GetName()] {
		if !v.IsVisited() {
			my_Graph.DFSExplore(v,end)
			v.SetVisited(false)
		}
	}
	MyStack.Pop()
}
