package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 960
	screenHeight = 720
)

type Point struct {
	x, y float64
}

type Game struct {
	width, height int
	points        []Point
}

func NewGame(width, height int, p []Point) *Game {
	return &Game{
		width:  width,
		height: height,
		points: p,
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.points {
		vector.DrawFilledCircle(screen, float32(v.x), float32(v.y), 3, color.RGBA{255, 0, 0, 255}, false)
	}
	a, b := approximation(g.points)
	vector.StrokeLine(screen, 0, float32(b), float32(g.width), float32(a*float64(g.width)+b), 3, color.White, false)
}

func approximation(p []Point) (a, b float64) {
	var sumX, sumY, sumXY, sumX2 float64
	for _, v := range p {
		sumX += v.x
		sumY += v.y
		sumXY += v.x * v.y
		sumX2 += v.x * v.x
	}
	a = (float64(len(p))*sumXY - sumX*sumY) / (float64(len(p))*sumX2 - sumX*sumX)
	b = (sumY - a*sumX) / float64(len(p))
	return
}

func main() {
	var p []Point
	var n int
	var x, y float64
	fmt.Println("Enter number of points: ")
	fmt.Scan(&n)
	fmt.Println("Enter coordinates of points: ")
	for ; n > 0; n-- {
		fmt.Scan(&x, &y)
		p = append(p, Point{x, y})
	}
	g := NewGame(screenWidth, screenHeight, p)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
