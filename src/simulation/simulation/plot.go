package main

import (
	"image/color"
    "math"
    "sort"
    "time"

    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	con "simulation/consensusInfo"
    ex "simulation/exception"
)

// INTERVALS is the constant that indicates the time interval in the cumulative graph
var INTERVALS = float64(50)

// XY is a struct
type coordinates struct {
    x float64
    y float64
}

func newCoordinates(x float64, y float64) *coordinates {
    coordinates := new(coordinates)
    coordinates.x = x
    coordinates.y = y

    return coordinates
}

func drawResults(results map[int]*con.Results, startTime time.Time, mutation string) {
	sent := newTriple(results, "sent")
    received := newTriple(results, "received")
    time := newCumulativeData(results, startTime)

	sentPlot, err := plot.New()
    receivedPlot, err := plot.New()
    timePlot, err := plot.New()
    
	ex.CheckError(err)

	drawData(sentPlot, sent, "sent_" + mutation)
    drawData(receivedPlot, received, "received_" + mutation)
    drawTimeData(timePlot, time, mutation)
}

func drawData(plot *plot.Plot, data plotter.XYZs, label string) {
	plot.Title.Text = "#Messages " + label
	plot.X.Label.Text = "Node"
	plot.Y.Label.Text = "Node"

	bubbles := newBubbles(data)
	plot.Add(bubbles)

	err := plot.Save(20 * vg.Inch, 20 * vg.Inch, label + ".png")
	ex.CheckError(err)
}

func drawTimeData(plot *plot.Plot, data plotter.XYs, mutation string) {
    plot.Title.Text = "Number of Nodes that decide over time"
	plot.X.Label.Text = "Time"
    plot.Y.Label.Text = "# Nodes"
    
    err := plotutil.AddLines(plot, mutation, data)
    err = plot.Save(8 * vg.Inch, 8 * vg.Inch, "time_" + mutation + ".png")

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
		c := color.RGBA{ R: 25, B: 178, A: 255 }
		minRadius, maxRadius := vg.Points(0), vg.Points(10)
        
        rng := maxRadius - minRadius
		_, _, z := data.XYZ(i)
		d := (z - minZ) / (maxZ - minZ)
		r := vg.Length(d) * rng + minRadius
        
        return draw.GlyphStyle{ Color: c, Radius: r, Shape: draw.CircleGlyph{} }
	}

	return sc
}

func newTriple(data map[int]*con.Results, set string) plotter.XYZs {
	size := len(data) * len(data)
	triples := make(plotter.XYZs, size)
	index := 0

	for idSource, results := range data {
		switch set {
		case "sent":
			index = fillTriple(idSource, results.Sent, triples, index)
		case "received":
			index = fillTriple(idSource, results.Received, triples, index)
		}
	}

	return triples
}

func fillTriple(idSource int, data []float64, triples plotter.XYZs, index int) int {
	i := index

	for idDestination, value := range data {
		triples[i].X = float64(idSource)
		triples[i].Y = float64(idDestination)
		triples[i].Z = value
		i++
	}

	return i
}

func newCumulativeData(data map[int]*con.Results, startTime time.Time) plotter.XYs {
    coordinates := make([]*coordinates, 0)
    durationsList := calculateDurationsList(data, startTime)

    maxDuration := durationsList[len(durationsList) - 1].Seconds()
    interval := float64(maxDuration / INTERVALS)
    numberNodes := float64(0)
    time := float64(0)

    for _, duration := range durationsList {
        for {
            if float64(duration.Seconds()) <= time {
                numberNodes++
                break     
            } else {
                coordinates = append(coordinates, newCoordinates(time, numberNodes))
                time += interval
            }
        }
    }
    
    coordinates = append(coordinates, newCoordinates(time, numberNodes)) 
    timeList := make(plotter.XYs, len(coordinates))

    for index, value := range coordinates {
        timeList[index].X = value.x
        timeList[index].Y = value.y
    }

    return timeList
}

func calculateDurationsList(data map[int]*con.Results, startTime time.Time) con.DurationSlice {
    list := make(con.DurationSlice, 0)

    for _, results := range data {
        duration := results.DecisionTime.Sub(startTime)
        list = append(list, duration)
    }

    sort.Sort(list)

    return list
}