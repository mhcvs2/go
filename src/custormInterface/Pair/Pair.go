package Pair

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Exchanger interface {
	Exchange()
}

func ExchangeThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

//----------------------------------------

type LowerCaser interface {
	LowerCase()
}

type UpperCaser interface {
	UpperCase()
}

type LowerUpperCaser interface {
	LowerCaser
	UpperCaser
}

type FixCaser interface {
	FixCase()
}

type ChangeCaser interface {
	LowerUpperCaser
	FixCaser
}

//------------------------------

type StringPair struct {
	First, Second string
}

func (pair *StringPair) Exchange() {
	pair.First, pair.Second = pair.Second, pair.First
}

func (pair StringPair) String() string {
	return fmt.Sprintf("%q+%q", pair.First, pair.Second)
}

func (pair *StringPair) Read(data []byte) (n int, err error) {
	if pair.First == "" && pair.Second == "" {
		return 0, io.EOF
	}
	if pair.First != "" {
		n = copy(data, pair.First)
		pair.First = pair.First[n:]
	}
	if n < len(data) && pair.Second != "" {
		m := copy(data[n:], pair.Second)
		pair.Second = pair.Second[m:]
		n += m
	}
	return n, nil
}

func (pair *StringPair) UpperCase()  {
	pair.First = strings.ToUpper(pair.First)
	pair.Second = strings.ToUpper(pair.Second)
}

func (pair *StringPair) LowerCase()  {
	pair.First = strings.ToLower(pair.First)
	pair.Second = strings.ToLower(pair.Second)
}

func (pair *StringPair) FixCase() {
	pair.First = FixCase(pair.First)
	pair.Second = FixCase(pair.Second)
}

//----------------------------

type Point [2]int

func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}

//-------------------------------

func FixCase(s string) string {
	var chars []rune
	upper := true
	for _, char := range s {
		if upper {
			char = unicode.ToUpper(char)
		} else {
			char = unicode.ToLower(char)
		}
		chars = append(chars, char)
		upper = unicode.IsSpace(char) || unicode.Is(unicode.Hyphen, char)
	}
	return string(chars)
}

