package lemin

import (
	"fmt"
	"regexp"
	"strings"
)

func BFS(AdjList map[string][]string) {

	//var s Start
	//start := s.start
	var Queue []string {"1", "2", "3"}

	x := Dequeue(Queue)
	fmt.Println(x)
	x = Dequeue(Queue)
	fmt.Println(x)


}

func Enqueue(Queue []string, toAdd string) []string{
	Queue = append(Queue, toAdd)

	return Queue
}

func Dequeue(Queue []string) []string{
	if len(Queue) > 1{

		Queue = append(Queue[1:2], Queue[2:]...)
		return Queue
	}
	return nil

}

func EdgeListToAdjList(List [][]string) map[string][]string{
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
	BFS(m)

	return m
}
