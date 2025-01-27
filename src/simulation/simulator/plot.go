package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
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
	// FONTSIZE gives an entry height of 10.
	FONTSIZE = 10.189054726368159
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

	drawData(sentPlot, sent, "sent_" + mutation, mutation)
	drawData(receivedPlot, received, "received_" + mutation, mutation)
	drawTimeData(timePlot, time, mutation)
}

func drawData(plot *plot.Plot, data plotter.XYZs, label string, mutation string) {
	font, err := vg.MakeFont("Helvetica-Bold", FONTSIZE * 1.2)

	if err != nil {
		println("failed to create font: %v", err)
	}

	if mutation == "old_ring" {
		plot.Title.Text = "Ring"
	}else {
		plot.Title.Text = strings.Title(mutation)
	}

	plot.Title.TextStyle = draw.TextStyle{
		Color:  color.Black,
		Font:   font,
		XAlign: draw.XCenter,
		YAlign: draw.YTop,
	}

	plot.X.Label.Text = "Peer i"
	plot.Y.Label.Text = "Peer j"
	plot.X.Label.TextStyle = draw.TextStyle{Color: color.Black, Font: font}
	plot.Y.Label.TextStyle = draw.TextStyle{Color: color.Black, Font: font}

	bubbles := newBubbles(data)
	plot.Add(bubbles)

	err = plot.Save(7 * vg.Inch, 7 * vg.Inch, DIRPNG + label + ".png")
	ex.CheckError(err)
}

func drawTimeData(plot *plot.Plot, data plotter.XYs, mutation string) {
	font, err := vg.MakeFont("Helvetica-Bold", FONTSIZE * 1.2)

	if err != nil {
		println("failed to create font: %v", err)
	}

	plot.Title.TextStyle = draw.TextStyle{
		Color:  color.Black,
		Font:   font,
		XAlign: draw.XCenter,
		YAlign: draw.YTop,
	}

	plot.Title.Text = "Number of nodes deciding over time"
	plot.X.Label.Text = "Time (s)"
	plot.Y.Label.Text = "Number of nodes"

	plot.X.Label.TextStyle = draw.TextStyle{Color: color.Black, Font: font}
	plot.Y.Label.TextStyle = draw.TextStyle{Color: color.Black, Font: font}

	err = plotutil.AddLines(plot, mutation, data)
	err = plot.Save(7 * vg.Inch, 7 * vg.Inch, DIRPNG + "time_" + mutation + ".png")

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
		color := color.RGBA{R: 25, B: 178, A: 255}
		minRadius, maxRadius := vg.Points(0), vg.Points(10)

		rangeRadius := maxRadius - minRadius
		_, _, z := data.XYZ(i)
		diameter := (z - minZ) / (maxZ - minZ)
		radius := vg.Length(diameter)*rangeRadius + minRadius

		return draw.GlyphStyle{Color: color, Radius: radius, Shape: draw.CircleGlyph{}}
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

	maxDuration := durationsList[len(durationsList)-1].Seconds()
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
		mutation + "_" +
		fmt.Sprintf("%f", defaultDelta) + "_" +
		strconv.Itoa(maxTries) + "_" +
		strconv.FormatFloat(percentageMiss, 'f', 2, 64) + ".csv")

	ex.CheckError(err)
	defer file.Close()

	for _, results := range data {
		length := len(results.Delays)

		for _, result := range results.Delays[:(length - 1)] {
			value := ensureValue(result)
			file.WriteString(value + ", ")
		}

		lastResult := results.Delays[length-1]
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

func saveResult(endTime time.Time, startTime time.Time, bandwidthExceeded bool, list map[int] *con.Results) {
	file, err := os.OpenFile(DIRCSV + "global_results.csv", os.O_APPEND | os.O_WRONLY, 0666)
	defer file.Close()

	ex.CheckError(err)

	duration := float64(endTime.Sub(startTime)) / float64(time.Millisecond)

	file.WriteString(strconv.Itoa(numberParticipants) + "," +
		fmt.Sprintf("%f", defaultDelta)               + "," +
		strconv.Itoa(maxTries)                        + "," +
		fmt.Sprintf("%f", percentageMiss)             + "," +
		fmt.Sprintf("%f", percentageFaults)           + "," +
		fmt.Sprintf("%f", probabilityToFail)          + "," +
		fmt.Sprintf("%f", latency)                    + "," +
		strconv.Itoa(bandwidth)                       + "," +
		strconv.FormatBool(bandwidthExceeded)         + "," +
		mutationName                                  + "," +
		fmt.Sprintf("%f", duration)                   + "\n")

	if mutationName != "adapted" {
		saveDelays(list)
	}
}

func saveToCsv(list map[int]*con.Results, duration int) {
	fileBandwidth, err := os.OpenFile(DIRCSV + mutationName + "_bandwidth_usage.csv", os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0666)
	fileRetransmission, err := os.OpenFile(DIRCSV + mutationName + "_retransmission.csv", os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0666)

	defer fileBandwidth.Close()
	defer fileRetransmission.Close()

	ex.CheckError(err)

	for _, result := range list {
		for _, bandwidthUsage := range result.ListOfBandwidthUsage {
			fileBandwidth.WriteString(bandwidthUsage + "\n")
		}
		for _, retransmission := range result.ListOfRetransmission {
			fileRetransmission.WriteString(retransmission + "\n")
		}
	}
}

func saveDelays(list map[int] *con.Results) {
	finalResults := make([]*con.Snapshot, 0)

	for _, result := range list {
		listOfSnapshots := make([]*con.Snapshot, 0)
		index := 0

		for {
			snapshot := result.Metrics.GetOther(index)

			if snapshot != nil {
				listOfSnapshots = append(listOfSnapshots, snapshot)
			} else {
				break
			}

			index++
		}

		finalResults = append(finalResults, listOfSnapshots...)
	}

	orderedSnapshots := make([] *con.Snapshot, 0)

	for (len(finalResults) != 0) {
		latestSnapshot := finalResults[0]
		latestSnapShotIndex := 0

		for id, snapshot := range finalResults {
			if latestSnapshot.Timestamp.After(snapshot.Timestamp) {
				latestSnapshot = snapshot
				latestSnapShotIndex = id
			}
		}

		finalResults = append(finalResults[:latestSnapShotIndex], finalResults[latestSnapShotIndex + 1:]...)
		orderedSnapshots = append(orderedSnapshots, latestSnapshot)
	}

	exportResults(orderedSnapshots, len(list))
}


func exportResults(results []*con.Snapshot, numberOfPeers int) {
	filepath := "deep-learning/results/snapshots_" + strconv.Itoa(numberParticipants) + ".csv"
	var fileSnapshot *os.File

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fileSnapshot, err = os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE, 0666)
		ex.CheckError(err)

		// header
		fileSnapshot.WriteString("PeerID,CoordID,Round,Phase,EstimatePeerID,EstimateValue,Decision,isMajority,")

		for id := 1; id <= numberOfPeers; id++ {
			fileSnapshot.WriteString(fmt.Sprintf("Peer%dVote,",id))
		}

		for id := 1; id <= numberOfPeers; id++ {
			fileSnapshot.WriteString(fmt.Sprintf("isFreshForPeer%d,",id))
		}

		for id := 1; id < numberOfPeers; id++ {
			fileSnapshot.WriteString(fmt.Sprintf("DelayToPeer%d,", id))
		}

		fileSnapshot.WriteString(fmt.Sprintf("DelayToPeer%d\n", numberOfPeers))
	} else {
		fileSnapshot, err = os.OpenFile(filepath, os.O_WRONLY | os.O_APPEND, 0666)
		ex.CheckError(err)
	}

	for _, peerResult := range results {
		// rows
		fileSnapshot.WriteString("" +
			strconv.Itoa(peerResult.PeerID)         + "," +
			strconv.Itoa(peerResult.CoordID)        + "," +
			strconv.Itoa(peerResult.Round)          + "," +
			strconv.Itoa(peerResult.Phase)          + "," +
			strconv.Itoa(peerResult.EstimatePeerID) + "," +
			peerResult.EstimateValue                + "," +
			peerResult.Decision                     + ",")

		if peerResult.IsMajority {
			fileSnapshot.WriteString("1,")
		} else {
			fileSnapshot.WriteString("0,")
		}

		for id := 1; id <= numberOfPeers; id++ {
			vote, present := peerResult.Voters[id]

			if present {
				fileSnapshot.WriteString(strconv.Itoa(vote) + ",")
			} else {
				fileSnapshot.WriteString("0,")
			}
		}

		for id := 0; id < numberOfPeers; id++ {
			isFreshValue := peerResult.IsFresh[id]

			if isFreshValue {
				fileSnapshot.WriteString("1,")
			} else {
				fileSnapshot.WriteString("0,")
			}
		}

		for id := 0; id < numberOfPeers - 1; id++ {
			delay := peerResult.Delays[id]

			fileSnapshot.WriteString(fmt.Sprintf("%f,", delay))
		}

		delay := peerResult.Delays[numberOfPeers - 1]
		fileSnapshot.WriteString(fmt.Sprintf("%f\n", delay))
	}

	fileSnapshot.Close()
}
