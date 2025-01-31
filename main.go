package main

import (
	"flag"
	"fmt"
	"os"
)

type Greeting struct {
	prefix string
	hello  string
}

var greetings = map[string]Greeting{
	"en": {"Mr/Ms", "Hello"},
	"ja": {"さん", "こんにちは"},
	"fr": {"M/Mme", "Bonjour"},
	"es": {"Sr/Sra", "Hola"},
	"de": {"Herr/Frau", "Hallo"},
}

type CLI struct {
	outStream, errStream *os.File
	name                 string
	lang                 string
}

func NewCLI(outStream, errStream *os.File) *CLI {
	return &CLI{outStream: outStream, errStream: errStream}
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("multilingual-hello", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	flags.StringVar(&c.name, "name", "", "your name")
	flags.StringVar(&c.lang, "lang", "en", "language (en/ja/fr/es/de)")

	if err := flags.Parse(args[1:]); err != nil {
		return 1
	}

	if c.name == "" {
		fmt.Fprintln(c.errStream, "Error: name is required")
		return 1
	}

	greeting, ok := greetings[c.lang]
	if !ok {
		fmt.Fprintf(c.errStream, "Error: unsupported language: %s\n", c.lang)
		return 1
	}

	fmt.Fprintf(c.outStream, "%s, %s %s!\n", greeting.hello, c.name, greeting.prefix)
	return 0
}

func main() {
	cli := NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
