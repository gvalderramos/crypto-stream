# CryptoStream

CryptoStream is a **real-time cryptocurrency price monitoring and event-based processing system**.  
It collects and processes market data from multiple exchanges using **Go** and an event-driven architecture.

![Crypto Stream Architecture](https://github.com/gvalderramos/crypto-stream/blob/main/crypto_stream.png?raw=true)

## 🚀 Features
- **Exchange Watchers**: Monitor real-time prices from Binance, CoinGecko, and Kraken.
- **Event Management**: Handle and process crypto-related events.
- **Database Management**: Store and manage processed crypto data.
- **Crypto Stream API**: Provide an interface for other applications to consume the data.

## 🏗️ Project Structure
```
crypto-stream/
│── go.mod
│── go.sum
│── cmd/                # Executables
│   ├── binance_watcher/
│   │   ├── main.go
│   ├── coingecko_watcher/
│   │   ├── main.go
│   ├── kraken_watcher/
│   │   ├── main.go
│   ├── event_mng/
│   │   ├── main.go
│   ├── crypto_stream/
│   │   ├── main.go
│── internal/           # Shared libraries
│   ├── crypto_stream/
│   │   ├── api.go
│   │   ├── client.go
│   │   ├── models.go
│── README.md
│── LICENSE
```

## ⚡ Getting Started
### **1️⃣ Install Dependencies**
Ensure you have **Go installed**:
```bash
go version
```
If not, install it from [Go's official website](https://golang.org/dl/).

Clone this repository:
```bash
git clone https://github.com/yourusername/crypto-stream.git
cd crypto-stream
```

Download dependencies:
```bash
go mod tidy
```

---

### **2️⃣ Running a Watcher**
Each watcher is a separate executable. To run one, use:
```bash
go run cmd/binance_watcher/main.go
```
To build and run:
```bash
go build -o bin/binance_watcher cmd/binance_watcher/main.go
./bin/binance_watcher
```

---

### **3️⃣ Running All Services**
To run all services at once:
```bash
go run cmd/binance_watcher/main.go &
go run cmd/coingecko_watcher/main.go &
go run cmd/kraken_watcher/main.go &
go run cmd/event_mng/main.go &
go run cmd/crypto_stream/main.go &
```

---

## ⚙️ Configuration
Set environment variables before running:
```bash
export BINANCE_API_KEY="your_api_key"
export BINANCE_API_SECRET="your_api_secret"
```
Or use a `.env` file.

---

## 🤝 Contributing
1. Fork the repo  
2. Create a new branch (`feature-xyz`)  
3. Commit changes  
4. Open a pull request  

---

## 📝 License
This project is licensed under the **MIT License** – see the [LICENSE](LICENSE) file for details.

