package main

import "fmt"

func main() {
	/* 使用interface{} 初始化1个一维map
	* 关键点interface{} 可以代表任意类型
	* 原理： interface{} 就是1个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	 */
	mainMap := make(map[string]interface{})
	subMap := make(map[string]string)
	subMap["key_1"] = "sub_value1"
	subMap["key_2"] = "sub_value2"
	mainMap["Map"] = subMap
	for key, val := range mainMap {
		// 此处必须实例化接口类型即 *.(map[string]string)
		for subKey, subVal := range val.(map[string]string) {
			fmt.Printf("mapName = %s key = %s value= %s\n", key, subKey, subVal)
		}
	}
}
