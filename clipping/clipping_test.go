package clipping

import "testing"

func TestOutcode(t *testing.T) {
	var out uint8

	out = outcode(100, 100)
	if out != in {
		t.Fatalf("expected %d, got %d", in, out)
	}
	out = outcode(650, 0)
	if out != right {
		t.Fatalf("expected %d, got %d", right, out)
	}
	out = outcode(300, -100)
	if out != btm {
		t.Fatalf("expected %d, got %d", btm, out)
	}
	out = outcode(-500, 300)
	if out != left {
		t.Fatalf("expected %d, got %d", left, out)
	}
	out = outcode(0, 900)
	if out != top {
		t.Fatalf("expected %d, got %d", top, out)
	}
}

func TestCohenSutherland(t *testing.T) {
	var clip bool

	var x1, y1, x2, y2 float32 = 450, 450, 550, 10
	clip = CohenSutherlandClipping2D(x1, y1, x2, y2)
	if !clip {
		t.Fatalf("CohenSutherlandClipping2D(%f, %f, %f, %f)=%t, got %t", x1, y1, x2, y2, true, false)
	}

	x1, y1, x2, y2 = 700, 50, 800, 100
	clip = CohenSutherlandClipping2D(x1, y1, x2, y2)
	if clip {
		t.Fatalf("CohenSutherlandClipping2D(%f, %f, %f, %f)=%t, got %t", x1, y1, x2, y2, false, true)
	}

	x1, y1, x2, y2 = -100, 400, 300, 700
	clip = CohenSutherlandClipping2D(x1, y1, x2, y2)
	if !clip {
		t.Fatalf("CohenSutherlandClipping2D(%f, %f, %f, %f)=%t, got %t", x1, y1, x2, y2, true, false)
	}
}
