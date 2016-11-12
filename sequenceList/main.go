package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	nums := []int{1, 2, 3, 4, 5, 6, 11, 22, 33, 44, 55, 66}
	key, err := Find(nums, 55)
	if err != nil {
		panic(err)
	}
	fmt.Println(key)

	nums = Add(nums, 77)
	fmt.Println(nums)

	nums, err = Del(nums, 11)
	if err != nil {
		panic(err)
	}
	fmt.Println(nums)

	fmt.Println(len(nums))

	nums, err = Mod(nums, 11, 89)
	if err != nil {
		panic(err)
	}
	fmt.Println(nums)
}

func Find(data []int, findData int) (int, error) {
	for singleKey, singleData := range data {
		if singleData == findData {
			return singleKey, nil
		}
	}
	return 0, errors.New("找不到该数字")
}

func Add(data []int, addData int) []int {
	return append(data, addData)
}

//删除第几个
func Del(data []int, delKey int) ([]int, error) {
	result := []int{}
	for singleKey, _ := range data {

		if singleKey == delKey {

			if singleKey == 0 {
				result = data[1:]
				return result, nil
			}

			result = data[:singleKey]
			result = append(result, data[singleKey+1:]...)
			return result, nil
		}
	}
	return data, errors.New("找不到该数字")
}

func Mod(data []int, modKey int, modData int) ([]int, error) {
	if modKey >= len(data) {
		return data, errors.New("修改超出了范围了")
	}

	data[modKey] = modData
	return data, nil
}
