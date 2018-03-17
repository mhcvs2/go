package main

import "fmt"

type SchoolReport interface {
	Report()
	Sign()
}

type FourSchoolReport struct {}

func (sr *FourSchoolReport) Report() {
	fmt.Println("report score")
}

func (sr *FourSchoolReport) Sign() {
	fmt.Println("sign name")
}

///////////////////////////////////////////////

type Decorator struct {
	sr SchoolReport
}

func NewDecorator(sr SchoolReport) *Decorator {
	return &Decorator{sr: sr}
}

func (d *Decorator) Report() {
	d.sr.Report()
}

func (d *Decorator) Sign() {
	d.sr.Sign()
}

////////////////////////////////////////////////

type HighScoreDecorator struct {
	Decorator
}

func NewHighScoreDecorator(sr SchoolReport) *HighScoreDecorator {
	return &HighScoreDecorator{Decorator: *NewDecorator(sr)}
}

func (h *HighScoreDecorator) reportHignScore() {
	fmt.Println("report high score")
}

func (h *HighScoreDecorator) Report() {
	h.reportHignScore()
	h.Decorator.Report()
}


////////////////////////////////////////////////

type SortDecorator struct {
	Decorator
}

func NewSortDecorator(sr SchoolReport) *SortDecorator {
	return &SortDecorator{Decorator: *NewDecorator(sr)}
}

func (h *SortDecorator) reportSort() {
	fmt.Println("report sort")
}

func (h *SortDecorator) Report() {
	h.Decorator.Report()
	h.reportSort()
}
///////////////////////////////////////////////

func main() {
	var sr SchoolReport
	sr = new(FourSchoolReport)
	sr = NewHighScoreDecorator(sr)
	sr = NewSortDecorator(sr)

	sr.Report()
	sr.Sign()
}

//report high score
//report score
//report sort
//sign name