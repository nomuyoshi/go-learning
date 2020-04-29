package geometry

import "math"

// Point は座標
type Point struct{ X, Y float64 }

// Distance は関数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance はPoint型のメソッド
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path は点のスライス、点を直線で結びつける道のりを表す
type Path []Point

// Distance はpathに沿って進んだ距離を返す。
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}

	return sum
}
