package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Dir struct {
	name     string
	parent   *Dir
	size     int
	files    map[string]struct{}
	children map[string]*Dir
}

func calcSizes(d *Dir) int {
	if len(d.children) == 0 {
		return d.size
	}
	children_total_size := 0
	for _, child := range d.children {
		children_total_size += calcSizes(child)
	}
	d.size += children_total_size
	return d.size
}

func getComplyingDirs1(d *Dir, treshold int) []*Dir {
	complyingDirs := []*Dir{}
	if d.size <= treshold {
		complyingDirs = append(complyingDirs, d)
	}
	for _, child := range d.children {
		complyingDirs = append(complyingDirs, getComplyingDirs1(child, treshold)...)
	}
	return complyingDirs
}

func getComplyingDirs2(d *Dir, treshold int) []*Dir {
	complyingDirs := []*Dir{}
	if d.size >= treshold {
		complyingDirs = append(complyingDirs, d)
		for _, child := range d.children {
			complyingDirs = append(complyingDirs, getComplyingDirs2(child, treshold)...)
		}
	}
	return complyingDirs
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	rootDir := &Dir{name: "/", parent: nil, size: 0, files: map[string]struct{}{}, children: map[string]*Dir{}}
	var curDir *Dir
	for _, line := range strings.Split(input, "\n") {
		splitLine := strings.Split(line, " ")
		if splitLine[0] == "$" {
			if splitLine[1] == "cd" {
				if splitLine[2] == ".." {
					curDir = curDir.parent
					continue
				}
				if splitLine[2] == "/" {
					curDir = rootDir
					continue
				}
				curDir = curDir.children[splitLine[2]]
			}
			continue
		}
		if splitLine[0] == "dir" {
			if _, ok := curDir.children[splitLine[1]]; !ok {
				curDir.children[splitLine[1]] = &Dir{name: splitLine[1], parent: curDir, size: 0, files: map[string]struct{}{}, children: map[string]*Dir{}}
			}
			continue
		}
		if _, ok := curDir.files[splitLine[1]]; !ok {
			curDir.files[splitLine[1]] = struct{}{}
			fileSize, _ := strconv.Atoi(splitLine[0])
			curDir.size += fileSize
		}
	}
	calcSizes(rootDir)
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		treshold := 30000000 - (70000000 - rootDir.size)
		complyingDirs := getComplyingDirs2(rootDir, treshold)
		return slices.MinFunc(complyingDirs, func(dL *Dir, dR *Dir) int { return dL.size - dR.size }).size
	}
	// solve part 1 here
	complyingDirs := getComplyingDirs1(rootDir, 100000)
	sum := 0
	for _, dir := range complyingDirs {
		sum += dir.size
	}
	return sum
}
