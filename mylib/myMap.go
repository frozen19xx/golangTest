package mytest

import "fmt"

func MapTest() {
	println("=====================Map loop=====================")
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	//

	/* 创建 map */
	//countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 删除元素 */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}
