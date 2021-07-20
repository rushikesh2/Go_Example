package primitive

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"bytes"

)

type Mode int

const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedrect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

// WithMode is option for Transform functionthat will define the
// mode you want to use. By Default, ModeTriangle willbe used
func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

func Transform(image io.Reader, numShapes int, opts ...func() []string) (io.Reader, error) {
	in, err := tempFile("in_", "jpg")
	if err != nil {
		return nil, errors.New("Failed to create input File")
	}
	defer os.Remove(in.Name())
	out, err := tempFile("out_", "jpg")
	if err != nil {
		return nil, errors.New("Failed to create output File")
	}
	defer os.Remove(out.Name())

	// Read image into File
	_, err = io.Copy(in, image)
	if err != nil {
		return nil, errors.New("Failed to copy image into input file")
	}

	// Run Primitive
	stdComb, err := primitive(in.Name(), out.Name(), numShapes, ModeRect)
	if err != nil {
		return nil, errors.New("Failed to run primitive command stdComb= %S")
	}
	fmt.Println(stdComb)

	// read out into reader
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, out)
	if err != nil {
		return nil, errors.New("Failed to copy output file into byte buffer")
	}

	return b, err
}

func primitive(inputFile, outputFile string, numShapes int, mode Mode) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d -m %d", inputFile, outputFile, numShapes, mode)
	cmd := exec.Command("primitive", strings.Fields(argStr)...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(b), err
}

func tempFile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", "in_")
	if err != nil {
		return nil, errors.New("Failed to create temp File")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}
