package main

import (
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y int
}

type Maze struct {
	maze          [][]int
	start, finish Point
	bestWay       []Point
	bestWayWeight int
}

func (m *Maze) PrintMaze() {
	for _, row := range m.maze {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
}

func (m *Maze) PrintBestWay() {
	for _, point := range m.bestWay {
		fmt.Printf("%d %d\n", point.y, point.x)
	}
}

func (m *Maze) FindBestWay() {
	visited := make(map[Point]int)
	m.walk(m.start, make([]Point, 0), m.maze[m.start.y][m.start.x], &visited)
}

// walk - finds the best way using deep-search algorithm. It checks all possible ways from "cur" to finish.
// Stores best way in Maze.bestWay, changes if it finds a better one.
// visited map - [Point]WayWeight. If new way to visited point is cheaper - visits it again and updates it`s WayWeight.
// Otherwise, it skips visited point.
func (m *Maze) walk(cur Point, curWay []Point, curWeight int, visited *map[Point]int) {
	curWay = append(curWay, cur)
	curWeight += m.maze[cur.y][cur.x]

	//pre-check
	if cur == m.finish {
		if curWeight < m.bestWayWeight {
			m.bestWayWeight = curWeight
			m.bestWay = make([]Point, len(curWay))
			copy(m.bestWay, curWay)
		}
		//delete cur from curWay (necessary for different "iteration" of this recursion
		// (because they use same array of points. Because slice is a struct with a pointer)).
		curWay = curWay[0 : len(curWay)-1]
		return
		//I dont mark finish as visited point (otherwise it`ll be reached only once)
	}

	(*visited)[cur] = curWeight

	toCheck := Point{
		x: cur.x - 1,
		y: cur.y,
	}
	//check left
	if toCheck.x >= 0 && m.maze[toCheck.y][toCheck.x] > 0 {
		if (*visited)[toCheck] > curWeight || (*visited)[toCheck] == 0 {
			m.walk(toCheck, curWay, curWeight, visited)
		}
	}

	//check right
	toCheck.x = cur.x + 1
	//toCheck.y = cur.y - its already this value. I don`t want to delete it to better understanding this code.
	if toCheck.x <= len(m.maze[toCheck.y])-1 && m.maze[toCheck.y][toCheck.x] > 0 {
		if (*visited)[toCheck] > curWeight || (*visited)[toCheck] == 0 {
			m.walk(toCheck, curWay, curWeight, visited)
		}
	}

	//check top
	toCheck.x = cur.x
	toCheck.y = cur.y - 1
	if toCheck.y >= 0 && m.maze[toCheck.y][toCheck.x] > 0 {
		if (*visited)[toCheck] > curWeight || (*visited)[toCheck] == 0 {
			m.walk(toCheck, curWay, curWeight, visited)
		}
	}

	//check bottom
	//toCheck.x = cur.x - its already this value. I don`t want to delete it to better understanding this code.
	toCheck.x = cur.x
	toCheck.y = cur.y + 1
	if toCheck.y <= len(m.maze)-1 && m.maze[toCheck.y][toCheck.x] > 0 {
		if (*visited)[toCheck] > curWeight || (*visited)[toCheck] == 0 {
			m.walk(toCheck, curWay, curWeight, visited)
		}
	}

	curWay = curWay[0 : len(curWay)-1]
	return
}

func main() {
	//prepare
	height, wight := 0, 0
	_, err := fmt.Scanf("%d %d", &height, &wight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input error: %v\n", err)
		os.Exit(1)
	}

	//read maze
	maze := Maze{
		maze:          make([][]int, height),
		start:         Point{},
		finish:        Point{},
		bestWay:       make([]Point, 0),
		bestWayWeight: math.MaxInt64,
	}
	for i := range maze.maze {
		maze.maze[i] = make([]int, wight)
		for j := range maze.maze[i] {
			_, err := fmt.Scanf("%d", &maze.maze[i][j])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Input error: %v\n", err)
				os.Exit(1)
			}
		}
	}

	//read start and finish
	_, err = fmt.Scanf("%d %d %d %d", &maze.start.y, &maze.start.x, &maze.finish.y, &maze.finish.x)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Input error: %v\n", err)
		os.Exit(1)
	}

	//go
	maze.FindBestWay()
	maze.PrintBestWay()
	fmt.Print(".")
}
