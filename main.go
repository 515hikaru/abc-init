package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "v0.0.3"

var (
	app        *cli.App
	targetName string
	fileNames  = []string{"A", "B", "C", "D", "E", "F"}
	contents   = `#include <iostream>
#include <cmath>
#include <string>
#include <algorithm>
#include <vector>
#include <stack>
#include <queue>
#include <functional>
#include <map>
#include <set>
#include <tuple>
#include <bitset>

using namespace std;
int main() {

    return 0;
}
`
)

func createFile(fileName string, template languageTemplate) error {
	fileNameWithExt := fmt.Sprintf("%s.%s", fileName, template.fileType)
	content := []byte(template.initialContent)
	return ioutil.WriteFile(fileNameWithExt, content, 0666)
}

func run(c *cli.Context) error {
	if fileName := c.String("output"); fileName != "" {
		err := createFile(fileName, cppTemplate)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
	for _, fileName := range fileNames {
		err := createFile(fileName, cppTemplate)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func init() {
	app = &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "",
				Usage:   "Place the output into `FILE`",
			},
		},
		Name:   "abc-init",
		Usage:  "Initialization Command For AtCoder Beginner Contest.",
		Action: run,
	}
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	os.Exit(0)
}
