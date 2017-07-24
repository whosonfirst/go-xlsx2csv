# go-xlsx2csv

Go tools for converting XLSX files to CSV

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Caveats

This is basically the same as @tealeg 's [xlsx2csv](https://github.com/tealeg/xlsx2csv/) program but uses the default `encoding/csv` package to generate CSV data and spits everything to STDOUT.

## See also

* github.com/tealeg/xlsx

