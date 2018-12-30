// Fairly wasteful in memory use. Don't bother with large data sets!

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

// MatchTable contains a set of Column data sets
type MatchTable struct {
	Columns  []map[string]bool
	Headings []string
}

// NewMatchTable creates a new MatchTable object and reads any
// supplied files in as column data sets
func NewMatchTable(filenames []string) (*MatchTable, error) {
	mt := &MatchTable{Columns: []map[string]bool{}, Headings: []string{}}
	for _, filename := range filenames {
		err := mt.AddColumn(filename)
		if err != nil {
			return nil, err
		}
	}
	return mt, nil
}

// AddColumn adds a new data set to a MatchTable
func (mt *MatchTable) AddColumn(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	col := make(map[string]bool)
	for scanner.Scan() {
		col[scanner.Text()] = true
	}
	mt.Columns = append(mt.Columns, col)
	mt.Headings = append(mt.Headings, filename)
	return nil
}

// Superset refreshes the superset of items
func (mt *MatchTable) Superset() []string {
	track := make(map[string]bool)
	var superset []string
	for _, column := range mt.Columns {
		for item := range column {
			if _, found := track[item]; found {
				continue
			}
			superset = append(superset, item)
			track[item] = true
		}
	}
	return superset
}

// GenerateTable renders a MatchTable into a 2D map
func (mt *MatchTable) GenerateTable(renderopts *TableRenderOptions) [][]string {
	var result [][]string
	var header []string
	header = append(header, "ITEM")
	for _, heading := range mt.Headings {
		header = append(header, heading)
	}
	result = append(result, header)
	superset := mt.Superset()
	if renderopts.Sort {
		sort.Strings(superset)
	}
	for _, item := range superset {
		var row []string
		// leftmost column is the item we're looking at
		row = append(row, item)
		for _, column := range mt.Columns {
			if _, present := column[item]; present {
				row = append(row, renderopts.YesValue)
			} else {
				row = append(row, renderopts.NoValue)
			}
		}
		result = append(result, row)
	}
	return result
}

// TableRenderOptions encapsulates configuration for rendering tables.
type TableRenderOptions struct {
	YesValue, NoValue string
	Sort              bool
}

func main() {
	yesValue := flag.String("yes-value", "X", "string used to indicate an item was present in a column")
	noValue := flag.String("no-value", "-", "string used to indicate an item was NOT present in a column")
	separator := flag.String("separator", " ", "string used to separate rendered columns in the output")
	sort := flag.Bool("sort", true, "sort the superset of items lexicographically")
	flag.Parse()
	renderopts := &TableRenderOptions{YesValue: *yesValue, NoValue: *noValue, Sort: *sort}
	mt, err := NewMatchTable(flag.Args())
	if err != nil {
		fmt.Printf("error setting up matchtable: %v\n", err)
		os.Exit(1)
	}
	for _, row := range mt.GenerateTable(renderopts) {
		fmt.Println(strings.Join(row, *separator))
	}
}
