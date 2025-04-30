package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/fatih/color"
	"github.com/gopasspw/gopass/pkg/gopass/apimock"
	"github.com/gopasspw/gopass/pkg/termio"
	"github.com/gopasspw/gopass/tests/gptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	require.NoError(t, act.Get(gptest.CliCtx(ctx, t, "foo")))
	assert.Equal(t, "bar\nbaz: zab\n", buf.String())
}
