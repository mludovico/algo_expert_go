package main

import (
	"fmt"
)

var inputArray = []interface{}{
	5,
	2,
	[]interface{}{
		7,
		-1,
	},
	3,
	[]interface{}{
		6,
		[]interface{}{
			-13,
			8,
		},
		4,
	},
}

func main() {
	result := ProductSum(inputArray)
	fmt.Println(result)
}

func ProductSum(array []interface{}) int {
	return productSum(array, 1)
}

func productSum(array []interface{}, level int) int {
	var acc int
	for _, val := range array {
		switch v := val.(type) {
		case int:
			acc += val.(int)
			break
		case []interface{}:
			acc += productSum(val.([]interface{}), level+1)
			break
		default:
			panic(fmt.Sprintf("Unexpected type %T", v))
		}
	}
	return level * acc
}
