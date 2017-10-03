package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func initCommandLine(args []string) error {
	cli.NewApp()

	cli.OsExiter = func(c int) {
		fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
	}

	app := cli.NewApp()
	app.Name = "cvs2xlsx"
	app.Usage = "Convert CSV data to xlsx - especially the big one"
	app.Version = "0.1.0"
	app.ArgsUsage = "[file of file's list with csv data]"

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "sheets, s",
			Usage: "sheet `names` in the same order like cvs files. If sheet with that name exists, data is inserted to this sheet. Usage: -s AA -s BB ",
		},
		cli.StringFlag{
			Name:  "template, t",
			Value: "",
			Usage: "`path` to xlsx file with template output",
		},
		cli.IntFlag{
			Name:  "row, r",
			Value: 0,
			Usage: "row `number` to use for create rows format. When '0' - not used. This row will be removed from xlsx file.",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "./output.xlsx",
			Usage: "path to result `xlsx file`",
		},
	}

	//cli.Command{}

	app.Action = func(c *cli.Context) error {

		p, err := checkAndReturnParams(c)
		if err != nil {
			return err
		}

		fmt.Println(p)

		//fmt.Printf("Args   %#v  \n", c.Args()) //  Get(0)
		//fmt.Printf("sheets  %#v   \n", c.StringSlice("sheets"))

		return nil
		return cli.NewExitError("oh err well", 0)
	}

	return app.Run(args)
}

func checkAndReturnParams(c *cli.Context) (*params, error) {
	p := &params{}

	output := c.String("output")
	if output == "" {
		return nil, cli.NewExitError("Path to output file not defined", 1)
	}

	output, err := filepath.Abs(output)
	if err != nil {
		return nil, cli.NewExitError("Wrong path to output file", 2)
	}
	p.output = output

	//

	p.input = make([]string, len(c.Args()))
	for i, f := range c.Args() {
		filename, err := filepath.Abs(f)
		if err != nil {
			return nil, cli.NewExitError("Wrong path to input file "+filename, 3)
		}
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return nil, cli.NewExitError("Input file does not exist ( "+filename+" )", 4)
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
			return nil, cli.NewExitError("Wrong path to template file", 5)
		}
		if _, err := os.Stat(xlsxTemplate); os.IsNotExist(err) {
			return nil, cli.NewExitError("Template file does not exist ( "+xlsxTemplate+" )", 6)
		}
		p.xlsxTemplate = xlsxTemplate
	}

	return p, nil
}

type params struct {
	output string
	input  []string

	xlsxTemplate string

	sheets []string
	row    int
}
