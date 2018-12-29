// Fairly wasteful in memory use. Don't bother with large data sets!

package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
  "strings"
)

// Column contains a set of items and a filename that they are read from
type Column struct {
  Filename string
  Items    map[string]bool
}

// Has returns true if an item is present in a column data set
func (c *Column) Has(item string) bool {
  _,found := c.Items[item]
  return found
}

// NewColumn reads lines from a file and returns them as a Column object
func NewColumn(filename string) (*Column, error) {
  file, err := os.Open(filename)
  if err != nil {
    return nil, err
  }
  defer file.Close()
  col := &Column{
    Filename: filename,
    Items: make(map[string]bool),
  }
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    col.Items[scanner.Text()] = true
  }
  return col, scanner.Err()
}

// MatchTable contains a set of Column data sets
type MatchTable struct {
  Columns  []*Column
  Superset map[string]bool
}

// NewMatchTable creates a new MatchTable object and reads any
// supplied files in as column data sets
func NewMatchTable(filenames []string) (*MatchTable, error){
  mt := &MatchTable{Columns: make([]*Column,0)}
  for _,filename := range filenames {
    err := mt.AddColumn(filename)
    if err != nil {
      return nil,err
    }
  }
  mt.UpdateSuperset()
  return mt, nil
}

// AddColumn adds a new data set to a MatchTable
func (mt *MatchTable) AddColumn(filename string) error {
  c,err := NewColumn(filename)
  if err != nil {
    return err
  }
  mt.Columns = append(mt.Columns, c)
  return nil
}

// UpdateSuperset refreshes the superset of items
func (mt *MatchTable) UpdateSuperset() {
  newsuperset := make(map[string]bool)
  for _,column := range mt.Columns {
    for item := range column.Items {
      newsuperset[item] = true
    }
  }
  mt.Superset = newsuperset
}

// GenerateTable renders a MatchTable into a 2D map
func (mt *MatchTable) GenerateTable(renderopts *TableRenderOptions) [][]string {
  var result [][]string
  var header []string
  header = append(header,"ITEM")
  for _,column := range mt.Columns {
    header = append(header, column.Filename)
  }
  result = append(result,header)
  for item := range mt.Superset {
    var row []string
    // leftmost column is the item we're looking at
    row = append(row, item)
    for _,column := range mt.Columns {
      if column.Has(item) {
        row = append(row,renderopts.YesValue)
      } else {
        row = append(row,renderopts.NoValue)
      }
    }
    result = append(result,row)
  }
  return result
}

// TableRenderOptions encapsulates configuration for rendering tables.
type TableRenderOptions struct {
  YesValue, NoValue string
}

func main() {
  yesValue := flag.String("yes-value", "X", "string used to indicate an item was present in a column")
  noValue := flag.String("no-value", "-", "string used to indicate an item was NOT present in a column")
  separator := flag.String("separator", " ", "string used to separate rendered columns in the output")
  flag.Parse()
  renderopts := &TableRenderOptions{YesValue: *yesValue, NoValue: *noValue}
  mt,err := NewMatchTable(flag.Args())
  if err != nil {
    fmt.Printf("error setting up matchtable: %v\n", err)
    os.Exit(1)
  }
  for _,row := range mt.GenerateTable(renderopts) {
    fmt.Println(strings.Join(row,*separator))
  }
}
