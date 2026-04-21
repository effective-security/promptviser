package command

import (
	"github.com/effective-security/promptviser/api/cli"
	"github.com/effective-security/promptviser/api/pb"
)

type SubmitCmd struct {
	// TODO
	Data string `help:"Data to submit"`
}

func (cmd *SubmitCmd) Run(c *cli.Cli) error {
	adviser, err := c.AdviserClient(true)
	if err != nil {
		return err
	}

	data, err := c.Resolve(cmd.Data)
	if err != nil {
		return err
	}

	res, err := adviser.Submit(c.Context(), &pb.SubmitRequest{
		Data: data,
	})
	if err != nil {
		return err
	}

	return c.Print(res)
}
