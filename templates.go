package main

type languageTemplate struct {
	fileType       string
	initialContent string
}

var cppTemplate = languageTemplate{
	fileType: "cpp",
	initialContent: `#include <iostream>
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
`,
}
