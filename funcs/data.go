package lemin

import (
	"fmt"
	"regexp"
	"strings"
)

type Start struct {
	start      string
	startCheck bool
	x          string
	y          string
}

type End struct {
	end      string
	endCheck bool
	x        string
	y        string
}

func ExtractStartAndEnd(data []string) {
	var s Start
	var e End
	fmt.Println(s.startCheck)
	for i, v := range data {
		if v == "##start" && i != len(data)-1 && !s.startCheck {
			hold := strings.Split(data[i+1], " ")
			s.start = hold[0]
			s.x = hold[0]
			s.y= hold[0]
			s.startCheck = true
		}
		if v == "##end" && i != len(data)-1 && !e.endCheck {
			hold := strings.Split(data[i+1], " ")
			e.end = hold[0]
			e.x= hold[1]
			e.y = hold[2]
			e.endCheck = true
		}
	}
}

func ExtractEdgeList(data []string) [][]string {

	var EdgeList [][]string
	for _, v := range data {
		pattern := regexp.MustCompile(`^[a-zA-Z0-9]+-[a-zA-Z0-9$]+`)
		edge := pattern.FindAllString(v, -1)
		if len(edge) > 0 {
			EdgeList = append(EdgeList, edge)
		}
	}

	return EdgeList
}
