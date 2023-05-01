# Merge 

## Description

This is a simple library to merge two or more map[string]interface{} into one, recursively. It is useful when you need to merge two or more maps with different structures.

## Usage

```go

package main

import (
	"fmt"
	"github.com/illarion/merge"
	"encoding/json"
)

func main() {

	var src1 map[string]interface{}
	var src2 map[string]interface{}
	var src3 map[string]interface{}

	json.Parse([]byte(`{"a": 1, "b": 2}`), &src1)
	json.Parse([]byte(`{"c": 3, "d": 4}`), &src2)
	json.Parse([]byte(`{"e": 5, "f": 6}`), &src3)

	result := merge.Maps(nil, src1, src2, src3)
	fmt.Sprintf("%#v", result)
}
```
