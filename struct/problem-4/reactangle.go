package main

import "fmt"

type Rectangle struct {
	width float64
	height float64
}

func brandNewRectangle(width float64, height float64) *Rectangle {
	rect := Rectangle{
		width: width,
		height: height,
	}
	return &rect
}

func (rec Rectangle) areaOfReactangle() float64 {
	return rec.width * rec.height
}

func main() {
	myRoom := brandNewRectangle(10, 5)
	fmt.Println(*myRoom)
	fmt.Println(myRoom.areaOfReactangle())
}
