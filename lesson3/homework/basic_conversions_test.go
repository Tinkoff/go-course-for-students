package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicConversions(t *testing.T) {
	binPath := composeBinaryPath()
	cmd := exec.Command("go", "build", "-o", binPath, "./")
	assert.NoError(t, cmd.Run())
	defer func() {
		assert.NoError(t, os.Remove(binPath))
	}()

	t.Run("ok with stdin input and stdout result, offset with trim spaces", func(t *testing.T) {
		cmd = exec.Command(binPath, "-offset", "4", "-conv", "trim_spaces")
		cmd.Stdin = strings.NewReader("HEAD  BA  ")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, "BA", stdout.String())
	})

	t.Run("ok with stdin input and stdout result, limit with trim spaces", func(t *testing.T) {
		cmd = exec.Command(binPath, "-limit", "6", "-conv", "trim_spaces")
		cmd.Stdin = strings.NewReader("  BA  TAIL")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, "BA", stdout.String())
	})

	t.Run("ok with stdin input and stdout result, two conversions", func(t *testing.T) {
		cmd = exec.Command(binPath, "-conv", "trim_spaces,lower_case")
		cmd.Stdin = strings.NewReader("  b ")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, "b", stdout.String())
	})

	t.Run("ok, lower_case", func(t *testing.T) {
		cmd = exec.Command(binPath, "-conv", "lower_case")
		cmd.Stdin = strings.NewReader("WШ")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, "wш", stdout.String())
	})

	t.Run("ok, upper_case", func(t *testing.T) {
		cmd = exec.Command(binPath, "-conv", "upper_case")
		cmd.Stdin = strings.NewReader("wш")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, "WШ", stdout.String())
	})

	t.Run("error with invalid conv", func(t *testing.T) {
		cmd = exec.Command(binPath, "-conv", "qweqwe")
		cmd.Stdin = strings.NewReader(testInput)
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
		assert.Zero(t, stdout.Len())
	})

	t.Run("error due to contradictory conv", func(t *testing.T) {
		cmd = exec.Command(binPath, "-conv", "upper_case,lower_case")
		cmd.Stdin = strings.NewReader(testInput)
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
		assert.Zero(t, stdout.Len())
	})
}
