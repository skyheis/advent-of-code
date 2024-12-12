package utils

type Coordinates struct {
	X, Y int
}

func (c Coordinates) Up() Coordinates {
	return Coordinates{X: c.X, Y: c.Y - 1}
}

func (c Coordinates) Down() Coordinates {
	return Coordinates{X: c.X, Y: c.Y + 1}
}

func (c Coordinates) Left() Coordinates {
	return Coordinates{X: c.X - 1, Y: c.Y}
}

func (c Coordinates) Right() Coordinates {
	return Coordinates{X: c.X + 1, Y: c.Y}
}

func (c Coordinates) UpLeft() Coordinates {
	return Coordinates{X: c.X - 1, Y: c.Y - 1}
}

func (c Coordinates) UpRight() Coordinates {
	return Coordinates{X: c.X + 1, Y: c.Y - 1}
}

func (c Coordinates) DownRight() Coordinates {
	return Coordinates{X: c.X + 1, Y: c.Y + 1}
}

func (c Coordinates) DownLeft() Coordinates {
	return Coordinates{X: c.X - 1, Y: c.Y + 1}
}
