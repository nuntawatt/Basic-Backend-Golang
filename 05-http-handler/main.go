package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Request รับ JSON Body
type Request struct {
	Name string `json:"name"`
}

// Response ส่งกลับ JSON Body
type Response struct {
	Message string `json:"message"`
}

func main() {
	// ตั้งค่า HTTP Handler สำหรับ path /hello
	http.HandleFunc("/hello", helloHandler)

	// แจ้งว่าเซิร์ฟเวอร์กำลังทำงาน
	fmt.Println("Server is running on http://localhost:8080")

	// เริ่มเซิร์ฟเวอร์ที่พอร์ต 8080
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// รับเฉพาะ POST เท่านั้น
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // ส่ง Error 405 ถ้าไม่ใช่ POST
		return
	}

	// แกะ JSON Body
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest) // ส่ง Error 400 ถ้า JSON ไม่ถูกต้อง
		return
	}

	// ตอบกลับเป็น JSON
	w.Header().Set("Content-Type", "application/json") // ตั้ง Header ว่าตอบกลับเป็น JSON
	if err := json.NewEncoder(w).Encode(Response{
		Message: fmt.Sprintf("Hello, %s", req.Name), // สร้างข้อความตอบกลับโดยใช้ชื่อที่ได้รับจาก JSON Request
	}); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError) // ส่ง Error 500 ถ้าเกิดปัญหาในการเข้ารหัส JSON
		return
	}
}
