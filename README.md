# 📘 Flip Bank Statement Viewer — Backend (Golang)

Backend service untuk memproses file CSV transaksi, menghitung balance, dan menampilkan transaksi bermasalah. Dibangun dengan **Golang**, menerapkan **Clean Architecture**, mendukung **Docker**, serta dilengkapi **GitHub Actions CI**.

---

## ⚙️ Tech Stack
- Go 1.20+
- Clean Architecture
- Net/HTTP
- Docker (Multi-stage)
- GitHub Actions (Build + Test)

---

## 📁 Project Structure
```ssh
├── cmd/
│ └── server/
│ └── main.go
├── internal/
│ ├── handler/ # HTTP handlers
│ ├── model/ # Struct + enums
│ ├── repository/ # Storage interface
│ ├── service/ # Business logic
│ └── storage/ # CSV parser
├── go.mod
├── go.sum
└── Dockerfile
```

---

# 🚀 Running the App (Local)

### 1. Clone repository
```sh
git clone <your-repo-url>
cd flip-bank-statement-viewer
```
### 2. Install dependencies
```sh
    go mod tidy
```
### 3. Run Server
```ssh
    go run cmd/server/main.go
```
server berjalan di:
```ssh 
    http://localhost:8080
```

# TESTING API
Upload CSV
```ssh 
curl -X POST -F "file=@sample.csv" http://localhost:8080/upload
```
Get Balance
```ssh
curl http://localhost:8080/balance
```
Get Issues
```ssh
curl http://localhost:8080/issues
```

# 🐳 Running with Docker
### 1. Build Image
```ssh
docker build -t flip-bank-viewer .
```
### 2. Run container
```ssh
docker run -p 8080:8080 flip-bank-viewer
```
### 3. Test
```ssh
http://localhost:8080/balance
```

# 📦 GitHub Actions (CI Pipeline)
### 1. go-ci.yml — Build & Test
Berjalan otomatis pada:
    push ke branch main
    pull_request ke branch main

Pipeline terdiri dari:
```ssh
    go mod download
    go test ./...
    go build ./cmd/server
```

Tujuan:
    Memastikan kode valid dan bisa di-compile
    Unit test harus lulus sebelum merge

### 2. docker.yml — Docker Build (Conditional)
Workflow hanya berjalan jika ada perubahan pada file:
Dockerfile
```ssh
    .github/workflows/docker.yml
```
Pipeline:
```ssh
    docker build -t flip-bank-viewer .
```
Tujuan:
    Validasi Dockerfile hanya saat dibutuhkan
    Mengurangi waktu CI

# 🧪 Unit Tests
Jalankan: 
```ssh
    go test ./...
```
Unit test mencakup:
    Validasi CSV
    Validasi enum TYPE & STATUS
    Hitung balance
    Deteksi transaction issues

# 📌 API Endpoints
| Method | Endpoint   | Deskripsi                       |
| ------ | ---------- | ------------------------------- |
| POST   | `/upload`  | Upload file CSV                 |
| GET    | `/balance` | Mendapatkan total balance       |
| GET    | `/issues`  | List transaksi FAILED & PENDING |

# 🎯 Status Fitur
| Fitur                    | Status |
| ------------------------ | ------ |
| Upload + Validasi CSV    | ✔      |
| Enum Type/Status         | ✔      |
| Hitung Balance           | ✔      |
| Issues (FAILED, PENDING) | ✔      |
| Unit Test                | ✔      |
| Docker Support           | ✔      |
| CI Pipeline              | ✔      |
