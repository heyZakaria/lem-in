package paths

import (
	"fmt"
	"sort"
	"strconv"
)

var (
	PathsList = make([]path, 0)
)

type path struct {
	length       int
	Ants         []string
	arrayOfRooms []string
}

func GetAllPaths(arrayOfPaths [][]string) {
	arrayOfPaths = sortPaths(arrayOfPaths)
	for _, p := range arrayOfPaths {
		mypath := path{length: len(p), arrayOfRooms: make([]string, 0)}
		for _, room := range p {
			mypath.arrayOfRooms = append(mypath.arrayOfRooms, room)
		}
		PathsList = append(PathsList, mypath)
	}
}

func ShowPathList() {
	for _, path := range PathsList {
		fmt.Println("====================================")
		fmt.Printf("length : %d \n", path.length)
		fmt.Printf("Ants : %v \n", path.Ants)
		for _, room := range path.arrayOfRooms {
			fmt.Printf("-> %v", room)
		}
		fmt.Println()
	}
}

func MakeAntsInPlaces(number_Ants int) {
	paths_size := len(PathsList)
	ant_Id := 1
	index := 0
	for ant_Id <= number_Ants {
		if index < paths_size-1 {
			room_size := PathsList[index].length
			ants_Size := len(PathsList[index].Ants)
			if room_size+ants_Size > PathsList[index+1].length {
				index++
				PathsList[index].Ants = append(PathsList[index].Ants, "L"+strconv.Itoa(ant_Id))
			} else {
				index = 0
				PathsList[index].Ants = append(PathsList[index].Ants, "L"+strconv.Itoa(ant_Id))
			}
		} else {
			index = 0
			PathsList[index].Ants = append(PathsList[index].Ants, "L"+strconv.Itoa(ant_Id))
		}
		ant_Id++
	}
}
func sortPaths(mypaths [][]string) [][]string {
	sort.Slice(mypaths, func(i, j int) bool {
		if len(mypaths[i]) != len(mypaths[j]) {
			return len(mypaths[i]) < len(mypaths[j])
		}
		return false
	})
	return mypaths
}

func arePathsEqual(path1, path2 []string) bool {
	for i := 0; i < len(path1[1:len(path1)-1]); i++ {
		for j := 0; j < len(path2[1:len(path2)-1]); j++ {
			if path1[1 : len(path1)-1][i] == path2[1 : len(path2)-1][j] {
				return true
			}
		}
	}
	return false
}

func GroupUniquePaths(paths [][]string) [][][]string {
	paths = sortPaths(paths)
	var groups [][][]string
	used := make(map[int]bool)
	for i, path := range paths {
		if used[i] {
			continue
		}
		group := [][]string{path}
		for j := i + 1; j < len(paths); j++ {
			if used[j] {
				continue
			}
			unique := true
			for _, groupPath := range group {
				if arePathsEqual(paths[j], groupPath) {
					unique = false
					break
				}
			}
			if unique {
				group = append(group, paths[j])
			}
		}
		groups = append(groups, group)
	}
	return groups
}

type Move struct {
	pathIndex int
	ant       string
	room      string
}

func MoveAnts() string{
	result := ""
	ListResult := make([]Move, 0)
	firstMove := true
	ListResultLen := 0
	for ListResultLen > 0 || firstMove {
		if ListResultLen > 0 {
			for i := 0; i < ListResultLen; i++ {
				roomIndex := nextRoom(ListResult[i].room, PathsList[ListResult[i].pathIndex].arrayOfRooms[1:])
				if roomIndex != -1 {
					ListResult[i].room = PathsList[ListResult[i].pathIndex].arrayOfRooms[1:][roomIndex]
					result += ListResult[i].ant + "-" + ListResult[i].room + " "
				} else {
					// remove ant from list
					ListResult = ListResult[1:]
					ListResultLen--
					i--
				}
			}
		}
		PathsListLen := len(PathsList)
		for index, path := range PathsList {
			if len(path.Ants) > 0 && len(path.arrayOfRooms) > 0 {
				ListResult = append(ListResult, Move{ant: path.Ants[0], room: path.arrayOfRooms[1], pathIndex: index})
				ListResultLen++
				result += path.Ants[0] + "-" + path.arrayOfRooms[1]
				PathsList[index].Ants = PathsList[index].Ants[1:]
				if index < PathsListLen - 1 {
					result += " "
				}
			}
		}
		if firstMove || ListResultLen > 0 {
			lenRes := len(result)
			if result[lenRes-1] ==' '{
				result = result[:lenRes-1]
			}
			result += "\n"
		}
		firstMove = false

	}
	return  result
}

func nextRoom(room string, path []string) int {
	size := len(path)
	for i := 0; i < size; i++ {
		if i < size-1 && room == path[i] {
			return i + 1
		}
	}
	return -1
}

func calculateTotalSteps(group [][]string) int {
	total := 0
	for _, path := range group {
		total += len(path) - 1
	}
	return total
}
func calculateRounds(group [][]string, numAnts int) int {
	paths_size := len(group)
	if paths_size == 0 {
		return 0
	}
	// Simulate MakeAntsInPlaces distribution to get actual ant counts per path
	distribution := make([]int, paths_size)
	index := 0
	ant_Id := 1

	// Distribute ants using same logic as MakeAntsInPlaces
	for ant_Id <= numAnts {
		if index < paths_size-1 {
			current_path_length := len(group[index]) - 1 // -1 for start room
			ants_in_path := distribution[index]
			next_path_length := len(group[index+1]) - 1

			if current_path_length+ants_in_path > next_path_length {
				index++
				distribution[index]++
			} else {
				index = 0
				distribution[index]++
			}
		} else {
			index = 0
			distribution[index]++
		}
		ant_Id++
	}
	// Find max rounds needed
	maxRounds := 0
	for i, antCount := range distribution {
		if antCount > 0 {
			pathLength := len(group[i]) - 1
			roundsNeeded := pathLength + (antCount - 1)
			if roundsNeeded > maxRounds {
				maxRounds = roundsNeeded
			}
		}
	}
	return maxRounds
}

func FindBestGroup(groups [][][]string, numAnts int) [][]string {
	if len(groups) == 0 {
		return nil
	}

	bestGroup := groups[0]
	bestRounds := calculateRounds(bestGroup, numAnts)
	bestSteps := calculateTotalSteps(bestGroup)

	// Compare each group
	for _, group := range groups {
		rounds := calculateRounds(group, numAnts)
		steps := calculateTotalSteps(group)

		// Update best group if current group is better
		if rounds < bestRounds || (rounds == bestRounds && steps < bestSteps) {
			bestRounds = rounds
			bestSteps = steps
			bestGroup = group
		}
	}

	//fmt.Printf("Best group needs %d rounds\n", bestRounds)
	return bestGroup
}

func PrintData(paths [][][]string) {
	for i, group := range paths {
		fmt.Printf("Group %d:\n", i+1)
		for _, path := range group {
			fmt.Printf("  %v\n", path)
		}
		fmt.Println()
	}
}
