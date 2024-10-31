package main

import (
	"fmt"
	graphs "lem-in/Graphs"
	paths "lem-in/Paths"
	readata "lem-in/ReadData"
	rooms "lem-in/Rooms"
	"log"
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
	if len(arg) != 2 {
		fmt.Println("[USAGE]: go run . example.txt")
		return
	}

	// Read file
	file, err := os.ReadFile(arg[1])
	if err != nil {
		log.Fatal(err)
	}

	// Process input
	slice := strings.Split(string(file), "\n")
	var hold []string
	for _, v := range slice {
		if v == "" {
			continue
		}
		hold = append(hold, v)
	}

	// Extract start, end rooms and ant count
	start, end, antCount := readata.ExtractStartAndEnd(hold)
	if start == "" || end == "" {
		fmt.Println("Error: Invalid start or end room")
		return
	}

	// Get edges and create adjacency list
	EdjeList := readata.ExtractEdgeList(hold)
	if len(EdjeList) == 0 {
		fmt.Println("Error: No valid paths found")
		return
	}
	AdjList := readata.EdgeListToAdjList(EdjeList)
	// Create graph
	myGraph := graphs.Graph{}
	// Create array of room names
	var arr []string
	for k := range AdjList {
		arr = append(arr, k)
	}
	// Initialize rooms and links
	CreateRooms(&myGraph, arr)
	LinkRoomsTogether(&myGraph, EdjeList)
	// Find start room object
	startRoom := getRoom(start)
	if startRoom == nil {
		fmt.Println("Error: Start room not found in graph")
		return
	}
	// Explore paths
	myGraph.DFSExplore(startRoom, end)
	if len(myGraph.Paths) == 0 {
		fmt.Println("Error: No valid paths found between start and end rooms")
		return
	}
	// Process paths and distribute ants
	groupedPaths := paths.GroupUniquePaths(myGraph.Paths)
	bestPath := paths.FindBestGroup(groupedPaths, antCount)
	paths.GetAllPaths(bestPath)
	paths.MakeAntsInPlaces(antCount)
	result := paths.MoveAnts()
	fmt.Printf("%v",result)
}
