package main

import (
	"fmt"
	"log"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/greymd/ojichat/generator"
)

var appVersion = `Ojisan Nanchatte (ojichat) command version 0.2.0
Copyright (c) 2019 Yamada, Yasuhiro
Released under the MIT License.
https://github.com/greymd/ojichat`

var usage = `Usage:
  ojichat [options] [<name>]

Options:
  -h, --help      ヘルプを表示.
  -V, --version   バージョンを表示.
  -e <number>     絵文字/顔文字の最大連続数 [default: 4].
  -p <level>      句読点挿入頻度レベル [min:0, max:3] [default: 0].`

// TODO: --type おじさんタイプ (絵文字乱用, 顔文字乱用, 句読点, 若作り)

func main() {
	parser := &docopt.Parser{
		OptionsFirst: true,
	}
	args, _ := parser.ParseArgs(usage, nil, appVersion)
	config := generator.Config{}
	err := args.Bind(&config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	result, err := generator.Start(config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", result)
}
