package main

import "fmt"

func testMap(){
	fmt.Println(" test Map")
	m:= map[string]string{
		//key 在map 中是无序的，打印顺序不一致
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"qualiity":"not bad",
	}
	fmt.Println(m)

	map2 := make(map[string]int) //map2 == nil = false
	fmt.Println(map2,map2==nil)

	var map3 map[string]string //map3 == nil = true
	fmt.Println(map3,map3==nil)
}

func testIterationMap(tmpMap map[string]string){
	for k,v := range tmpMap{
		fmt.Println(k,v)
	}
}

func testDeleteMapElem(){
	fmt.Println("Test delete map element")
	m:= map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"qualiity":"not bad",
	}
	fmt.Println(m)
	delete(m,"name")
	fmt.Println(m)
}

func main() {
	testMap()
	m:= map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"qualiity":"not bad",
	}
	testIterationMap(m)

	string1,ok := m["course"]
	fmt.Println(string1,ok)

	string2,ok := m["courseaa"]
	fmt.Println(string2,ok)

	//判断键是否存在
	if courseName,ok := m["courseaaa"];ok{
		fmt.Println(courseName)
	}else{
		fmt.Println("courseaaa not in map")
	}

	testDeleteMapElem()
}
