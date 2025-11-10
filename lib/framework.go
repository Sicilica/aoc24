package lib

import (
	"fmt"
	"reflect"
)

func Day(
	input any,
	parts ...any,
) {
	inputV := reflect.ValueOf(input)
	Assert(inputV.Kind() == reflect.Func)
	inputs := inputV.Call(nil)

	for i, p := range parts {
		partV := reflect.ValueOf(p)
		Assert(partV.Kind() == reflect.Func)
		ret := partV.Call(inputs)
		Assert(len(ret) == 1)
		fmt.Println("part", i+1, "result:", ret[0].Interface())
	}
}
