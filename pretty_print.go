package main

import (
	"io/fs"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	// "github.com/jedib0t/go-pretty/v6/text"
)

func PrintFiles(files []fs.FileInfo){
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.SetTitle("Todos")
	t.AppendHeader(table.Row{"NAME", "SIZE", "MODE","LAST-MODIFIED"})

	for _ , file := range files {
		modTime := file.ModTime().Format(time.RFC1123) 
		t.AppendRow(table.Row{file.Name(),file.Size(),file.Mode().String(),modTime})
	}


}