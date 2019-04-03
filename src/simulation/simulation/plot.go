package main

import (
    "image/color"
    "fmt"
    "strconv"
    "math"
    "sort"
    "time"
    "os"

    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
    "gonum.org/v1/plot/vg/draw"
    con "simulation/consensus"
    ex "simulation/exception"
)

const (
    // DIRCSV is the directory for csv datasets
    DIRCSV = "results/csv-datasets/"
    // DIRPNG is the directory for plots in png formate
    DIRPNG = "results/png-plots/"
)

// INTERVALS is the constant that indicates the time interval in the cumulative graph
var INTERVALS = float64(50)

// coordinates is a struct that represents the 2D coordinates
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
    plot.Title.Text = "# Messages " + label
    plot.X.Label.Text = "Node"
    plot.Y.Label.Text = "Node"

    bubbles := newBubbles(data)
    plot.Add(bubbles)

    err := plot.Save(20 * vg.Inch, 20 * vg.Inch, DIRPNG + label + ".png")
    ex.CheckError(err)
}

func drawTimeData(plot *plot.Plot, data plotter.XYs, mutation string) {
    plot.Title.Text = "Number of Nodes that decide, over time"
    plot.X.Label.Text = "Time"
    plot.Y.Label.Text = "# Nodes"

    err := plotutil.AddLines(plot, mutation, data)
    err = plot.Save(8 * vg.Inch, 8 * vg.Inch, DIRPNG + "time_" + mutation + ".png")

    ex.CheckError(err)
}

func newBubbles(data plotter.XYZs) *plotter.Scatter {
    scatter, err := plotter.NewScatter(data)
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

    scatter.GlyphStyleFunc = func(i int) draw.GlyphStyle {
        color := color.RGBA{ R: 25, B: 178, A: 255 }
        minRadius, maxRadius := vg.Points(0), vg.Points(10)

        rangeRadius := maxRadius - minRadius
        _, _, z := data.XYZ(i)
        diameter := (z - minZ) / (maxZ - minZ)
        radius := vg.Length(diameter) * rangeRadius + minRadius

        return draw.GlyphStyle{ Color: color, Radius: radius, Shape: draw.CircleGlyph{} }
    }

    return scatter
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

func save(data map[int]*con.Results, mutation string) {
    file, err := os.Create(DIRCSV +
        mutation                                       + "_" +
        fmt.Sprintf("%f", defaultDelta)                + "_" +
        strconv.Itoa(maxTries)                         + "_" +
        strconv.FormatFloat(percentageMiss,'f', 2, 64) + "_" +
        ".csv")

    ex.CheckError(err)
    defer file.Close()

    for _, results := range data {
        length := len(results.Delays)

        for _, result := range results.Delays[:(length - 1)] {
            value := ensureValue(result)
            file.WriteString(value + ", ")
        }

        lastResult := results.Delays[length - 1]
        value := ensureValue(lastResult)
        file.WriteString(value + "\n")
    }
}

func ensureValue(value float64) string {
    var result string

    if value < 0 {
        result = "0"
    } else {
        result = fmt.Sprintf("%f", value)
    }

    return result
}

func saveResult(endTime time.Time, startTime time.Time) {
    file, err := os.OpenFile(DIRCSV + "global_results.csv", os.O_APPEND|os.O_WRONLY, 0666)
    ex.CheckError(err)
    defer file.Close()

    // file.WriteString("#Nodes, default delta, max tries, percentage of miss (messages), mutation, result time")

    duration := float64(endTime.Sub(startTime)) / float64(time.Millisecond)

    file.WriteString(strconv.Itoa(numberParticipants) + "," +
        fmt.Sprintf("%f", defaultDelta)     + "," +
        strconv.Itoa(maxTries)              + "," +
        fmt.Sprintf("%f", percentageMiss)   + "," +
        fmt.Sprintf("%f", percentageFaults) + "," +
        mutation                            + "," +
        strconv.FormatBool(withFaults)      + "," +
        fmt.Sprintf("%f", duration)         + "\n")
}