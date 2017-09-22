package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func t1() {
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Sign", "Rating"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

/*
+------+-----------------------+--------+
| NAME |         SIGN          | RATING |
+------+-----------------------+--------+
| A    | The Good              |    500 |
| B    | The Very very Bad Man |    288 |
| C    | The Ugly              |    120 |
| D    | The Gopher            |    800 |
+------+-----------------------+--------+
*/

func t2() {
	data := [][]string{
		[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
		[]string{"1/1/2014", "January Hosting", "2345", "$54.95"},
		[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
		[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

func t3() {
	data := [][]string{
		[]string{"id", "'testredis'", ""},
		[]string{"name", "'testrediscluster'", ""},
		[]string{"endpoint", "'redis-testredis.svc.41d254d6-6b12-43a6-d8b8-d79b0662346b.bst-1.cns.bstjpc.com'", ""},
		[]string{"port", "'6379'", ""},
		[]string{"state", "'avaiable'", ""},
		[]string{"images", "dbproxy", "'a7b8d9a51803'"},
		[]string{"images", "redis", "'5962c9dddbe0'"},
		[]string{"runnings", "dbproxy", "'1/1:[6e29bccdd5cc]'"},
		[]string{"runnings", "mgmtapi", "'1/1:[54ca4fb07d91]'"},
		[]string{"runnings", "redis", "'3/3:[2ccbf521bc99,794c61b820dd,e81d33e43533]'"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	//table.SetFooter([]string{"", "", "Total", "$146.93"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}

func main() {
	t3()
}