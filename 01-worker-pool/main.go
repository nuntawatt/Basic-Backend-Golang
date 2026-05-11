package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// รันโปรแกรมด้วยจำนวน worker 3 คน และจำนวนงาน 5 งาน
	RunWorkerPool(3, 5)
}

func RunWorkerPool(numWorkers int, numJobs int) {
	// คิวงาน buffered ให้ส่งงานได้โดยไม่ต้องรอให้ worker พร้อม
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// เปิด worker ตามจำนวนที่กำหนด
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // บอกว่ามี Goroutine เพิ่มขึ้น 1
		go func(workerID int) {
			defer wg.Done() // แจ้งเมื่อ Goroutine นี้ทำงานเสร็จ
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
				time.Sleep(time.Second)
			}
		}(i) // ส่ง i เข้ามาตรงๆ ป้องกัน closure bug
	}

	// ส่งงานเข้าคิว
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // ปิดคิวบอก worker ว่าหมดงานแล้ว

	// รอทุก worker ทำงานเสร็จ
	wg.Wait()
	fmt.Println("All jobs completed")
}
