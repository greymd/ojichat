package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/greymd/ojichat/generator"
)

var appVersion = "0.0.1"
var usage = `Usage:
  ojichat [options] [<name>]

Options:
  -h, --help                        Show this screen.
  -V, --version                     Show version.
  -e <num>, --emoji=<num>           Maximum number of continuous Emojis [default: 4].
  -p <level>, --punctuation=<level> Punctuation level [default: 1].`

// TODO: --type おじさんタイプ (絵文字乱用, 顔文字乱用, 句読点, 若作り)

func main() {
	parser := &docopt.Parser{
		OptionsFirst: false,
	}
	args, _ := parser.ParseArgs(usage, nil, appVersion)
	config := generator.Config{}
	args.Bind(&config)

	result, err := generator.Start(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)
}
