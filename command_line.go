package main

import (
	"fmt"
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

		fmt.Printf("Args   %#v  \n", c.Args()) //  Get(0)
		fmt.Printf("sheets  %#v   \n", c.StringSlice("sheets"))

		return cli.NewExitError("oh err well", 0)
	}

	return app.Run(args)
}
