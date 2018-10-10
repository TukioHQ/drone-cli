package build

import (
	"strconv"

	"github.com/drone/drone-cli/drone/internal"
	"github.com/urfave/cli"
)

var buildPromoteCmd = cli.Command{
	Name:      "promote",
	Usage:     "promote a build",
	ArgsUsage: "<repo/name> <build> <environment>",
	Action:    buildPromote,
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:  "param, p",
			Usage: "custom parameters to be injected into the job environment. Format: KEY=value",
		},
	},
}

func buildPromote(c *cli.Context) (err error) {
	repo := c.Args().First()
	owner, name, err := internal.ParseRepo(repo)
	if err != nil {
		return err
	}
	number, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return err
	}
	target := c.Args().Get(2)
	params := internal.ParseKeyPair(c.StringSlice("param"))

	client, err := internal.NewClient(c)
	if err != nil {
		return err
	}

	_, err = client.Promote(owner, name, number, target, params)
	if err != nil {
		return err
	}

	return nil
}
