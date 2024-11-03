package main

import (
	"fmt"
	graphs "lem-in/Graphs"
	paths "lem-in/Paths"
	data "lem-in/ReadData"
	rooms "lem-in/Rooms"
	"os"
	"strings"
)

var (
	ArrayRooms = make([]*rooms.Room, 0)
)

func getRoom(name string) *rooms.Room {
	for _, v := range ArrayRooms {
		if v != nil && v.GetName() == name {
			return v
		}
	}
	return nil
}

func CreateRooms(graph *graphs.Graph, sourceArray []string) {
	for _, name := range sourceArray {
		myRoom := rooms.NewRoom(name, 1, 2)
		ArrayRooms = append(ArrayRooms, myRoom)
	}
	graph.NewGraph(sourceArray)
}

func LinkRoomsTogether(graph *graphs.Graph, RoomsLinkedSource []string) {
	for _, link := range RoomsLinkedSource {
		roomsLinked := strings.Split(link, "-")
		graph.AddRoom(getRoom(roomsLinked[0]), getRoom(roomsLinked[1]))
	}
}

func main() {
    // Check arguments
    arg := os.Args
    rooms,linkedRooms, num_ants, _ := data.GetData(arg[1])
   /* var linksRooms = []string{}
	fmt.Printf("rooms before : %v \n",rooms)
    for i := 0; i < len(EdjeList); i++ {	
		linksRooms = append(linksRooms, (EdjeList[i][0]))
    }
    for i := 0; i < len(rooms); i++ {	
		if rooms[i] == start {
			rooms[0], rooms[i] = rooms[i], rooms[0]	
        }
        if rooms[i] == end {
			rooms[len(rooms)-1], rooms[i] = rooms[i], rooms[len(rooms)-1]	
        }
    }*/
	fmt.Printf("rooms after : %v \n",rooms)
	fmt.Printf("Linked rooms  : %v \n",linkedRooms)

    myGraph := graphs.Graph{}
    // Initialize rooms and links
    CreateRooms(&myGraph, rooms)
    LinkRoomsTogether(&myGraph, linkedRooms)
    // Explore paths
    myGraph.DFSExplore(ArrayRooms[0], ArrayRooms[len(ArrayRooms)-1].GetName())
    if len(myGraph.Paths) == 0 {
        fmt.Println("Error: No valid paths found between start and end rooms")
        return
    }
    // Process paths and distribute ants
    groupedPaths := paths.GroupUniquePaths(myGraph.Paths)
    bestPath := paths.FindBestGroup(groupedPaths, num_ants)
    paths.GetAllPaths(bestPath)
    paths.MakeAntsInPlaces(num_ants)
	paths.ShowPathList()
    result := paths.MoveAnts()
    fmt.Printf("%v", result)
	fmt.Printf("type : %v \n",data.CheckType("##cstart")) 
}