package solutions

type Coordinate struct {
	x int
	y int
}

func getVisitedCoordinates(endingYCoordinate int) []Coordinate {
	var result []Coordinate = make([]Coordinate, endingYCoordinate)
	for i := 1; i <= endingYCoordinate; i++ {
		coord := Coordinate{
			x: i * 3,
			y: i,
		}
		result[i-1] = coord
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

func Day03Part1(input []string) int {
	endingYCoordinate := len(input) - 1
	visitedCoordinates := getVisitedCoordinates(endingYCoordinate)
	treesSeen := 0
	for _, coord := range visitedCoordinates {
		characterAt := getCharAt(input[coord.y], coord.x)
		if isTree(characterAt) {
			treesSeen++
		}
	}
	return treesSeen
}
