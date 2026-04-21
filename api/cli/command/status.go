package command

import (
	"github.com/effective-security/promptviser/api/cli"
	"github.com/effective-security/promptviser/api/client"
)

// VersionCmd prints remote server version
type VersionCmd struct {
}

// Run the command
func (a *VersionCmd) Run(ctx *cli.Cli) error {
	c, err := ctx.HTTPClient(true)
	if err != nil {
		return err
	}

	res, err := client.NewHTTPStatusClient(c).Version(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.Print(res)
}

// ServerCmd prints remote server status
type ServerCmd struct {
}

// Run the command
func (a *ServerCmd) Run(ctx *cli.Cli) error {
	c, err := ctx.HTTPClient(true)
	if err != nil {
		return err
	}

	res, err := client.NewHTTPStatusClient(c).Status(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.Print(res)
}

/*
// CallerCmd shows the caller status
type CallerCmd struct {
}

// Run the command
func (a *CallerCmd) Run(ctx *cli.Cli) error {
	c, err := ctx.AuthClient(false)
	if err != nil {
		return err
	}

	res, err := c.Caller(ctx.Context(), &emptypb.Empty{})
	if err != nil {
		return err
	}

	return ctx.Print(res)
}
*/
