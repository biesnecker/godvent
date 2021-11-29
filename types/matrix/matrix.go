package matrix

import "fmt"

type Matrix struct {
	height, width int
	data          []int
}

func (m *Matrix) index(x, y int) int {
	return x*m.width + y
}

func SquareMatrix(side int) *Matrix {
	return RectangularMatrix(side, side)
}

func RectangularMatrix(height, width int) *Matrix {
	data := make([]int, height*width)
	return &Matrix{height: height, width: width, data: data}
}

func (m *Matrix) Get(x, y int) int {
	return m.data[m.index(x, y)]
}

func (m *Matrix) Set(x, y, value int) {
	m.data[m.index(x, y)] = value
}

func (m *Matrix) Transpose() {
	newMat := RectangularMatrix(m.width, m.height)
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			newMat.data[newMat.index(j, i)] = m.data[m.index(i, j)]
		}
	}
	*m = *newMat
}

func (m *Matrix) MirrorVertical() {
	for col := 0; col < m.width/2; col++ {
		for row := 0; row < m.height; row++ {
			m.data[m.index(row, col)],
				m.data[m.index(row, m.width-1-col)] =
				m.data[m.index(row, m.width-1-col)],
				m.data[m.index(row, col)]
		}
	}
}

func (m *Matrix) rotateInPlace() {
	var a, b, c, d int
	n := m.height

	for i := 0; i <= n/2-1; i++ {
		for j := 0; j <= n-(2*i)-2; j++ {
			a = m.data[m.index(i+j, i)]
			b = m.data[m.index(n-1-i, i+j)]
			c = m.data[m.index(n-1-i-j, n-1-i)]
			d = m.data[m.index(i, n-1-i-j)]

			m.data[m.index(i+j, i)] = b
			m.data[m.index(n-1-i, i+j)] = c
			m.data[m.index(n-1-i-j, n-1-i)] = d
			m.data[m.index(i, n-1-i-j)] = a
		}
	}
}

func (m *Matrix) Rotate() {
	if m.height == m.width {
		m.rotateInPlace()
	} else {
		m.Transpose()
		m.MirrorVertical()
	}
}

func (m *Matrix) Copy() *Matrix {
	n := RectangularMatrix(m.height, m.width)
	copy(n.data, m.data)
	return n
}

func (m *Matrix) DebugPrint() {
	for x := 0; x < m.height; x++ {
		for y := 0; y < m.width; y++ {
			fmt.Print(m.data[m.index(x, y)], "\t")
		}
		fmt.Print("\n")
	}
}
