package main

import (
	"fmt"
	"sync"
)

// SafeCounter เป็น struct ที่มี mutex เพื่อป้องกันการเข้าถึงพร้อมกันของหลาย goroutine
type SafeCounter struct {
	mu    sync.RWMutex // ล็อคเพื่อป้องกันการเข้าถึงพร้อมกัน
	count int          // ค่าที่นับ
}

// method Inc จะเพิ่มค่า count โดยล็อคก่อนเพื่อป้องกัน race condition
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock() // ปลดล็อคเมื่อทำงานเสร็จ
	c.count++
}

// method Value จะอ่านค่า count โดยล็อคแบบอ่านเพื่อให้สามารถอ่านได้พร้อมกันหลาย goroutine
func (c *SafeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock() // ปลดล็อคหลังอ่านเสร็จ
	return c.count
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// เปิด 1000 goroutine พร้อมกัน
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	// รอทุก goroutine เสร็จก่อน
	wg.Wait()
	
	// พิมพ์ค่าที่นับได้ ซึ่งควรจะเป็น 1000
	fmt.Println("Final count:", counter.Value())
}
