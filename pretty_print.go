package main

import (
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrintFiles(files []fs.FileInfo,tableTitle string){
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.SetTitle(tableTitle)
	t.AppendHeader(table.Row{"NAME", "SIZE", "MODE","LAST-MODIFIED"})

	for _ , file := range files {
		modTime := file.ModTime().Format(time.RFC1123)
		fileSize := fmt.Sprintf("%d bytes",file.Size())
		t.AppendRow(table.Row{file.Name(),fileSize,file.Mode().String(),modTime})
	}

	// Applying dark style
	style := t.Style()
	style.Color.Header = text.Colors{text.BgHiBlack, text.FgWhite}
	style.Color.Row = text.Colors{text.BgBlack, text.FgHiWhite}
	t.SetStyle(*style)

	t.Render()

}