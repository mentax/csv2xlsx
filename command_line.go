package main

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

type params struct {
	output string
	input  []string

	xlsxTemplate string

	sheets     []string
	exampleRow int

	startFrom int

	delimiter rune
}

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func initCommandLine(args []string) error {
	app := cli.NewApp()
	app.Suggest = true
	app.EnableBashCompletion = true
	app.Name = "csv2xlsx"
	app.Usage = "Convert CSV data to XLSX - especially the big one. \n\n" +
		"Example: \n" +
		"   csv2xlsx --template example/template.xlsx --sheet Sheet_1 --sheet Sheet_2 --exampleRow 2 --output result.xlsx data.csv data2.csv \n" +
		"   csv2xlsx.exe -t example\\template.xlsx -s Sheet_1 -s Sheet_2 -r 2 -o result.xlsx data.csv data2.csv "

	app.Version = version + " built in " + date + " from commit: [" + commit + "] by " + builtBy
	app.ArgsUsage = "[file or file's list with csv data]"

	app.DisableSliceFlagSeparator = true

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
			Usage:   "`path` to xlsx file with template file",
		},
		&cli.StringFlag{
			Name:    "delimiter",
			Aliases: []string{"d"},
			Value:   ",",
			Usage:   "one `letter` delimiter used in csv file",
		},
		&cli.IntFlag{
			Name:    "exampleRow",
			Aliases: []string{"r"},
			Value:   0,
			Usage:   "exampleRow `number` to use for create rows format. When '0' - not used. This exampleRow will be overwrite in result file.",
		},
		&cli.IntFlag{
			Name:    "startFrom",
			Aliases: []string{"sf"},
			Value:   0,
			Usage:   "startFrom `number` decide which row is used as first row from csv file. Counting from 0.",
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

		if err = buildXls(p); err != nil {
			return cli.Exit(err.Error(), 99)
		}
		return nil
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

	delimiter := []rune(c.String("delimiter"))
	p.delimiter = delimiter[0]
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

	if len(p.input) < 1 {
		return nil, cli.Exit("Missing path to input file", 5)
	}
	//

	p.startFrom = c.Int("startFrom")

	p.exampleRow = c.Int("exampleRow")
	p.sheets = c.StringSlice("sheets")

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

	if p.exampleRow != 0 && xlsxTemplate == "" {
		return nil, cli.Exit("Defined `exampleRow in template` without template file", 7)
	}

	return p, nil
}
