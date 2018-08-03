package main

import "fmt"

func main() {
	// 定义
	map1 := map[string]string {
		"key1" : "value1",
		"key2" : "value2",
		"key3" : "value3",
	}
	fmt.Println(map1)

	map2 := make(map[string]string) // map2 == empty map
	fmt.Println(map2)

	var map3 map[string]string; // map3 == nil
	fmt.Println(map3)

	// 遍历
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	for k := range map1 {
		fmt.Println(k)
	}

	for _, v := range map1 {
		fmt.Println(v)
	}

	// 取值
	value1 := map1["key1"]
	fmt.Println(value1)

	value4 := map1["key4"]
	fmt.Printf("如果key不存在，value输出zero value：%x\n",value4)

	value1, ok1 := map1["key1"]
	fmt.Println(value1, ok1)

	value4, ok4 := map1["key4"]
	fmt.Printf("如果key不存在，value输出zero value，ok4表示key不存在：%x, %v\n",value4, ok4)

	if value4, ok4 := map1["key4"]; ok4 {
		fmt.Println("key存在，value=", value4)
	} else {
		fmt.Println("key不存在")
	}

	// 删除
	delete(map1, "key3")

	len := len(map1)
	fmt.Println(len)
}
