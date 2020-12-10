package solutions

type Coordinate struct {
	x int
	y int
}

func (coord *Coordinate) moveRightBy(points int) {
	coord.x = coord.x + points
}

func (coord *Coordinate) moveDownBy(points int) {
	coord.y = coord.y + points
}

func getVisitedCoordinates(endingYCoordinate int, slopesToTraverse *Coordinate) []Coordinate {
	coord := Coordinate{
		x: 0,
		y: 0,
	}
	var result []Coordinate = make([]Coordinate, endingYCoordinate)
	index := 0
	for {
		if coord.y == endingYCoordinate {
			break
		}
		coord.moveRightBy(slopesToTraverse.x)
		coord.moveDownBy(slopesToTraverse.y)
		result[index] = coord
		index++
	}
	return result
}

func getCharAt(line string, position int) string {
	strlen := len(line)
	newPosition := position % strlen
	return string(line[newPosition])
}

func isTree(charAtSquare string) bool {
	return charAtSquare == "#"
}

func traverseSlopeWithTreeCount(input []string, slopeToTraverse *Coordinate) int {
	endingYCoordinate := len(input) - 1
	visitedCoordinates := getVisitedCoordinates(endingYCoordinate, slopeToTraverse)
	treesSeen := 0
	for _, coord := range visitedCoordinates {
		characterAt := getCharAt(input[coord.y], coord.x)
		if isTree(characterAt) {
			treesSeen++
		}
	}
	return treesSeen
}

func Day03Part1(input []string) int {
	slopeToTraverse := Coordinate{
		x: 3,
		y: 1,
	}
	return traverseSlopeWithTreeCount(input, &slopeToTraverse)
}

func Day03Part2(input []string) int {
	slopesToTraverse := []Coordinate{
		{
			x: 1,
			y: 1,
		},
		{
			x: 3,
			y: 1,
		},
		{
			x: 5,
			y: 1,
		},
		{
			x: 7,
			y: 1,
		},
		{
			x: 1,
			y: 2,
		},
	}
	result := 1
	for _, slopeToTraverse := range slopesToTraverse {
		result *= traverseSlopeWithTreeCount(input, &slopeToTraverse)
	}
	return result
}
