# 🎬 Reelix

Reelix is a RESTful API built with **Go (Golang)** and **Gin Gonic** — a fast and lightweight web framework.

## 🏗️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin Gonic
- **Database:** MongoDB
- **ORM/Driver:** MongoDB Driver for Go
- **Authentication:** JWT (JSON Web Tokens)


## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/reelix.git
cd reelix
````


## 🧩 API Endpoints

| Method | Endpoint          | Description            | Auth Required |
| ------ | ----------------- | ---------------------- | ------------- |
| POST   | `/users/register` | Register a new user    | ❌             |
| POST   | `/users/login`    | Login user & get token | ❌             |
| GET    | `/movies`         | Get all movies         | ❌             |
| POST   | `/movies`         | Add a new movie        | ✅             |
| GET    | `/movies/:id`     | Get movie by ID        | ✅             |
| PUT    | `/movies/:id`     | Update a movie         | ✅             |
| DELETE | `/movies/:id`     | Delete a movie         | ✅             |
