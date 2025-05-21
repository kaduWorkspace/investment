# 💸 Investment API

A lightweight and extensible backend service for managing investment accounts and their events (deposits, withdrawals, transfers, and more). Built with Go.

---

## 🚀 Features

- Create and manage investment accounts
- Register different types of account events (e.g. deposits, withdrawals)
- Track account balances over time
- RESTful API with clean JSON responses
- Fully tested core logic

---

## 📦 Requirements

- Go 1.21+
- Make (for running dev/test commands)

---

## 🛠️ Running Locally

```bash
# Clone the repo
git clone https://github.com/kaduWorkspace/investment.git
cd investment

# Run the API
make run
```

The server should be running at: [http://localhost:8080](http://localhost:8080)

---

## 🧪 Running Tests

```bash
make test
```

All core business logic is tested, including account creation, event handling, and balance tracking.

---

## 📬 API Endpoints (WIP)

| Method | Endpoint         | Description                  |
|--------|------------------|------------------------------|
| POST   | `/accounts`      | Create a new account         |
| GET    | `/accounts/:id`  | Get account by ID            |
| POST   | `/events`        | Register a new event         |
| GET    | `/balance/:id`   | Get current account balance  |

> You can use tools like Postman or cURL to interact with the API.

---

## 📁 Project Structure

| Folder         | Description                      |
|----------------|----------------------------------|
| `cmd/api`      | API setup and HTTP routing       |
| `core/domain`  | Business entities and interfaces |
| `core/usecase` | Application logic                |
| `core/service` | Concrete service implementations |
| `tests`        | Unit tests                       |

---

## 👨‍💻 Author

Developed by [@kaduWorkspace](https://github.com/kaduWorkspace)

---

## 📝 License

This project is open-source and available under the [MIT License](LICENSE).

