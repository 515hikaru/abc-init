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

func createFile(content []byte, fileName string, fileType string) error {
	fileType = "cpp" // TODO: support other language
	fileNameWithExt := fmt.Sprintf("%s.%s", fileName, fileType)
	return ioutil.WriteFile(fileNameWithExt, content, 0666)
}

func run(c *cli.Context) error {
	b := []byte(contents)
	if fileName := c.String("output"); fileName != "" {
		err := createFile(b, fileName, "cpp")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		return nil
	}
	for _, fileName := range fileNames {
		err := createFile(b, fileName, "cpp")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
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
	}
}
