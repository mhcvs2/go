package main

import "fmt"

func t1() {
	points := []struct{x,y int} {{4,6}, {}, {-7, 11}, {15, 17}, {14, -8}}
	for _, point := range points {
		fmt.Printf("(%d, %d)", point.x, point.y)
	}
}

//=====================================================================

type optioner interface {
	Name() string
	IsValid() bool
}

type OptionCommon struct {
	ShortName string "short option name"
	LongName string "long option name"
}

//-------------------------------

type IntOption struct {
	OptionCommon
	Value, Min, Max int
}

func (option IntOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option IntOption) IsValid() bool {
	return option.Value <= option.Max && option.Value >= option.Min
}

//---------------------------------

type FloatOption struct {
	optioner
	Value float64
}

type GenericOption struct {
	OptionCommon
}

func (option GenericOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option GenericOption) IsValid() bool {
	return true
}

//-------------------------------------
func name(shortName, longName string) string {
	if longName == "" {
		return shortName
	}
	return longName
}
//---------------------------------------
func t2() {
	topOption := IntOption{OptionCommon:OptionCommon{"t", "top"}, Max:100}
	sizeOption := FloatOption{GenericOption{OptionCommon{"s", "size"}}, 19.5}

	for _, option := range []optioner{topOption, sizeOption} {
		fmt.Print("Name=", option.Name(), "-valid=", option.IsValid())
		fmt.Print("-value=")
		switch option := option.(type) {
		case IntOption:
			fmt.Print(option.Value, "-min=",option.Min, "-max=",option.Max,"\n")
		case FloatOption:
			fmt.Println(option.Value)
		}
	}
	//Name=top-valid=true-value=0-min=0-max=100
	//Name=size-valid=true-value=19.5
}

func main() {
	t2()
}
