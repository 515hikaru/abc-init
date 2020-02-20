"""
entry point
"""
import pathlib


__version__ = "0.1.0"


def main():
    current = pathlib.Path(".")
    for char in ["A", "B", "C", "D", "E", "F"]:
        file_path = current / f"{char}.cpp"
        with file_path.open("w") as f:
            f.write(
                """#include <iostream>
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
"""
            )
