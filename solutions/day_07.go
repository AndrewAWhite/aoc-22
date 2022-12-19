package solutions

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	Name     string
	Children []*File
	Parent   *File
	Size     int
	IsDir    bool
	Total    int
}

func parseInput() [][]string {
	input := readInput("./input/day_07/p1.txt")
	cmdSplit := strings.Split(input, "\n$")
	parsedInput := make([][]string, len(cmdSplit))
	for i, cmdFull := range cmdSplit {
		termLines := strings.Split(cmdFull, "\n")
		termLines[0] = strings.Trim(termLines[0], " \t")
		parsedInput[i] = termLines
	}
	return parsedInput
}

func calcSizeWalk(dir *File) int {
	for _, file := range dir.Children {
		dir.Total += file.Size
		if file.Children != nil {
			dir.Total += calcSizeWalk(file)
		}
	}
	return dir.Total
}

func parseCmd(cmdSet []string) (string, string) {
	cmd := cmdSet[0]
	cmdSplit := strings.Split(cmd, " ")
	cmd = cmdSplit[0]
	cmdArg := ""
	if len(cmdSplit) > 1 {
		cmdArg = cmdSplit[1]
	}
	return cmd, cmdArg
}

func buildFsModel() *File {
	parsedInput := parseInput()
	// Create the root directory
	root := &File{Name: "/", IsDir: true}
	ctx := root
	for _, cmdSet := range parsedInput {
		cmd, cmdArg := parseCmd(cmdSet)
		switch cmd {
		case "ls":
			{
				for _, file := range cmdSet[1:] {
					fSplit := strings.Split(file, " ")
					if fSplit[0] == "dir" {
						d := &File{Name: fSplit[1], Children: []*File{}, Parent: ctx, IsDir: true}
						ctx.Children = append(ctx.Children, d)
					} else {
						size, err := strconv.Atoi(fSplit[0])
						if err != nil {
							fmt.Printf("ERROR PARSING FILE SIZE: %s\n", fSplit[0])
							os.Exit(1)
						}
						f := &File{Name: fSplit[1], Size: size}
						ctx.Children = append(ctx.Children, f)
					}
				}
			}
		case "cd":
			{
				if cmdArg == ".." {
					ctx = ctx.Parent
				} else if cmdArg == "/" {
					ctx = root
				} else {
					for _, tgt := range ctx.Children {
						if tgt.Name == cmdArg {
							ctx = tgt
							break
						}
					}
				}
			}
		}
	}
	calcSizeWalk(root)
	return root
}

func findCandidateTotal(dir *File, cdtSize int) int {
	candidateTotal := 0
	for _, file := range dir.Children {
		if file.IsDir && file.Total <= cdtSize {
			candidateTotal += file.Total
		}
		if file.Children != nil {
			candidateTotal += findCandidateTotal(file, cdtSize)
		}
	}
	return candidateTotal
}

func filterDirs(dir *File, candidates *[]*File, minTotal int) {
	for _, file := range dir.Children {
		if file.IsDir {
			if file.Total >= minTotal {
				*candidates = append(*candidates, file)
			}
		}
		if file.Children != nil {
			filterDirs(file, candidates, minTotal)
		}
	}
}

func Solution_07_1() int {
	root := buildFsModel()
	return findCandidateTotal(root, 100000)
}

func Solution_07_2() int {
	root := buildFsModel()
	freeDisk := 70000000 - root.Total
	needed := 30000000 - freeDisk
	candidates := &[]*File{}
	filterDirs(root, candidates, needed)
	cd := *candidates
	sort.Slice(cd, func(i, j int) bool {
		return cd[i].Total < cd[j].Total
	})
	return cd[0].Total
}
