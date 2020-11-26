package main

import (
	"fmt"
)

// type Vertex struct {
// 	X, Y, Z float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z)
// }

// func main() {
// 	v := Vertex{3, 4, 1}
// 	fmt.Println(v.Abs())
// }
// type qwe struct {
// 	X, Y float64
// }

// func abs(v qwe) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := qwe{3, 4}
// 	fmt.Println(abs(v))
// }
// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// func main() {
// 	f := MyFloat(-math.Sqrt2)
// 	fmt.Println(f.Abs())
// }

// type Vertex struct {
// 	X, Y float64
// }

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func Scale(v *Vertex, f float64) {
// 	fmt.Println(f)
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	Scale(&v, 100)
// 	fmt.Println(Abs(v))
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func ScaleFunc(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	v.Scale(2)
// 	fmt.Println(v)
// 	ScaleFunc(&v, 10)
// 	fmt.Println(v)

// 	fmt.Println("p")

// 	p := &Vertex{4, 3}
// 	p.Scale(3)
// 	fmt.Println(p)
// 	ScaleFunc(p, 8)
// 	fmt.Println(p)

// 	fmt.Println("res")
// 	fmt.Println(v, p)
// }
// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func ScaleFunc(v *Vertex, f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v)
// 	v.Scale(2)
// 	ScaleFunc(&v, 10)
// 	fmt.Println(v)
// 	v = Vertex{3, 4}
// 	fmt.Println(v)
// 	p := &v
// 	p.Scale(2)
// 	ScaleFunc(p, 10)

// 	fmt.Println(v, p)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func AbsFunc(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// 	fmt.Println(AbsFunc(v))

// 	p := &v
// 	fmt.Println(p.Abs())
// 	fmt.Println(AbsFunc(*p))
// }
// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := &Vertex{3, 4}
// 	fmt.Printf("Before scaling: %v, Abs: %v\n", v, v.Abs())
// 	v.Scale(5)
// 	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
// }
// type Abser interface {
// 	Abs() float64
// }

// func main() {
// 	// var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	fmt.Println("f", f)
// 	v := Vertex{5, 4}
// 	fmt.Println("v", v)

// 	// a = f // a MyFloat implements Abser
// 	// fmt.Println("a", a)
// 	a := &v // a *Vertex implements Abser
// 	fmt.Println("a", a)
// 	// In the following line, v is a Vertex (not *Vertex)
// 	// and does NOT implement Abser.
// 	// a = v

// 	fmt.Println(a.Abs())
// 	fmt.Println(f.Abs())
// }

// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (a *Vertex) Abs() float64 {
// 	return math.Sqrt(a.X*a.X + a.Y*a.Y)
// }
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (k T) M() {
	fmt.Println(k.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
