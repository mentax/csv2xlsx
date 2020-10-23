package main

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

type params struct {
	output string
	input  []string

	xlsxTemplate string

	sheets []string
	row    int
}

func initCommandLine(args []string) error {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "csv2xlsx"
	app.Usage = "Convert CSV data to XLSX - especially the big one. \n\n" +
		"Example: \n" +
		"   csv2xlsx --template example/template.xlsx --sheet Sheet_1 --sheet Sheet_2 --row 2 --output result.xlsx data.csv data2.csv \n" +
		"   csv2xlsx.exe -t example\\template.xlsx -s Sheet_1 -s Sheet_2 -r 2 -o result.xlsx data.csv data2.csv "

	app.Version = "0.2.3"
	app.ArgsUsage = "[file of file's list with csv data]"

	app.Flags = []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "sheets",
			Aliases: []string{"s"},
			Usage:   "sheet `names` in the same order like csv files. If sheet with that name exists, data is inserted to this sheet. Usage: -s AA -s BB ",
		},
		&cli.StringFlag{
			Name:    "template",
			Aliases: []string{"t"},
			Value:   "",
			Usage:   "`path` to xlsx file with template output",
		},
		&cli.IntFlag{
			Name:    "row",
			Aliases: []string{"r"},
			Value:   0,
			Usage:   "row `number` to use for create rows format. When '0' - not used. This row will be removed from xlsx file.",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Value:   "./output.xlsx",
			Usage:   "path to result `xlsx file`",
		},
	}

	app.Action = func(c *cli.Context) error {

		p, err := checkAndReturnParams(c)
		if err != nil {
			return err
		}

		return buildXls(p)
	}

	return app.Run(args)
}

func checkAndReturnParams(c *cli.Context) (*params, error) {
	p := &params{}

	output := c.String("output")
	if output == "" {
		return nil, cli.Exit("Path to output file not defined", 1)
	}

	output, err := filepath.Abs(output)
	if err != nil {
		return nil, cli.Exit("Wrong path to output file", 2)
	}
	p.output = output

	//

	p.input = make([]string, c.Args().Len())
	for i, f := range c.Args().Slice() {
		filename, err := filepath.Abs(f)
		if err != nil {
			return nil, cli.Exit("Wrong path to input file "+filename, 3)
		}
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return nil, cli.Exit("Input file does not exist ( "+filename+" )", 4)
		}

		p.input[i] = filename
	}

	//

	p.row = c.Int("row")
	p.sheets = c.StringSlice("sheets")

	//

	xlsxTemplate := c.String("template")
	if xlsxTemplate != "" {
		xlsxTemplate, err = filepath.Abs(xlsxTemplate)
		if err != nil {
			return nil, cli.Exit("Wrong path to template file", 5)
		}
		if _, err := os.Stat(xlsxTemplate); os.IsNotExist(err) {
			return nil, cli.Exit("Template file does not exist ( "+xlsxTemplate+" )", 6)
		}
		p.xlsxTemplate = xlsxTemplate
	}

	if p.row != 0 && xlsxTemplate == "" {
		return nil, cli.Exit("Defined `row template` without xlsx template file", 7)
	}

	return p, nil
}
