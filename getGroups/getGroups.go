package main

import (
	"fmt"
	"sort"
)

func arePathsEqual(path1, path2 []string) bool {
	fmt.Println(path1[1:len(path1)-1], " , ", path2[1:len(path2)-1])
	for i := 0; i < len(path1[1:len(path1)-1]); i++ {
		for j := 0; j < len(path2[1:len(path2)-1]); j++ {
			if path1[1 : len(path1)-1][i] == path2[1 : len(path2)-1][j] {
				return true
			}
		}
	}
	return false
}

func sortPaths(paths [][]string) [][]string {
	sort.Slice(paths, func(i, j int) bool {
		if len(paths[i]) != len(paths[j]) {
			return len(paths[i]) < len(paths[j])
		}
		return false
	})
	return paths
}

func groupUniquePaths(paths [][]string) [][][]string {
	paths = sortPaths(paths)
	var groups [][][]string
	used := make(map[int]bool)
	for i, path := range paths {
		if used[i] {
			continue
		}
		group := [][]string{path}
		used[i] = true
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

func main() {
	paths := [][]string{
		{"1", "3", "9", "0"}, {"1", "u", "z", "a", "n", "o", "0"}, {"1", "3", "4", "2", "7", "6", "0"},
		{"1", "3", "4", "7", "6", "0"}, {"1", "3", "4", "7", "2", "5", "6", "0"}, {"1", "3", "5", "2", "4", "0"},
		{"1", "3", "5", "2", "4", "7", "6", "0"}, {"1", "3", "5", "2", "7", "6", "0"}, {"1", "3", "5", "2", "7", "4", "0"},
		{"1", "3", "8", "7", "0"}, {"1", "3", "5", "6", "7", "2", "4", "0"}, {"1", "3", "5", "6", "7", "4", "0"},
		{"1", "2", "5", "3", "4", "0"}, {"1", "2", "5", "3", "4", "7", "6", "0"}, {"1", "2", "5", "6", "0"},
		{"1", "2", "5", "9", "7", "4", "0"}, {"1", "2", "4", "0"}, {"1", "2", "4", "3", "5", "6", "0"},
		{"1", "2", "4", "7", "6", "0"}, {"1", "2", "7", "6", "0"}, {"1", "2", "7", "6", "5", "3", "4", "0"},
		{"1", "2", "7", "4", "0"}, {"1", "2", "7", "4", "3", "5", "6", "0"},
	}

	groupedPaths := groupUniquePaths(paths)

	for i, group := range groupedPaths {
		fmt.Printf("Group %d:\n", i+1)
		for _, path := range group {
			fmt.Printf("  %v\n", path)
		}
		fmt.Println()
	}
}
