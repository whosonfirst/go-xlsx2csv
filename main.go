package main

import (
	"encoding/csv"
	"flag"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

func main() {

	flag.Parse()

	w := csv.NewWriter(os.Stdout)

	for _, path := range flag.Args() {

		doc, err := xlsx.OpenFile(path)

		if err != nil {
			log.Fatal(err)
		}

		for _, sheet := range doc.Sheets {

			for _, row := range sheet.Rows {

				out := make([]string, 0)

				for _, cell := range row.Cells {
					text := cell.String()
					out = append(out, text)
				}

				err = w.Write(out)

				if err != nil {
					log.Fatal(err)
				}
			}
		}

	}

	w.Flush()
}
