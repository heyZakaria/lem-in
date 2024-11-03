package data

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseFile(arg []string) ([][]string, []string, int, string, string) {
	if len(arg) != 2 {
		fmt.Println("[USAGE]: go run . example.txt")
		os.Exit(1)
	}

	file, err := os.ReadFile(arg[1])
	if err != nil {
		log.Fatal(err)
	}

	slice := strings.Split(string(file), "\n")
	fmt.Printf("slice : %v \n", slice)
	var num_ants int = 0

	firstline := true
	for _, line := range slice {
		if firstline && !(strings.HasPrefix(line, "#") || strings.HasPrefix(line, "L")) {
			numOfAnts, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("ERROR: invalid data format", err)
				os.Exit(1)
			}

			if numOfAnts <= 0 {
				fmt.Println("ERROR: invalid data format, there is no ant")
				os.Exit(1)
			}
			num_ants = numOfAnts
			firstline = false
		}
		if line == "" {
			continue
		}

	}

	// unique1 from rooms
	start, end, unique1 := ExtractStartAndEnd(slice)
	EdjeList := ExtractEdgeList(slice)
	// unique2 from room links
	unique2 := EdgeListToUniques(EdjeList)

	// we check if there is an imposter
	if !Equal(unique1, unique2, start, end) {
		fmt.Println("ERROR: invalid data format, there is an imposter")
		os.Exit(1)
	}
	return EdjeList, unique2, num_ants, start, end
}

func ExtractStartAndEnd(data []string) (string, string, []string) {
	Start := ""
	End := ""
	sCount := 0
	eCount := 0
	checkS := false
	checkE := false
	unique := []string{}

	for i, line := range data {

		// append rooms
		if strings.Contains(line, " ") {
			hold := strings.Split(line, " ")
			if len(hold) == 3 {
				CheckValidRoom(hold)
				unique = append(unique, hold[0])

			} else {
				fmt.Println("ERROR: invalid data format, over or miss corrdinates in line", i, "--->", line)
				os.Exit(1)
			}
		}

		if line == "##start" {

			checkS = true
			sCount++
			continue
		}
		if checkS {

			if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "L") {
				continue
			}
			S := GetStartAndEnd(line, Start, checkS, sCount)
			Start = S
			checkS = false
		}
		if line == "##end" {
			checkE = true
			eCount++
			continue
		}
		if checkE {
			if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "L") {
				continue
			}
			E := GetStartAndEnd(line, End, checkE, eCount)
			End = E
			checkE = false

		}
		if i == len(data)-1 && (Start == "" || End == "") {
			fmt.Println("ERROR: invalid data format, there is no start or no end", line)
			os.Exit(1)
		}

	}

	// return unique from rooms A1 0 1,, A2 0 1, A3 0 1 ....
	return Start, End, unique
}

func ExtractEdgeList(data []string) [][]string {
	var EdgeList [][]string
	for _, line := range data {
		pattern := regexp.MustCompile(`^[a-zA-Z0-9]+-[a-zA-Z0-9$]+`)
		edge := pattern.FindAllString(line, -1)

		if len(edge) > 0 {
			EdgeList = append(EdgeList, edge)
		}
	}
	return EdgeList
}

func EdgeListToUniques(List [][]string) []string {
	var hold [][]string
	Map := make(map[string][]string)
	for _, Node := range List {
		for _, v := range Node {

			roomToRoom := strings.Split(v, "-")
			if roomToRoom[0] == roomToRoom[1] {
				fmt.Println("ERROR: invalid data format, room link to itself --->", roomToRoom[0])
				os.Exit(1)
			}
			hold = append(hold, roomToRoom)
		}
	}

	unique := []string{}
	for _, v := range hold {

		Map[v[0]] = append(Map[v[0]], v[1])

		Map[v[1]] = append(Map[v[1]], v[0])

	}
	for key := range Map {
		unique = append(unique, key)
	}
	// return unique from room links [H3-F2] [H3-H4] [H4-A2] [0-G0] [G0-G1]....
	return unique
}

func Equal(unique1, unique2 []string, start, end string) bool {
	// to check if there is roooms in links didn't present in rooms with coordinates
	if len(unique1) < len(unique2) {
		return false
	}
	count := 0
	// check is to check if the start and end presnt in room links
	checkS := false
	checkE := false
	for _, room1 := range unique1 {
		for _, room2 := range unique2 {
			if room1 == room2 {
				count++
			}
			if start == room2 {
				checkS = true
			}
			if end == room2 {
				checkE = true
			}

		}
	}
	if count != len(unique2) || !checkE || !checkS {
		return false
	}
	return true
}

func CheckValidRoom(room []string) {
	if _, err := strconv.Atoi(room[1]); err != nil {
		fmt.Println("ERROR: invalid data format in coordinates --->", room)
		os.Exit(1)
	}

	if _, err := strconv.Atoi(room[2]); err != nil {
		fmt.Println("ERROR: invalid data format in room --->", room)
		os.Exit(1)
	}
}

func GetStartAndEnd(v, StartOrEnd string, check bool, Count int) string {
	if strings.Contains(v, " ") {
		hold := strings.Split(v, " ")
		if len(hold) == 3 {
			CheckValidRoom(hold)
			StartOrEnd = hold[0]
			check = false
		}
		if Count > 1 {
			fmt.Println("ERROR: invalid data format, more than one start or end")
			os.Exit(1)
		}
	}

	return StartOrEnd
}

//============================================ abdelouahab khiri Code ============================================================

func ReadFromFile(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(file), "\n")
}

func CheckType(data string) string {
	if data == "##start" {
		return "start"
	}
	if data == "##end" {
		return "end"
	}
	result := strings.Split(data, " ")
	size := len(result)
	if size > 0 {
		L_index := search(data, '#')
		if size == 3 && L_index == -1 {
			_, err1 := strconv.Atoi(result[1])
			_, err2 := strconv.Atoi(result[2])
			if err1 == nil && err2 == nil && L_index == -1 {
				return "room"
			}
		} else if size >= 1 {
			if num, err := strconv.Atoi(data); err == nil {
				if num > 0 {
					return "ants"
				}
			} else {
				if L_index == 0 {
					return "comment"
				} else {
					result = strings.Split(data, "-")
					if len(result) == 2 {
						return "link"
					}
				}
			}
		}
	}
	return "error"
}

func GetData(fileName string) ([]string, []string, int, string) {
	slice := ReadFromFile(fileName)
	arrayRooms := []string{}
	LinkedRooms := []string{}
	start := ""
	end := ""
	isStart := false
	isEnd := false
	n_Ants := 0
	for index, line := range slice {
		data_Type := CheckType(line)
		if index == 0 {
			if data_Type == "ants" {
				n_Ants, _ = strconv.Atoi(line)
			} else {
				return arrayRooms, LinkedRooms, -1, "ERROR:invalid number of Ants"
			}
		} else {
			if data_Type == "start" {
				isStart = true
			} else if data_Type == "end" {
				isEnd = true
			} else if data_Type == "room" {
				if isStart {
					if start == "" {
						start = strings.Split(line, " ")[0]
						isStart = false
						continue
					} else {
						return arrayRooms, LinkedRooms, -1, "ERROR: invalid data format"
					}
				}
				if isEnd {
					if end == "" {
						end = strings.Split(line, " ")[0]
						isEnd = false
						continue
					}else {
						return arrayRooms, LinkedRooms, -1, "ERROR: invalid data format"
					}
				} 
				arrayRooms = append(arrayRooms, strings.Split(line, " ")[0])
			} else if data_Type == "link" {
				LinkedRooms = append(arrayRooms, line)
			}
		}
	}
	if (start == "" || end == ""){
		return  arrayRooms, LinkedRooms, -1, "ERROR: invalid data format, no start/end room found"
	}
	newArray := []string{}
	newArray = append(newArray, start)
	arrayRooms = append(arrayRooms, end)
	newArray = append(newArray, arrayRooms...)
	return newArray, LinkedRooms, n_Ants, ""
}

func search(str string, char byte) int {
	size := len(str)
	lastIndex := -1
	for index := 0; index < size; index++ {
		if str[index] == char {
			lastIndex = index
		}
	}
	return lastIndex
}
