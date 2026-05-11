package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	// เรียกใช้ฟังก์ชัน TwoSum เพื่อหาคู่ตัวเลขที่บวกกันได้ target
	result, found := TwoSum(nums, target)

	if found {
		// เช็คว่าเคยเจอ complement นี้มาก่อนหรือยัง
		fmt.Println("Result:", result)
	} else {
		// ถ้าไม่เคยเจอคู่ที่บวกกันได้ target ให้พิมพ์ข้อความว่าไม่พบ
		fmt.Println("No two sum found")
	}
}

// method TwoSum จะรับ slice ของตัวเลขและ target แล้วคืน index ของคู่ตัวเลขที่บวกกันได้ target
func TwoSum(nums []int, target int) ([]int, bool) {
	// เก็บตัวเลขที่เคยเจอใน map เพื่อให้ค้นหาได้เร็วขึ้น
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num // คำนวณตัวเลขที่ต้องการเพื่อให้ได้ target

		// เช็คว่าเคยเจอ complement นี้มาก่อนหรือยัง
		if j, found := seen[complement]; found {
			return []int{j, i}, true // ถ้าเจอคู่ที่บวกกันได้ target ให้คืน index ของทั้งสองตัวเลขออกมา
		}

		// เก็บค่าปัจจุบันไว้ เพื่อให้รอบถัดไปตรวจสอบย้อนหลังได้
		seen[num] = i
	}

	// ไม่พบคู่ที่รวมกันได้ target
	return nil, false
}
