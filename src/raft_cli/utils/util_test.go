package utils

import (
	"fmt"
	"reflect"
)

func getStringSlice() []string {
	return []string{
		"mhc",
		"mhc222",
		"hello",
		"mhc_hello",
	}
}

func ExampleSliceFilter_String() {
	ss := getStringSlice()
	fmt.Println(SliceFilter(ChangeSliceType(ss), ""))
	//Output: [mhc mhc222 hello mhc_hello]
}

func ExampleSliceFilter_String2() {
	ss := getStringSlice()
	fmt.Println(SliceFilter(ChangeSliceType(ss), "mhc"))
	//Output: [mhc mhc222 mhc_hello]
}

func ExampleSliceFilter_String3() {
	ss := getStringSlice()
	fmt.Println(SliceFilter(ChangeSliceType(ss), "mhc&hello"))
	//Output: [mhc_hello]
}

func ExampleSliceFilter_String4() {
	ss := getStringSlice()
	fmt.Println(SliceFilter(ChangeSliceType(ss), "mhc", func(s reflect.Value) bool {
		return Contains(s.String(), "222")
	}))
	//Output: [mhc mhc_hello]
}

type testStructure struct {
	key1 string
	key2 string
}

func getStructureSlice() (res []testStructure) {
	res = append(res, testStructure{key1:"mhc", key2:"111"})
	res = append(res, testStructure{key1:"mhc222",key2:"222"})
	res = append(res, testStructure{key1:"hello", key2:"333"})
	res = append(res, testStructure{key1:"mhc_hello",key2:"444"})
	return
}

func ExampleSliceFilter_Structure() {
	ShowInTable(SliceFilter(ChangeSliceType(getStructureSlice()), ""), "key1", "key2")
	//Output:
	//+-----------+------+
	//|   KEY1    | KEY2 |
	//+-----------+------+
	//| mhc       |  111 |
	//| mhc222    |  222 |
	//| hello     |  333 |
	//| mhc_hello |  444 |
	//+-----------+------+
}

func ExampleSliceFilter_Structure2() {
	ShowInTable(SliceFilter(ChangeSliceType(getStructureSlice()), "key1&mhc"), "key1", "key2")
	//Output:
	//+-----------+------+
	//|   KEY1    | KEY2 |
	//+-----------+------+
	//| mhc       |  111 |
	//| mhc222    |  222 |
	//| mhc_hello |  444 |
	//+-----------+------+
}

func ExampleSliceFilter_Structure3() {
	ShowInTable(SliceFilter(ChangeSliceType(getStructureSlice()), "key1&mhc&hello"), "key1", "key2")
	//Output:
	//+-----------+------+
	//|   KEY1    | KEY2 |
	//+-----------+------+
	//| mhc_hello |  444 |
	//+-----------+------+
}

func testCondition(s reflect.Value) bool {
	if s.FieldByName("key2").String() == "222" {
		return true
	}
	return false
}

func ExampleSliceFilter_Structure4() {
	ShowInTable(SliceFilter(ChangeSliceType(getStructureSlice()), "key1&mhc", testCondition), "key1", "key2")
	//Output:
	//+-----------+------+
	//|   KEY1    | KEY2 |
	//+-----------+------+
	//| mhc       |  111 |
	//| mhc_hello |  444 |
	//+-----------+------+
}