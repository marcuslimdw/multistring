package main

import (
	"bufio"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"

	"multistring/internal"
)

//go:generate go run . ../../multistring_gen.go

const preamble = `package multistring

import (
	"strings"
	"unicode"
)

`

func main() {
	if len(os.Args) < 2 {
		panic("must provide destination path")
	}

	bc := generationContext{}
	bc.generate(os.Args[1])
	if bc.err != nil {
		panic(bc.err.Error())
	}
}

type generationContext struct {
	destFile *os.File
	err      error
	writer   *bufio.Writer
}

func (c *generationContext) generate(destPath string) {
	c.openDestinationFile(destPath)
	defer c.cleanup()

	c.writePreamble()

	stringsDecls := c.parseBuiltinStringsDecls()
	for _, decl := range internal.ExtractExportedFuncDecls(stringsDecls) {
		transformedDecl := internal.Transform(decl)
		c.writeFuncDecl(transformedDecl)
		c.writeFuncFooter()
	}
}

func (c *generationContext) openDestinationFile(destPath string) {
	if c.err != nil {
		return
	}

	file, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		c.err = fmt.Errorf("failed to open destination file: %w", err)
	}

	c.destFile = file
	c.writer = bufio.NewWriter(file)
}

func (c *generationContext) parseBuiltinStringsDecls() []ast.Decl {
	if c.err != nil {
		return nil
	}

	goroot, exists := os.LookupEnv("GOROOT")
	if !exists {
		c.err = errors.New("GOROOT not set")
		return nil
	}

	path := fmt.Sprintf("%s/src/strings/strings.go", goroot)
	file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		c.err = fmt.Errorf("failed to load strings source code: %w", err)
		return nil
	}

	return file.Decls
}

func (c *generationContext) writePreamble() {
	if c.err != nil {
		return
	}

	_, c.err = c.writer.WriteString(preamble)
}

func (c *generationContext) writeFuncDecl(decl *ast.FuncDecl) {
	if c.err != nil {
		return
	}

	defer func() {
		if v := recover(); v != nil {
			c.err = fmt.Errorf("recovered from panic while writing transformed function %s: %s", decl.Name, v)
		}
	}()

	if err := printer.Fprint(c.writer, token.NewFileSet(), decl); err != nil {
		c.err = err
	}
}

func (c *generationContext) writeFuncFooter() {
	if c.err != nil {
		return
	}

	_, c.err = c.writer.WriteString("\n\n")
}

func (c *generationContext) cleanup() {
	if c.err != nil || c.writer == nil || c.destFile == nil {
		return
	}

	if err := c.writer.Flush(); err != nil {
		c.err = fmt.Errorf("failed to write to file: %w", err)
		return
	}

	offset, err := c.destFile.Seek(0, io.SeekCurrent)
	if err != nil {
		c.err = fmt.Errorf("failed to get current position: %w", err)
		return
	}

	if err := c.destFile.Truncate(offset); err != nil {
		c.err = fmt.Errorf("failed to truncate file: %w", err)
		return
	}

	if err := c.destFile.Close(); err != nil {
		c.err = fmt.Errorf("failed to close file: %w", err)
	}
}
