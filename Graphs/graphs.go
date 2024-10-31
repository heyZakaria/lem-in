package graphs

import (
	"lem-in/Rooms"
	"lem-in/stack"
)

type Graph struct{
	AdjList map[string][]*rooms.Room
	Paths   [][]string
}

func (dfs *Graph) NewGraph(listRooms []string){
	dfs.AdjList = make(map[string][]*rooms.Room)
	dfs.Paths = make([][]string, 0)
	for _, v := range listRooms {
		dfs.AdjList[v] = nil
	}
}

func (dfs *Graph) AddRoom(src , destRoom *rooms.Room) bool {
	_, src_room := dfs.AdjList[src.GetName()]
	_, dest_room := dfs.AdjList[destRoom.GetName()]
	if !src_room || !dest_room {
		return false
	}
	dfs.AdjList[src.GetName()] = append(dfs.AdjList[src.GetName()], destRoom)
	dfs.AdjList[destRoom.GetName()] = append(dfs.AdjList[destRoom.GetName()],src)
	return true
}

/*func (dfs *Graph) ShowAdjList(){
	for key , val := range dfs.AdjList {
		fmt.Printf("key : %v values : ",key)
		for _, v := range val {
			fmt.Printf("%v ,",v)
		}
		fmt.Println()
	}
}*/


var MyStack = stack.NewStack()
func (dfs *Graph) DFSExplore(root *rooms.Room,end string) {
	root.SetVisited(true)
	MyStack.Push(root.GetName())
	if root.GetName() == end {
		root.SetVisited(false)
		dfs.Paths = append(dfs.Paths, MyStack.ConvertToArray())
		//MyStack.Show()
		MyStack.Pop()
		return
	}
	for _,v := range dfs.AdjList[root.GetName()] {
		if !v.IsVisited() {
			dfs.DFSExplore(v,end)
			v.SetVisited(false)
		}
	}
	MyStack.Pop()
}
