package graphs

import (
	"fmt"
	"lem-in/rooms"
	"lem-in/stack"
)

type Graph struct{
	AdjList map[string][]*rooms.Room
}

func (dfs *Graph) NewGraph(listRooms []string){
	dfs.AdjList = make(map[string][]*rooms.Room)
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

func (dfs *Graph) ShowAdjList(){
	for key , val := range dfs.AdjList {
		fmt.Printf("key : %v values : ",key)
		for _, v := range val {
			fmt.Printf("%v ,",v)
		}
		fmt.Println()
	}
}

var MyStack = stack.NewStack()
func (dfs *Graph) DFSExplore(root *rooms.Room) {
	root.SetVisited(true)
	MyStack.Push(root.GetName())
	if root.GetName() == "t" {
		root.SetVisited(false)
		MyStack.Show()
		MyStack.Pop()
		return
	}
	for _,v := range dfs.AdjList[root.GetName()] {
		if !v.IsVisited() {
			dfs.DFSExplore(v)
			v.SetVisited(false)
		}
	}
	MyStack.Pop()
}


/*
func (dfs *Graph) DFSExplore(root *rooms.Room) {
	myStack := stack.NewStack()
	myStack.Push(root.GetName())
	root.SetVisited(true)
	//dfs.ShowAdjList()				// visited [t,t,f,t,f,f]
	for !myStack.IsNull() {
		n := myStack.Pop().GetValue() //  n = 3                      stack = 0 , 1 , 3
		myStack.Push(n)
		isDone := true
		for _, room := range dfs.AdjList[n] {
			if room.GetName() == "t" {
				myStack.Show()
				break
			}
				//fmt.Printf("%v childs : %v \n",n,*room)
			if  !room.IsVisited() {
				myStack.Push(room.GetName())
				room.SetVisited(true)
				isDone = false
				break
			}
		}
		if isDone {
			node := myStack.Pop().GetValue()
			//dfs.MakeEveryChildUnVisited(node)
			fmt.Printf("visited : %v \n",node)
		}
	}
}
*/

/*
func (dfs *Graph) BFSExplore(root int){
	myQueu := queu.NewQueu(dfs.size)
	myQueu.Push(root)
	for !myQueu.IsNull(){
		n,_ :=myQueu.Front()
		fmt.Printf("visited : %d \n",n)
		current := dfs.array[n].head
		for current != nil {
			myQueu.Push(current.dest)
			current = current.Next
		}
		myQueu.Pop()
	}
}

func (dfs *Graph) NodeRoot(root int){
	head := dfs.array[root].head
	fmt.Printf("%d : ",root)

	for head != nil {
		fmt.Printf(" -> %d",head.dest)
		head = head.Next
	}
	fmt.Println()
}

func (dfs *Graph) DeleteNode(root,dest int) {
	if dfs.array[root].head.dest == dest {
		dfs.array[root].head = dfs.array[root].head.Next
		return
	}
	current := &Node{
		dest: -1,
		Next: nil,
	}
	current.Next = dfs.array[root].head
	for current.Next != nil {
		if current.Next.dest == dest {

			current.Next = current.Next.Next
			current.Next.Next = nil
			//current = nil
		}
		current = current.Next
	}
}



func (dfs *Graph) DFSsearch(root,dest int) {
	visited := make([]bool, dfs.size)
	for i := 0; i < dfs.size; i++ {
		visited = append(visited, false)
	}
	current := dfs.array[root].head
	visited[root] = true
	myStack := stack.NewStack(dfs.size)
	myStack.Push(root)											// visited [t,t,f,t,f,f]
	for !myStack.IsNull() {
		n := myStack.Pop() //  n = 3                      stack = 0 , 1 , 3
		myStack.Push(n)
		current = dfs.array[n].head // dest = 3
		isDone := true
		for current != nil {
			if current.dest == dest {
				myStack.Push(current.dest)
				fmt.Println("Path : ")
				myStack.Show()
				myStack.Pop()
				break
			} else if !visited[current.dest] {
				myStack.Push(current.dest)
				visited[current.dest] = true
				isDone = false
				break
			} else {
				current = current.Next
			}
		}
		if isDone {
			child := myStack.Pop()
			dfs.MakeEveryChildUnVisited(child,visited)
		}
	}
}

func (dfs *Graph) DFSFindAllPaths(start , end int) {
	myStack := stack.NewStack(dfs.size)
	myStack.Push(1)
}
*/