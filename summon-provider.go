package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gopasspw/gopass/pkg/ctxutil"
	"github.com/gopasspw/gopass/pkg/gopass"
	"github.com/urfave/cli/v3"
)

// Stdout is exported for tests.
var Stdout io.Writer = os.Stdout

type gc struct {
	gp gopass.Store
}

// Get outputs the password for given path on stdout.
func (s *gc) Get(ctx context.Context, cmd *cli.Command) error {
	ctx = ctxutil.WithNoNetwork(ctx, true)
	path := cmd.Args().Get(0)
	secret, err := s.gp.Get(ctx, path, "latest")
	if err != nil {
		return err
	}

	fmt.Fprintln(Stdout, secret.Password())

	return nil
}
