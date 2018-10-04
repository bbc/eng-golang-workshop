package main

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel Struct {
	Circle  Circle
	Spokes int

}
