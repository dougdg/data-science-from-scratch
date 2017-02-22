package main

import (
	"html/template"
	"os"

	"github.com/go-template/nvd3-go-template/browser"
)

var lineChartTemplate *template.Template
var barChartTemplate *template.Template

type Line_type struct {
	X	int
	Y	float64
}

type Bar_type struct {
	Label string
	Value float64
}

var data []Line_type
var data_bar []Bar_type

func init() {
	lineChartTemplate = template.Must(template.ParseFiles("templates/linechart.html",
		"templates/linechartjs.tmpl"))
	barChartTemplate = template.Must(template.ParseFiles("templates/discreteBarChart.html",
		"templates/discreteBarChart.tmpl"))


	data = []Line_type {
		Line_type {X: 1950, Y: 300.2},
		Line_type {X: 1950, Y: 300.2},
		Line_type {X: 1960, Y: 543.3},
		Line_type {X: 1970, Y: 1075.9},
		Line_type {X: 1980, Y: 2862.5},
		Line_type {X: 1990, Y: 5979.6},
		Line_type {X: 2000, Y: 10289.7},
		Line_type {X: 2010, Y: 14958.3},
	}

	data_bar = []Bar_type {
		Bar_type {Label: "Annie Hall", Value: 5},
		Bar_type {Label: "Ben-Hur", Value: 11},
		Bar_type {Label: "Casablanca", Value: 3},
		Bar_type {Label: "Gandhi", Value: 8},
		Bar_type {Label: "West Side Story", Value: 10},
	}
}

func renderLineChart(path string) {
	var file = createFile(path)
	lineChartTemplate.ExecuteTemplate(file, "linechart.html", data)
	defer file.Close()
}

func renderBarChart(path string) {
	var file = createFile(path)
	barChartTemplate.ExecuteTemplate(file, "discreteBarChart.html",
	 data_bar)
	defer file.Close()
}

func createFile(path string) (f *os.File) {
	f, err := os.Create(path)
	if err != nil {
    panic(err)
	}
	return f
}

func deleteDirOutput() {
		os.Remove("output")
}

func createDirOutput(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
    os.Mkdir(path, os.ModePerm)
	}
}

func main() {
	deleteDirOutput()
	createDirOutput("output")

	chartType := os.Args[1]
	path := "output/" + chartType

	if (chartType == "linechart") {
			renderLineChart(path)
	} else {
			renderBarChart(path)
	}

	browser.Open(path)
}
