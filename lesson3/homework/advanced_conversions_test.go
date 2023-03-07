package main

import (
	"context"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAdvancedConversions(t *testing.T) {
	binPath := composeBinaryPath()
	cmd := exec.Command("go", "build", "-o", binPath, "./")
	assert.NoError(t, cmd.Run())
	defer func() {
		assert.NoError(t, os.Remove(binPath))
	}()

	t.Run("ok with stdin input and stdout result, conv and block-size options", func(t *testing.T) {
		cmd = exec.Command(binPath, "-conv", "trim_spaces,upper_case", "-block-size", "1")
		cmd.Stdin = strings.NewReader(testInput)
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, strings.ToUpper(strings.TrimSpace(testInput)), stdout.String())
	})

	t.Run("ok with unlimited stdin with unicode spaces", func(t *testing.T) {
		limit := 9999
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		cmd = exec.CommandContext(ctx, binPath, "-limit", strconv.Itoa(limit), "-block-size", "1024", "-conv", "trim_spaces")

		cmd.Stdin = &unlimitedReader{input: []byte(" ")}
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, ctx.Err(), "process timed out")
		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, "", stdout.String())
	})

	t.Run("ok with unlimited stdin with unicode spaces and with ansi prefix", func(t *testing.T) {
		limit := 10000
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		cmd = exec.CommandContext(ctx, binPath, "-limit", strconv.Itoa(limit), "-block-size", "1024", "-conv", "trim_spaces")

		cmd.Stdin = &unlimitedReader{input: []byte(" "), prefix: []byte(" ")}
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, ctx.Err(), "process timed out")
		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, "", stdout.String())
	})
}
