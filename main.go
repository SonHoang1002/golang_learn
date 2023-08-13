package main

import (
	"awesomeProject/server_side"
)

func main() {

	//input := [][]interface{}{{[]interface{}{1, 2, 3, 4, 5, 6}, 2}}
	//result := expandArray(input[0][0].([]interface{}), input[0][1].(int))
	//fmt.Println(result)

	//server_side.CreateServer()
	//server_side.StaticServer()
	server_side.RouteServer()

}

func expandArray(arr []interface{}, step int) [][]interface{} {
	result := make([][]interface{}, len(arr)-step+1)

	for i := 0; i <= len(arr)-step; i++ {
		subArray := []interface{}{}
		for j := i; j < i+step; j++ {
			subArray = append(subArray, arr[j])
		}
		subArray = append(subArray, arr[i+step-1])
		result[i] = subArray
	}

	return result
}
