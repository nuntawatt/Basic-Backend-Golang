package main

import (
	"fmt"
	"math"
)

// Shape interface ใครก็ได้ที่มี Area() ถือว่าเป็น Shape
type Shape interface {
	Area() float64
}

// Rectangle สี่เหลี่ยม
type Rectangle struct {
	Width  float64
	Height float64
}

// Area คำนวณพื้นที่สี่เหลี่ยม = กว้าง * สูง
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle วงกลม
type Circle struct {
	Radius float64
}

func main() {
	rect := Rectangle{Width: 5, Height: 3} // สร้างสี่เหลี่ยม
	circle := Circle{Radius: 7}            // สร้างวงกลม

	// พิมพ์พื้นที่ของทั้งสองรูป
	PrintArea(rect)
	PrintArea(circle)
}

// method Area คำนวณพื้นที่วงกลม
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius // พื้นที่ = π * r^2
}

// method PrintArea รับ Shape อะไรก็ได้ แล้วพิมพ์พื้นที่
func PrintArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area()) // เรียก method Area() ของ Shape ที่ส่งเข้ามา
}
