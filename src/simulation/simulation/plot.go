package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	con "simulation/consensusInfo"
	ex "simulation/exception"
)

func drawResults(results map[int]*con.Results, mutation string) {
	sent := newTriple(results, "sent")
	received := newTriple(results, "received")
	sentPlot, err := plot.New()
	receivedPlot, err := plot.New()
	ex.CheckError(err)

	drawData(sentPlot, sent, "sent_"+mutation)
	drawData(receivedPlot, received, "received_"+mutation)
}

func drawData(plot *plot.Plot, data plotter.XYZs, label string) {
	plot.Title.Text = "#Messages " + label
	plot.X.Label.Text = "Node"
	plot.Y.Label.Text = "Node"

	bs := newBubbles(data)
	plot.Add(bs)

	err := plot.Save(20*vg.Inch, 20*vg.Inch, label+".png")
	ex.CheckError(err)
}

func newBubbles(data plotter.XYZs) *plotter.Scatter {
	sc, err := plotter.NewScatter(data)
	ex.CheckError(err)

	minZ, maxZ := math.Inf(1), math.Inf(-1)
	for _, xyz := range data {
		if xyz.Z > maxZ {
			maxZ = xyz.Z
		}
		if xyz.Z < minZ {
			minZ = xyz.Z
		}
	}

	sc.GlyphStyleFunc = func(i int) draw.GlyphStyle {
		c := color.RGBA{R: 25, B: 178, A: 255}
		minRadius, maxRadius := vg.Points(1), vg.Points(10)
		rng := maxRadius - minRadius
		_, _, z := data.XYZ(i)
		d := (z - minZ) / (maxZ - minZ)
		r := vg.Length(d)*rng + minRadius
		return draw.GlyphStyle{Color: c, Radius: r, Shape: draw.CircleGlyph{}}
	}

	return sc
}

func newTriple(data map[int]*con.Results, set string) plotter.XYZs {
	size := len(data) * len(data)
	triples := make(plotter.XYZs, size)
	index := 0

	for idSrc, results := range data {
		switch set {
		case "sent":
			index = fillTriple(idSrc, results.Sent, triples, index)
		case "received":
			index = fillTriple(idSrc, results.Received, triples, index)
		}
	}

	return triples
}

func fillTriple(idSrc int, data []float64, triples plotter.XYZs, index int) int {
	i := index

	for idDest, value := range data {
		triples[i].X = float64(idSrc)
		triples[i].Y = float64(idDest)
		triples[i].Z = value
		i++
	}

	return i
}
