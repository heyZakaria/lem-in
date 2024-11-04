package main

import (
	"fmt"
	"os"
	"strings"

	graphs "lem-in/Graphs"
	paths "lem-in/Paths"
	data "lem-in/ReadData"
	rooms "lem-in/Rooms"
)

var ArrayRooms = make([]*rooms.Room, 0)

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
	// rooms,linkedRooms, num_ants, _ := data.GetData(arg[1])
	dataFile,EdjeList, rooms, num_ants, start, end := data.ParseFile(arg)
	linksRooms := []string{}
	for i := 0; i < len(EdjeList); i++ {
		linksRooms = append(linksRooms, (EdjeList[i][0]))
	}
	for i := 0; i < len(rooms); i++ {
		if rooms[i] == start {
			rooms[0], rooms[i] = rooms[i], rooms[0]
		}
	}

	for i := 0; i < len(rooms); i++ {
		if rooms[i] == end {
			rooms[len(rooms)-1], rooms[i] = rooms[i], rooms[len(rooms)-1]
		}
	}
	myGraph := graphs.Graph{}
	// Initialize rooms and links
	CreateRooms(&myGraph, rooms)
	LinkRoomsTogether(&myGraph, linksRooms)
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
	Priint(dataFile)
	result := paths.MoveAnts()
	fmt.Printf("%v", result)
}

func Priint(dataFile []string) {
	for _, v := range dataFile {
		if  data.CheckType(v) != "comment"{
			fmt.Println(v)
		}
	}
	fmt.Println()
}
