package clipping

const top uint8 = 8
const btm uint8 = 4
const right uint8 = 2
const left uint8 = 1
const in uint8 = 0

// 600x600 clipping window
const xmin = 0
const xmax = 600
const ymin = 0
const ymax = 600

func outcode(x float32, y float32) uint8 {
	out := in
	if x < xmin {
		out |= left
	}

	if x > xmax {
		out |= right
	}

	if y < ymin {
		out |= btm
	}

	if y > ymax {
		out |= top
	}

	return out
}

func CohenSutherlandClipping2D(x1 float32, y1 float32, x2 float32, y2 float32) bool {
	o1 := outcode(x1, y1)
	o2 := outcode(x2, y2)

	// both inside
	if o1|o2 == 0 { // o1 = o2 = 0
		return true
	}
	// outside, and on the same side of the window
	if o1&o2 != 0 {
		return false
	}

	// we get past these checks, we need to compute intersections

	outside := o1
	if o2 > o1 {
		outside = o2
	}
	var x, y float32
	if outside&top != 0 {
		x = x1 + (x2-x1)*(ymax-y1)/(y2-y1)
		y = ymax
	} else if outside&btm != 0 {
		x = x1 + (x2-x1)*(ymax-y1)/(y2-y1)
		y = ymin
	} else if outside&right != 0 {
		y = y1 + (y2-y1)*(xmax-x1)/(x2-x1)
		x = xmax
	} else if outside&left != 0 {
		y = y1 + (y2-y1)*(xmax-x1)/(x2-x1)
		x = xmin
	}

	if outside == o1 {
		return CohenSutherlandClipping2D(x, y, x2, y2)
	} else {
		return CohenSutherlandClipping2D(x1, y1, x, y)
	}
}
