package main

import (
	"custormInterface/Pair"
	"fmt"
	"io"
	"strings"
)

func t1() {
	jekyll := Pair.StringPair{"Henry", "Jekyll"}
	hyde := Pair.StringPair{"Edward", "Hyde"}
	point := Pair.Point{5, -3}
	fmt.Println("befeore:", jekyll, hyde, point)

	jekyll.Exchange()
	hyde.Exchange()
	point.Exchange()
	fmt.Println("after #1:", jekyll, hyde, point)

	Pair.ExchangeThese(&jekyll, &hyde, &point)
	fmt.Println("after #2:", jekyll, hyde, point)

	//befeore: "Henry"+"Jekyll" "Edward"+"Hyde" [5 -3]
	//after #1: "Jekyll"+"Henry" "Hyde"+"Edward" [-3 5]
	//after #2: "Henry"+"Jekyll" "Edward"+"Hyde" [5 -3]
}

func t2() {
	const size = 100
	robert := &Pair.StringPair{"Robert L.", "Stevensonsssssssssssssss"}
	david := Pair.StringPair{"David", "Balfour"}
	for _, reader := range []io.Reader{robert, &david} {
		raw, err := ToBytes(reader, size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%q\n", raw)
	}
	//"Robert L.Stevensonsssssssssssssss"
	//"DavidBalfour"
}

func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	return data[:n], nil
}

//=========================================================================
type IsValider interface {
	IsValid() bool
}

type test struct {
	text string
}

func (t *test) IsValid() bool {
	fmt.Println("i am test-IsValid")
	return true
}

func t3() {
	a := test{"haha"}
	//for _, x := range []interface{}{&a} {
	//	if x, ok := x.(IsValider); ok {
	//		x.IsValid()
	//	}
	//}
	var d interface{} = &a
	if thing, ok := d.(IsValider); ok {
		fmt.Println(thing.IsValid(), "isvalider")
	} else {
		fmt.Println("not")
	}
	//i am test-IsValid
	//true isvalider
}

//==========================================================
type Part struct {
	Id int
	Name string
}

func (part *Part)LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part)UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string {
	return fmt.Sprintf("<<%d %q>>", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

func (part *Part) FixCase() {
	part.Name = Pair.FixCase(part.Name)
}


func t4() {
	toaskRack := Part{8427, "TOAST RACK"}
	toaskRack.LowerCase()
	lobelia := Pair.StringPair{"LOBELIA", "SACKVILLE-BAGGINS"}
	lobelia.FixCase()
	fmt.Println(toaskRack, lobelia)
	// <<8427 "toast rack">> "Lobelia"+"Sackville-Baggins"

	for _, x := range []interface{}{&toaskRack, &lobelia} {
		if x, ok := x.(Pair.LowerCaser); ok {
			x.LowerCase()
		}
	}
	fmt.Println(toaskRack, lobelia)
	//<<8427 "toast rack">> "lobelia"+"sackville-baggins"
}


func main() {
	t3()
}
