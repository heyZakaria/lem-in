package readata

import (
	"regexp"
	"strconv"
	"strings"
)

type Start struct {
	start      string
	startCheck bool
}

type End struct {
	end      string
	endCheck bool
}

func ExtractStartAndEnd(data []string) (string, string, int) {
	var s Start
	var e End
	antCount := 0

	// Get ant count from first line
	if len(data) > 0 {
		count, err := strconv.Atoi(data[0])
		if err == nil {
			antCount = count
		}
	}

	// Extract start and end rooms
	for i, v := range data {
		if v == "##start" && i != len(data)-1 && !s.startCheck {
			hold := strings.Split(data[i+1], " ")
			s.start = hold[0]
			s.startCheck = true
		}
		if v == "##end" && i != len(data)-1 && !e.endCheck {
			hold := strings.Split(data[i+1], " ")
			e.end = hold[0]
			e.endCheck = true
		}
	}

	return s.start, e.end, antCount
}

func ExtractEdgeList(data []string) []string {
	var EdgeList []string
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+-[a-zA-Z0-9]+`)

	for _, v := range data {
		edge := pattern.FindString(v)
		if edge != "" {
			EdgeList = append(EdgeList, edge)
		}
	}

	return EdgeList
}

func EdgeListToAdjList(List []string) map[string][]string {
	var hold [][]string
	Map := make(map[string][]string)

	// Process each edge
	for _, v := range List {
		pattern := regexp.MustCompile(`^[a-zA-Z0-9]+-[a-zA-Z0-9]+$`)
		found := pattern.FindString(v)
		if found != "" {
			h := strings.Split(found, "-")
			if len(h) == 2 {
				hold = append(hold, h)
			}
		}
	}

	// Build adjacency list (bidirectional)
	for _, v := range hold {
		Map[v[0]] = append(Map[v[0]], v[1])
		Map[v[1]] = append(Map[v[1]], v[0])
	}

	return Map
}
