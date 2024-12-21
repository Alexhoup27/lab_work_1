package main

import (
	"fmt"
	"math"
)

var A_l, B_l, C_l, D_l, E_l, A_g, B_g, C_g, D_g, E_g float64
var x, y float64

func eval_delta(A, B, C, D, E float64) float64 {
	return (math.Pow(C, 2) / (4 * A)) + (math.Pow(D, 2) / (4 * B)) - E
}

func is_intersect(y0_g, y0_l, A_g, A_l, B_g, B_l, delta float64) bool {
	if math.Abs((B_g+B_l)*(B_g*math.Pow(y0_g, 2)+B_l*math.Pow(y0_l, 2)-delta)-
		math.Pow(B_g*y0_g+B_l*y0_l, 2))/(-((B_g + B_l) * (A_g + A_l))) != 0 {
		return true
	}
	return false
}

// Проверка через дельу
func это_гипербола(A, B, C, D, E, a, b float64) bool {
	if a <= 0 || b <= 0 {
		return false
	}
	delta := eval_delta(A, B, C, D, E)
	if delta != 0 && A*B < 0 {
		return true
	}
	return false
}

func это_линии(A, B, C, D, E, a, b float64) bool {
	if a <= 0 || b <= 0 {
		return false
	}
	delta := eval_delta(A, B, C, D, E)
	if delta == 0 && A*B < 0 {
		return true
	}
	return false
}

func on_lines(A, B, C, D, E, x, y float64) bool {
	return A*math.Pow(x, 2)+B*math.Pow(y, 2)+C*x+D*y+E == 0
}

func on_hyperbola(A, B, C, D, E, x, y float64) bool {
	return A*math.Pow(x, 2)+B*math.Pow(y, 2)+C*x+D*y+E == 0
}

func under_hyperbola(a, b, x_0, y_0, x, y float64) bool {
	desc := math.Sqrt(math.Pow(2*a*y_0, 2) - 4*a*
		(a*math.Pow(y_0, 2)-b*math.Pow(x, 2)+2*x*x_0*b-b*math.Pow(x_0, 2)-a*b))
	if desc < 0 {
		return false
	}
	y_1 := (2*a*y_0 + desc) / (2 * a)
	y_2 := (2*a*y_0 - desc) / (2 * a)
	if y < y_1 && y < y_2 {
		return true
	}
	return false
}

func over_hyperbola(a, b, x_0, y_0, x, y float64) bool {
	desc := math.Sqrt(math.Pow(2*a*y_0, 2) - 4*a*
		(a*math.Pow(y_0, 2)-b*math.Pow(x, 2)+2*x*x_0*b-b*math.Pow(x_0, 2)-a*b))
	if desc < 0 {
		return false
	}
	y_1 := (2*a*y_0 + desc) / (2 * a)
	y_2 := (2*a*y_0 - desc) / (2 * a)
	if y > y_1 && y > y_2 {
		return true
	}
	return false
}

func over_under_lines(a, b, x_0, y_0, x, y float64) bool {
	desc := math.Sqrt(math.Pow(2*a*y_0, 2) + 4*a*
		(-a*math.Pow(y_0, 2)+b*math.Pow(x, 2)-2*x*x_0*b+b*math.Pow(x_0, 2)))
	if desc < 0 {
		return false
	}
	y_1 := (2*a*y_0 + desc) / (2 * a)
	y_2 := (2*a*y_0 - desc) / (2 * a)
	if y > math.Min(y_1, y_2) && y < math.Max(y_1, y_2) {
		return true
	}
	return false
}

func over_lines(a, b, x_0, y_0, x, y float64) bool {
	desc := math.Sqrt(math.Pow(2*a*y_0, 2) + 4*a*
		(-a*math.Pow(y_0, 2)+b*math.Pow(x, 2)-2*x*x_0*b+b*math.Pow(x_0, 2)))
	if desc < 0 {
		return false
	}
	y_1 := (2*a*y_0 + desc) / (2 * a)
	y_2 := (2*a*y_0 - desc) / (2 * a)
	if y > y_1 && y > y_2 {
		return true
	}
	return false
}
func correct_graphics(y0_l, y0_g, b_g, a_g, b_l, a_l float64) bool {
	if y0_l != y0_g-math.Pow(math.Abs(b_g), 0.5) {
		return false //проверка на правильное положение точки симметрии линей
	}
	if math.Abs(math.Abs(a_g-b_g)*2-math.Abs(a_l-b_l)*1.9) <= 2 {
		return false // проверка на пересечение графиков
	}
	return true
}

func new_correct_graphics(x0_l, y0_l, x0_g, y0_g, a_h, b_h, c_h, d_h, e_h, a_l, b_l float64) bool {
	if on_hyperbola(a_h, b_h, c_h, d_h, e_h, x0_l, y0_l) == false ||
		x0_l != x0_g || y0_l >= y0_g { // проверка на положение графиков относительно друг друга
		return false
	}
	delta := eval_delta(a_h, b_h, c_h, d_h, e_h)
	if is_intersect(y0_g, y0_l, a_h, a_l, b_h, b_l, delta) == false {
		return false
	}
	return true
}

func main() {
	fmt.Println("Enter A, B, C, D, E for intersecting lines:")
	fmt.Scan(&A_l, &B_l, &C_l, &D_l, &E_l)
	b_l := A_l
	a_l := -B_l
	x0_l := -C_l / (2 * b_l)
	y0_l := D_l / (2 * a_l)
	if это_линии(A_l, B_l, C_l, D_l, E_l, a_l, b_l) == false {
		fmt.Println("Not lines!!!")
	}
	fmt.Println("Enter A, B, C, D, E for hyperbola:")
	fmt.Scan(&A_g, &B_g, &C_g, &D_g, &E_g)
	b_g := -A_g
	a_g := B_g
	x0_g := C_g / (2 * b_g)
	y0_g := -D_g / (2 * a_g)
	if это_гипербола(A_g, B_g, C_g, D_g, E_g, a_l, b_l) == false {
		fmt.Println("Not hyperbola!!!")
	}

	if x0_g != x0_l || y0_l >= y0_g {
		fmt.Println("Wrong graphics !!!")
		return
	}
	if new_correct_graphics(x0_l, y0_l, x0_g, y0_g, A_g, B_g, C_g, D_g, E_g, A_l, B_l) == false {
		fmt.Println("Wrong graphics !!!")
		return
	}
	fmt.Println("Enter cords:")
	fmt.Scan(&x, &y)
	flag_l := on_lines(A_l, B_l, C_l, D_l, E_l, x, y)
	flag_h := on_hyperbola(A_g, B_g, C_g, D_g, E_g, x, y)
	if flag_l && flag_h {
		fmt.Println("Paint booth grafics")
	} else if flag_l {
		fmt.Println("Paint lines")
	} else if flag_h {
		fmt.Println("Paint hyperbola")
	} else if y < y0_l && x > x0_l {
		if over_under_lines(a_l, b_l, x0_l, y0_l, x, y) && under_hyperbola(a_g, b_g, x0_g, y0_g, x, y) {
			fmt.Println("Paint in white")
		} else {
			fmt.Println("Paint in blue")
		}
	} else if y > y0_l && x < x0_l {
		if over_under_lines(a_l, b_l, x0_l, y0_l, x, y) && over_hyperbola(a_g, b_g, x0_g, y0_g, x, y) {
			fmt.Println("Paint in white")
		} else {
			fmt.Println("Paint in blue")
		}
	} else if y > y0_l {
		if over_lines(a_l, b_l, x0_l, y0_l, x, y) &&
			(under_hyperbola(a_g, b_g, x0_g, y0_g, x, y) == false &&
				over_hyperbola(a_g, b_g, x0_g, y0_g, x, y) == false) {
			fmt.Println("Paint in white")
		} else {
			fmt.Println("Paint in blue")
		}
	} else {
		fmt.Println("Paint in blue")
	}
	return
}
