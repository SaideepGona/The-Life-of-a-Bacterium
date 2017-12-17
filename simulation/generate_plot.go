package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"image/color"
	"strconv"
)

func Plot(n plotter.XYs, title, xLabel, yLabel string) {

	// Create a new plot, set its title and axis labels.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	// Make a line plotter with points and set its style.
	lpLine, lpPoints, err := plotter.NewLinePoints(n)
	if err != nil {
		panic(err)
	}
	lpLine.Color = color.RGBA{G: 255, A: 255}
	lpPoints.Shape = draw.PyramidGlyph{}
	lpPoints.Color = color.RGBA{R: 255, A: 255}

	p.Add(lpLine, lpPoints)
	//	p.Legend.Add("line", l)

	// Save the plot to a PNG file.
	name := "data/" + title + ".png"
	if err := p.Save(4*vg.Inch, 4*vg.Inch, name); err != nil {
		panic(err)
	}
}

// To draw multiple lines of plot in single grid
func MultiPlot(n []plotter.XYs, title, xLabel, yLabel string) {

	// Create a new plot, set its title and
	// axis labels.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	l, err := plotter.NewLine(n[0])
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = MakeColor(0, 0, 0)
	p.Add(l)
	p.Legend.Add("Predator", l)

	for strain := 1; strain < len(n); strain++ {
		// Make a line plotter and set its style.
		l, err := plotter.NewLine(n[strain])
		if err != nil {
			panic(err)
		}
		l.LineStyle.Width = vg.Points(1)

		// assign different color for each strain of bacteria
		color := MakeColor(0, 0, 0)
		switch {
		case strain == 1:
			color = MakeColor(255, 0, 0)
		case strain == 2:
			color = MakeColor(0, 0, 255)
		case strain == 3:
			color = MakeColor(0, 255, 0)
		case strain == 4:
			color = MakeColor(255, 100, 255)
		case strain == 5:
			color = MakeColor(255, 255, 100)
		}

		l.LineStyle.Color = color

		p.Add(l)
		legend := "Strain " + strconv.Itoa(strain)
		p.Legend.Add(legend, l)
	}

	name := "data/" + title + ".png"
	if err := p.Save(4*vg.Inch, 4*vg.Inch, name); err != nil {
		panic(err)
	}
}
