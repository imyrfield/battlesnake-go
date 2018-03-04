package main

func (p Point) IsOutOfBounds(board *MoveRequest) bool {
	if p.X < 0 || p.Y < 0 {
		return true
	}

	if p.X >= board.Width || p.Y >= board.Height {
		return true
	}

	return false
}

func (p Point) Equals(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p Point) IsInList(list PointList) bool {
	for _, point := range list {
		if p.Equals(point) {
			return true
		}
	}
	return false
}

func (p Point) SurroundingPoints() PointList {
	var list PointList

	list = append(list,
		Point{p.X - 1, p.Y},
		Point{p.X + 1, p.Y},
		Point{p.X, p.Y - 1},
		Point{p.X, p.Y + 1})

	return list
}

func (p Point) IsValidPoint(data *MoveRequest) bool {
	if p.IsOutOfBounds(data) {
		return false
	}

	if p.IsInList(OccupiedPoints(data)) {
			return false
	}
	return true
}

func (p Point) ValidSurroundingPoints(data *MoveRequest) PointList {
	var list PointList

	for _, point := range p.SurroundingPoints() {
		if point.IsValidPoint(data) {
			list = append(list, point)
		}
	}

	return list
}

func (p1 Point) DistanceTo(p2 Point) int {
	return abs(p1.X - p2.X) + abs(p1.Y - p2.Y)
}