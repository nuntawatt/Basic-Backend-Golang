# Go Assessment

ชุดแบบทดสอบ Backend Developer ด้วยภาษา Go มี 5 ข้อ

## โครงสร้างโปรเจค

```
go-assessment/
├── 01-worker-pool/       # Basic Concurrency
├── 02-safe-counter/      # Thread-Safety
├── 03-shape-interface/   # Interface
├── 04-logic-map/         # Algorithm
└── 05-http-handler/      # HTTP API
```

---

## การติดตั้งและรันโปรแกรม

```bash
# 1. Clone โปรเจค
git clone https://github.com/username/go-assessment.git
cd go-assessment

# 2. รันแต่ละข้อ
cd 01-worker-pool && go run main.go
cd 02-safe-counter && go run main.go
cd 03-shape-interface && go run main.go
cd 04-logic-map && go run main.go
cd 05-http-handler && go run main.go
```

> หมายเหตุ: โปรเจคนี้ใช้เฉพาะ Standard Library ไม่ต้องรัน `go mod tidy`

---

## 01 - Worker Pool

สร้าง Worker Pool ที่รันแบบ Concurrent โดยใช้ Goroutine, Buffered Channel และ WaitGroup

---

## 02 - Safe Counter

SafeCounter ที่รองรับ 1,000 Goroutine เรียกใช้พร้อมกันโดยไม่เกิด Race Condition ใช้ `sync.RWMutex` แยกการล็อคระหว่างอ่านและเขียน

---

## 03 - Shape Interface

ระบบคำนวณพื้นที่ผ่าน Interface โดย Rectangle และ Circle implement `Area()` และรับผ่าน `PrintArea(s Shape)` ได้ทันที

---

## 04 - Two Sum

หาคู่ index ที่บวกกันได้เท่ากับ target ด้วย HashMap แทน Double Loop

---

## 05 - HTTP Handler

REST API ด้วย `net/http` รองรับ POST `/hello` และตอบกลับเป็น JSON

---