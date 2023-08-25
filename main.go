package main

import (
	"fmt"
	"sync"
)

func main() {

	//input := [][]interface{}{{[]interface{}{1, 2, 3, 4, 5, 6}, 2}}
	//result := expandArray(input[0][0].([]interface{}), input[0][1].(int))
	//fmt.Println(result)

	//server_side.CreateServer()
	//server_side.GetMusicApi()

	//server_side.StaticServer()

	//server_side.RouteServer()

	//server_side.InitMySQLConnection()

	//if err := server_side.CreateUser("hts", 22); err != nil {
	//	fmt.Println("Error adding user:", err)
	//}

	//if err := server_side.DeleteUser(2); err != nil {
	//	fmt.Println("Error deleting user:", err)
	//}

	//if err := server_side.UpdateUser(3, "sco len cb", 22); err != nil {
	//	fmt.Println("Error deleting user:", err)
	//}

	//var allUser = server_side.GetAllUser()
	//for _, item := range allUser {
	//	fmt.Printf("User: username -> %s, age -> %d\n", item.Username, item.Age)
	//}

	//server_side.TemplateServer()
	//server_side.WebSocketServer()

	var (
		wg sync.WaitGroup
	)
	sum := 0
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for item := range list {
		wg.Add(1)
		sum += item
		wg.Done()
	}
	fmt.Printf("sum %d ", sum)

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
