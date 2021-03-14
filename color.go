package identicon

import "math"

type HSV struct {
	H, S, V float64
}

type RGB struct {
	R, G, B float64
}

func (c *HSV) RGB() *RGB {
	var r, g, b float64
	if c.S == 0 {
		r = c.V * 255
		g = c.V * 255
		b = c.V * 255
	} else {
		h := c.H * 6
		if h == 6 {
			h = 0
		}
		i := math.Floor(h)
		v1 := c.V * (1 - c.S)
		v2 := c.V * (1 - c.S*(h-i))
		v3 := c.V * (1 - c.S*(1-(h-i)))

		switch i {
		case 0:
			r = c.V
			g = v3
			b = v1
		case 1:
			r = v2
			g = c.V
			b = v1
		case 2:
			r = v1
			g = c.V
			b = v3
		case 3:
			r = v1
			g = v2
			b = c.V
		case 4:
			r = v3
			g = v1
			b = c.V
		default:
			r = c.V
			g = v1
			b = v2
		}

		r = r * 255
		g = g * 255
		b = b * 255
	}
	rgb := &RGB{r, g, b}
	return rgb
}
