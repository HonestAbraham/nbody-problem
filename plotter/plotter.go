package plotter

import (
	"fmt"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"nbody-problem/enum"
)

func PlotBodies(bodies []enum.Body, filename string) error {
	// Ensure directory exists
	dir := "plotter/plots"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Full path inside directory
	path := fmt.Sprintf("%s/%s", dir, filename)

	// Create plot
	p := plot.New()
	p.Title.Text = "N-Body Snapshot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(bodies))
	for i, b := range bodies {
		pts[i].X = b.X
		pts[i].Y = b.Y
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	p.Add(s)

	// Save
	if err := p.Save(6*vg.Inch, 6*vg.Inch, path); err != nil {
		return err
	}

	fmt.Println("Saved:", path)
	return nil
}
