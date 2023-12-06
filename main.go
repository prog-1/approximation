package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

const (
	screenWidth  = 1800
	screenHeight = 1080
)

type game struct {
	points  plotter.XYs
	NewPlot func() *plot.Plot
	plot    *plot.Plot
}

func NewGame(points plotter.XYs) *game {
	return &game{
		points,
		func() *plot.Plot {
			p := plot.New()
			p.X.Min = -10
			p.X.Max = 10
			p.Y.Min = -10
			p.Y.Max = 10

			p.BackgroundColor = color.Black
			return p
		},
		nil,
	}
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.plot = g.NewPlot()

	s, err := plotter.NewScatter(g.points)
	if err != nil {
		return err
	}
	s.Color = color.RGBA{255, 0, 0, 255}
	g.plot.Add(s)

	f := plotter.NewFunction(Approximate(g.points))
	f.Color = color.RGBA{255, 255, 255, 255}
	g.plot.Add(f)

	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	DrawPlot(screen, g.plot)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Approximation line plotter")
	g := NewGame(plotter.XYs{{1, 1}, {2, 4}, {3, 1}, {4, 5}, {5, 3}})
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func DrawPlot(screen *ebiten.Image, p *plot.Plot) {
	// https://github.com/gonum/plot/wiki/Drawing-to-an-Image-or-Writer:-How-to-save-a-plot-to-an-image.Image-or-an-io.Writer,-not-a-file.
	img := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	c := vgimg.NewWith(vgimg.UseImage(img))
	p.Draw(draw.New(c))

	screen.DrawImage(ebiten.NewImageFromImage(c.Image()), &ebiten.DrawImageOptions{})
}

// Returns approximating linear function
func Approximate(points plotter.XYs) func(float64) float64 {
	return func(x float64) float64 { return 2*x + 1 }
}
