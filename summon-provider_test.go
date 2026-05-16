package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/gopasspw/gopass/pkg/gopass/apimock"
	"github.com/gopasspw/gopass/pkg/termio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func TestSummonProviderOutputsOnlySecret(t *testing.T) { //nolint:paralleltest
	ctx := t.Context()
	act := &gc{
		gp: apimock.New(),
	}
	require.NoError(t, act.gp.Set(ctx, "foo", &apimock.Secret{Buf: []byte("bar\nbaz: zab")}))

	buf := &bytes.Buffer{}
	Stdout = buf
	color.NoColor = true
	defer func() {
		termio.Stdin = os.Stdin
		Stdout = os.Stdout
	}()

	app := &cli.Command{
		Action: act.Get,
	}
	require.NoError(t, app.Run(ctx, []string{"app", "foo"}))
	assert.Equal(t, "bar\n", buf.String())
}
