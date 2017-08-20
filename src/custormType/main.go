package main

import (
	"unicode"
	"fmt"
	"strings"
	"custormType/Place"
)

type RuneForRuneFunc func(rune) rune

func t1()  {
	var removePunctuation RuneForRuneFunc
	phrases := []string{"Day; dusk, and night.", "All day long"}
	removePunctuation = func(char rune) rune {
		if unicode.Is(unicode.Terminal_Punctuation, char) {
			return -1
		}
		return char
	}
	processPhrases(phrases, removePunctuation)
}

func processPhrases(phrases []string, function RuneForRuneFunc) {
	for _, phrase := range phrases {
		fmt.Println(strings.Map(function, phrase))
	}
}
//=======================================================
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

func t2()  {
	part := Part{5, "wrench"}
	part.UpperCase()
	part.Id += 11
	fmt.Println(part, part.HasPrefix("w"))
	// <<16 "WRENCH">> false
}

//=======================================================
type Item struct {
	id string
	price float64
	quantity int
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item
	catalogId int
}

func t3()  {
	special := SpecialItem{Item{"Green", 3,5}, 207}
	fmt.Println(special.id, special.price, special.quantity, special.catalogId)
	fmt.Println(special.Cost())
	//Green 3 5 207
	//15
}

//方法重载------------------------------------------------------------------

type LuxuryItem struct {
	Item
	markup float64
}

func (item *LuxuryItem) Cost() float64 {
	return item.Item.Cost() * item.markup
}

func t4()  {
	luxu := LuxuryItem{Item{"Green", 3, 5}, 0.2}
	fmt.Println(luxu.Cost())
	// 3
}
//=========================================================

func t5()  {
	newYork := Place.New(40.716667, -74, "New York")
	fmt.Println(newYork)
	baltimore := newYork.Copy()
	baltimore.SetLatitude(newYork.Latitude() - 1.43333)
	baltimore.SetLongtitude(newYork.Longtitude() - 2.61667)
	baltimore.Name = "Baltimore"
	fmt.Println(baltimore)
	//(40.717度, -74.000度 ) "New York"
	//(39.283度, -76.617度 ) "Baltimore"
}


//=======================================================
func main()  {
	t5()
}
