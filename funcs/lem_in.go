package lemin

import (
	"fmt"
	"regexp"
	"strings"
)

func BFS(AdjList map[string][]string) {
	var s Start
	start := s.start
	Queue := []string{start}
	Visited := []string{start}

	for len(Queue) > 0{

	}
	Enqueue()

	Queue, out = Dequeue(Queue)
	fmt.Println(x, out)
}

func Enqueue(Queue map[string][]string, toAdd string) []string {
	Queue = append(Queue, toAdd)

	return Queue
}

func Dequeue(Queue []string) ([]string, string) {
	var out string
	if len(Queue) > 0 {
		out = Queue[0]
	}
	if len(Queue) > 1 {

		x := make([]string, len(Queue)-1)
		x = append(Queue[:0], Queue[1:]...)
		return x, out
	}
	return nil, out
}

func EdgeListToAdjList(List [][]string) map[string][]string {
	var hold [][]string
	m := make(map[string][]string)
	for _, Node := range List {
		for _, v := range Node {

			pattern := regexp.MustCompile(`^[a-zA-Z0-9]+-[a-zA-Z0-9]+$`)
			found := pattern.FindAllString(v, -1)
			if len(found) > 0 {
				h := strings.Split(found[0], "-")
				hold = append(hold, h)
			}

		}
	}

	for _, v := range hold {

		m[v[0]] = append(m[v[0]], v[1])

		m[v[1]] = append(m[v[1]], v[0])

	}
	fmt.Println(m)
	BFS(m)

	return m
}

// room.name visited
