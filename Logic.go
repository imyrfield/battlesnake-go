package main

const (
	UP    			= "up"
	RIGHT           = "right"
	DOWN            = "down"
	LEFT            = "left"
)
/* Primary move function - replace return function to alter game logic */
func getMove(data *MoveRequest) string {

	food := GetClosestFood(data, data.You.Head())

	if data.You.Length <= averageLength(data) {
		return getFoodMove(data)
	}

	if NeedFood(data, data.You.Head().DistanceTo(food)) {
		return getFoodMove(data)
	}

	return getLoopMove(data)
}

func getDirection(current, next Point) string {
	vertical := next.Y - current.Y
	horizontal := next.X - current.X

	if vertical == 0 {
		if horizontal > 0 {
			return RIGHT
		}
		return LEFT
	}

	if vertical < 0 {
		return UP
	}

	return DOWN
}


func getFoodMove(data *MoveRequest) string {
	surroundingPoints := data.You.Head().ValidSurroundingPoints(data)
	food := GetClosestFood(data, data.You.Head())
	distance := data.You.Head().DistanceTo(food)
	nextMove := surroundingPoints[0]
	for _, point := range surroundingPoints {
		if (point.DistanceTo(food) < distance) {
			nextMove = point
		}
	}

	return getDirection(data.You.Head(), nextMove)
}

func getLoopMove(data *MoveRequest) string {
	path := data.You.Head().getPathTo(data.You.Tail(), data)

	if len(path) > 2 {
		return getDirection(path[0], path[1])
	}

	return getFoodMove(data)
}

func GetClosestFood(data *MoveRequest, point Point) Point {
	closest := data.Food[0]
	closestDistance := point.DistanceTo(closest)
	mostSpace := 0

	for _, food := range data.Food {
		distance := point.DistanceTo(food)
		emptySpaces := FloodFill(food, data)

		if (distance < closestDistance && emptySpaces > data.You.Length + 3) || emptySpaces > mostSpace {
			closest = food
			closestDistance = distance
			mostSpace = emptySpaces
		}
	}

	return closest
}

func NeedFood(data *MoveRequest, distanceToFood int) bool {
	buffer := 25
	if data.You.Health <= distanceToFood+buffer {
		return true
	}
	return false
}

func OccupiedPoints(data *MoveRequest) PointList {
	var list PointList

	for _, snake := range data.Snakes{
		if snake.Health == 0 {
			continue
		}

		if snake.Health == 100 {
			list = append(list, snake.Body...)
		} else {
			list = append(list, snake.Body[:len(snake.Body) - 1]...)
		}

	}

	return list
}

func FloodFill(point Point, data *MoveRequest) int {
	board := make([][]int, data.Width)

	for row := range board {
		board[row] = make([]int, data.Height)
	}

	return fillBoard(point, data, board, 0)
}

func fillBoard(point Point, data *MoveRequest, board [][]int, numPoints int) int {
	if point.IsValidPoint(data) && board[point.X][point.Y] == 0 {
		// mark as visited
		board[point.X][point.Y] = 1
		numPoints++

		for _, neighbor := range point.ValidSurroundingPoints(data) {
			numPoints = fillBoard(neighbor, data, board, numPoints)
		}
	}

	return numPoints
}

func (p1 Point) getPathTo(p2 Point, data *MoveRequest) PointList {
	visited := make(PointList, 0)
	toVisit := make(PointList, 0)
	toVisit = append(toVisit, p1)
	cameFrom := make(map[Point]Point)
	fScore := make(map[Point]int)
	gScore := make(map[Point]int)

	for i := 0; i < data.Width; i++ {
		for j:= 0; j < data.Height; j++ {
			fScore[Point{i,j}] = 1000
			gScore[Point{i,j}] = 1000
		}
	}

	gScore[p1] = 0
	fScore[p1] = p1.DistanceTo(p2)

	for len(toVisit) > 0 {
		min := toVisit[0]
		minIndex := 0

		for i := 0; i < len(toVisit); i++ {
			if fScore[toVisit[i]] < fScore[min] {
				min = toVisit[i]
				minIndex = i
			}
		}

		if min.Equals(p2) {
			return path(p2, cameFrom)
		}

		toVisit[minIndex] = toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]
		visited = append(visited, min)

		neighbors := min.ValidSurroundingPoints(data)

		for _, p := range neighbors {
			if p.IsInList(visited) {
				continue
			}

			s := gScore[min] + p.DistanceTo(min)

			if !p.IsInList(toVisit) {
				toVisit = append(toVisit, p)
			} else if s >= gScore[p] {
				continue
			}

			cameFrom[p] = min

			gScore[p] = s
			fScore[p] = s + p.DistanceTo(min)
		}
	}

	return nil
}

func path(p Point, path map[Point]Point) PointList {
	list := make(PointList, 0)
	list = append(list, p)

	_, exists := path[p]

	for ; exists; _,exists = path[p] {
		p = path[p]
		list = append(list, p)
	}

	return reversedPath(list)
}

func reversedPath(list PointList) PointList {
	for i:= 0; i < len(list)/2; i++ {
		j := len(list) - i - 1
		list[i], list[j] = list[j], list[i]
	}

	return list
}

func averageLength(data *MoveRequest) int {
	sum := 0
	for _, snake := range data.Snakes {
		sum += snake.Length
	}
	return sum / len(data.Snakes)
}