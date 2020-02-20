package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	fileNames = []string{"A", "B", "C", "D", "E", "F"}
	contens   = `#include <iostream>
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

using namespace std;
int main() {

    return 0;
}
`
)

func main() {
	b := []byte(contens)
	for _, fileName := range fileNames {
		fileNameWExt := fmt.Sprintf("%s.cpp", fileName)
		err := ioutil.WriteFile(fileNameWExt, b, 0666)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
