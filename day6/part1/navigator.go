package main

const (
	RIGHT int = iota
	DOWN
	LEFT
	UP
)

type Navigator struct {
	data      []string
	position  [2]int
	direction int
}

func NewNavigator(data []string) *Navigator {
	return &Navigator{
		data: data,
	}
}

func (n *Navigator) SetDirection(direction int) {
	n.direction = direction
}

func (n *Navigator) SetPosition(x, y int) {
	n.position = [2]int{y, x}
}

func (n *Navigator) Position() (x, y int) {
	return n.position[1], n.position[0]
}

func (n *Navigator) RotateClockwise() {
	if n.direction+1 > 3 {
		n.direction = 0
	} else {
		n.direction += 1
	}
}

func (n *Navigator) RotateCounterClockwise() {
	if n.direction-1 < 0 {
		n.direction = 3
	} else {
		n.direction -= 1
	}
}

func (n *Navigator) String() string {
	return string(n.data[n.position[0]][n.position[1]])
}

func (n *Navigator) Next() bool {
	switch n.direction {
	case UP:
		n.position[0] -= 1
		if n.position[0] < 0 {
			return false
		}
	case DOWN:
		n.position[0] += 1
		if n.position[0] >= len(n.data) {
			return false
		}
	case RIGHT:
		n.position[1] += 1
		if n.position[1] >= len(n.data[n.position[0]]) {
			return false
		}
	case LEFT:
		n.position[1] -= 1
		if n.position[1] < 0 {
			return false
		}
	}

	return true
}

func (n *Navigator) Previous() bool {
	switch n.direction {
	case UP:
		n.position[0] += 1
		if n.position[0] >= len(n.data) {
			return false
		}
	case DOWN:
		n.position[0] -= 1
		if n.position[0] < 0 {
			return false
		}
	case RIGHT:
		n.position[1] -= 1
		if n.position[1] < 0 {
			return false
		}
	case LEFT:
		n.position[1] += 1
		if n.position[1] >= len(n.data[n.position[0]]) {
			return false
		}
	}

	return true
}
