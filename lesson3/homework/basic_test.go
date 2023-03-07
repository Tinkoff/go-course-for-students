package main

import (
	"context"
	"math"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const testInput = `

  
hELlO evEryOnE!
Машинное обучение – это наука о разработке алгоритмов и статистических моделей, которые компьютерные системы используют для выполнения задач без явных инструкций, полагаясь вместо этого на шаблоны и логические выводы.
Компьютерные системы используют алгоритмы машинного обучения для обработки больших объемов статистических данных и выявления шаблонов данных.
Таким образом, системы могут более точно прогнозировать результаты на основе заданного набора входных данных. 😊🎉💋😍😋
Например, специалисты по работе с данными могут обучить медицинское приложение диагностировать рак по рентгеновским изображениям, сохраняя миллионы отсканированных изображений и соответствующие диагнозы.

 

  
`

type unlimitedReader struct {
	input  []byte
	prefix []byte
	offset int
}

func (s *unlimitedReader) Read(p []byte) (n int, err error) {
	for idx := range p {
		if s.offset < len(s.prefix) {
			p[idx] = s.prefix[s.offset]
		} else {
			p[idx] = s.input[(s.offset-len(s.prefix))%len(s.input)]
		}
		s.offset++
	}
	return len(p), nil
}

func composeBinaryPath() string {
	binName := "go-course-2023-lesson3-tests"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	return path.Join(os.TempDir(), binName)
}

func TestBasicIntegration(t *testing.T) {
	binPath := composeBinaryPath()
	cmd := exec.Command("go", "build", "-o", binPath, "./")
	assert.NoError(t, cmd.Run())
	defer func() {
		assert.NoError(t, os.Remove(binPath))
	}()

	t.Run("ok with stdin input and stdout result", func(t *testing.T) {
		cmd = exec.Command(binPath)
		cmd.Stdin = strings.NewReader(testInput)
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, testInput, stdout.String())
	})

	t.Run("ok with unlimited stdin input, limit option and stdout result", func(t *testing.T) {
		limit := 1000000
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		cmd = exec.CommandContext(ctx, binPath, "-limit", strconv.Itoa(limit))

		cmd.Stdin = &unlimitedReader{input: []byte(testInput)}
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, ctx.Err(), "process timed out")
		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Len(t, stdout.String(), limit)
	})

	t.Run("ok with file input and stdout result", func(t *testing.T) {
		testFile, err := os.Create("in.txt")
		assert.NoError(t, err)
		defer os.Remove(testFile.Name())
		_, err = testFile.WriteString(testInput)
		assert.NoError(t, err)

		cmd = exec.Command(binPath, "-from", testFile.Name())
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err = cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, testInput, stdout.String())
	})

	t.Run("ok with stdin input and file result", func(t *testing.T) {
		testFileName := "out.txt"
		cmd = exec.Command(binPath, "-to", testFileName)
		cmd.Stdin = strings.NewReader(testInput)
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Zero(t, stdout.Len())

		data, err := os.ReadFile(testFileName)
		defer os.Remove(testFileName)
		assert.NoError(t, err)
		assert.Equal(t, testInput, string(data))
	})

	t.Run("ok with stdin input and stdout result, limit and offset options", func(t *testing.T) {
		limit := 100
		offset := 1200
		end := int(math.Min(float64(offset+limit), float64(len(testInput))))
		cmd = exec.Command(binPath, "-limit", strconv.Itoa(limit), "-offset", strconv.Itoa(offset))
		cmd.Stdin = strings.NewReader(testInput)
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.NoError(t, err)
		assert.Zero(t, stderr.Len(), stderr.String())
		assert.Equal(t, testInput[offset:end], stdout.String())
	})

	t.Run("error, offset greater than input size", func(t *testing.T) {
		cmd = exec.Command(binPath, "-offset", "100", "-limit", "1000")
		cmd.Stdin = strings.NewReader("test")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
		assert.Zero(t, stdout.Len())
	})

	t.Run("error with invalid limit", func(t *testing.T) {
		cmd = exec.Command(binPath, "-limit", "qweqwe")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
		assert.Zero(t, stdout.Len())
	})

	t.Run("error with invalid offset", func(t *testing.T) {
		cmd = exec.Command(binPath, "-offset", "-90")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
		assert.Zero(t, stdout.Len())
	})

	t.Run("error with non-existent input file", func(t *testing.T) {
		cmd = exec.Command(binPath, "-from", "non-exist.txt")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr

		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
		assert.Zero(t, stdout.Len())
	})

	t.Run("error with existing output file", func(t *testing.T) {
		cmd = exec.Command(binPath, "-to", "main.go")
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
