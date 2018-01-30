package main

import (
	memdb "github.com/hashicorp/go-memdb"
	"fmt"
)

type Person struct {
	Email string
	Name  string
	Age   int
}

func t11() {
	schema := &memdb.DBSchema{
		Tables:map[string]*memdb.TableSchema{
			"person": &memdb.TableSchema{
				Name:"person",
				Indexes:map[string]*memdb.IndexSchema{
					"id":&memdb.IndexSchema{
						Name: "id",
						Unique: true,
						Indexer: &memdb.StringFieldIndex{Field:"Email"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	txn := db.Txn(true)
	p := Person{"joe@aol.com", "Joe", 30}
	if err := txn.Insert("person", p); err != nil {
		panic(err)
	}
	txn.Commit()

	txn = db.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("person", "id", "joe@aol.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s\n", raw.(Person).Name)
	fmt.Printf("Age: %d\n", raw.(Person).Age)
}
//Hello Joe
//Age: 30


func main() {
	t11()
}
