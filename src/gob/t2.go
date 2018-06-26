package main

import (
	"os"
	"encoding/gob"
	"runtime"
	"fmt"
)

const file = "./test.gob"

type User struct {
	Name, Pass string
}

func t1() {
	var datato = &User{"Donald1", "DuckPass1"}
	var datafrom = new(User)
	err := Save(file, datato)
	Check(err)
	err = Load(file, datafrom)
	Check(err)
	fmt.Println(datafrom)
}

func t2() {
	var datato = &User{"Donal22d", "22"}
	var datafrom = new(User)
	err := Save2(file, datato)
	Check(err)
	err = Load(file, datafrom)
	Check(err)
	fmt.Println(datafrom)
}

func main() {
	t1()
}

func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func Save2(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	} else {
		panic(err)
	}
	file.Close()
	return err
}

func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}