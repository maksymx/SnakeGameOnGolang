package bll

type frame struct {
	figure
}
type iframe interface {
	Draw()
	isHitPoint(point ipoint) bool
}

func newIFrame(h, w int, value rune, writer IWriter) iframe {

	points := make([]ipoint, 0)
	points = addHorizontal(w, 0, points, value, writer)
	points = addHorizontal(w, h, points, value, writer)
	points = addVertical(h, 0, points, value, writer)
	points = addVertical(h, w, points, value, writer)

	return frame{figure{points}}
}

func addHorizontal(w, y int, points []ipoint, value rune, writer IWriter) []ipoint {
	for i := 0; i <= w; i++ {
		points = append(points, newIPoint(i, y, value, writer))
	}
	return points
}

func addVertical(h, x int, points []ipoint, value rune, writer IWriter) []ipoint {
	for i := 0; i <= h; i++ {
		points = append(points, newIPoint(x, i, value, writer))
	}
	return points
}

func (this frame) isHitPoint(point ipoint) bool {
	for _, p := range this.points {
		if p.Overlaps(point) {
			return true
		}
	}
	return false
}