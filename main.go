package main

import (
	"fmt"
	lemin "lemin/funcs"
	"log"
	"os"
	"strings"
)

func main() {

	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("[USAGE]: go run . example.txt")
		return
	}

	file, err := os.ReadFile(arg[1])
	if err != nil {
		log.Fatal(err)
	}

	slice := strings.Split(string(file), "\n")

	var hold []string
	for _, v := range slice {
		if v == "" {
			continue
		}
		hold = append(hold, v)
	}

	
	lemin.ExtractStartAndEnd(hold)
	EdjeList := lemin.ExtractEdgeList(hold)

	lemin.EdgeListToAdjList(EdjeList)
}
