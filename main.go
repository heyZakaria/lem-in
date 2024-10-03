package main

import (
	//"fmt"
	"lem-in/graphs"
	"lem-in/rooms"
)



func main(){
	myGraph := graphs.Graph{}
	 myGraph.NewGraph([]string{"s","0","1","2","3","4","5","6","7","8","t"})
	//room_start := rooms.NewRoom("s",1,2)
	room_s := rooms.NewRoom("s",1,2)
	room_0 := rooms.NewRoom("0",1,2)
	room_1 := rooms.NewRoom("1",1,2)
	room_2 := rooms.NewRoom("2",1,2)
	room_3 := rooms.NewRoom("3",1,2)
	room_4 := rooms.NewRoom("4",1,2)
	room_5 := rooms.NewRoom("5",1,2)
	room_6 := rooms.NewRoom("6",1,2)
	room_7 := rooms.NewRoom("7",1,2)
	room_8 := rooms.NewRoom("8",1,2)

	room_t := rooms.NewRoom("t",1,2)


	 myGraph.AddRoom(room_s,room_0)
	 myGraph.AddRoom(room_s,room_1)
	 myGraph.AddRoom(room_s,room_2)
	 myGraph.AddRoom(room_0,room_3)
	 myGraph.AddRoom(room_0,room_4) 
	 myGraph.AddRoom(room_1,room_4)
	 myGraph.AddRoom(room_1,room_5)
	 myGraph.AddRoom(room_2,room_3)
	 myGraph.AddRoom(room_2,room_7)
	 myGraph.AddRoom(room_3,room_6)
	 myGraph.AddRoom(room_3,room_7)
	 myGraph.AddRoom(room_4,room_6)
	 myGraph.AddRoom(room_4,room_8)
	 myGraph.AddRoom(room_4,room_5)
	 myGraph.AddRoom(room_5,room_8)
	 myGraph.AddRoom(room_6,room_t)
	 myGraph.AddRoom(room_7,room_t)
	 myGraph.AddRoom(room_8,room_t)
	 myGraph.DFSExplore(room_s)
	//myGraph.ShowAdjList()
	/*myGraph := NewGraph(11)
	myGraph.AddNode(0,4)
	myGraph.AddNode(0,6)
	myGraph.AddNode(1,3)
	myGraph.AddNode(4,3)
	myGraph.AddNode(5,2)
	myGraph.AddNode(3,5)
	myGraph.AddNode(4,2)
	myGraph.AddNode(2,1)
	myGraph.AddNode(7,6)
	myGraph.AddNode(7,2)
	myGraph.AddNode(7,4)
	myGraph.AddNode(6,5)
	myGraph.DFSsearch(0,5)*/
}