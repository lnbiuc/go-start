package main

import (
	"fmt"
	"time"
)

func main() {
	typeStruct()
}

func typeArray() {
	var arr1 []int
	arr2 := []int{1, 2, 3}
	arr3 := [...]int{1, 2, 3}

	arr1 = append(arr1, 1)
	arr2 = append(arr2, 3)

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func typeMap() {
	dataMap1 := make(map[string]string)
	dataMap2 := map[string]string{}

	dataMap3 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	fmt.Println(dataMap1)
	fmt.Println(dataMap2)
	fmt.Println(dataMap3)

	// add
	dataMap1["key1"] = "value1"
	dataMap1["key2"] = "value2"
	dataMap1["key3"] = "value3"
	val := dataMap1["key3"]
	fmt.Println(val)
	// delete
	delete(dataMap1, "key1")

	fmt.Println(dataMap1)

	for k, v := range dataMap1 {
		fmt.Println(k, v)
	}
}

func typeStruct() {
	type Article struct {
		id, title, content string
		tags               []string
		createdAt          time.Time
		views              int
	}

	article1 := Article{}

	fmt.Println(article1)

	article2 := Article{
		id:        "1",
		title:     "Go Type Struct",
		content:   "CONTENT",
		tags:      []string{"tag1", "tag2"},
		createdAt: time.Now(),
		views:     10,
	}

	fmt.Println(article2)
	fmt.Println(article2.title)

	article2.id = "2"
	article2.createdAt = time.Now()
	article2.tags = []string{"tag4", "tag5"}
	fmt.Println(article2)
}
