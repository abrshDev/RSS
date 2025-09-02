# 📰 RSS Aggregator

A backend service for aggregating and serving RSS feeds.  
Built in **Go**, using **Fast, lightweight chi router**, **PostgreSQL**, and **sqlc** for type-safe queries.

---

## 🚀 Features
- Create and manage users with API keys
- Add RSS feeds to the system
- Follow/unfollow feeds as a user
- Background scraper fetches latest posts
- Retrieve latest posts for the feeds you follow
- JSON-based REST API with versioned routes (`/v1`)

---

## 🛠 Tech Stack
- **Go (chi router, cors, net/http)**
- **PostgreSQL**
- **sqlc** (auto-generates type-safe DB code)
- **Docker** (optional, for running Postgres easily)
- **godotenv** (for environment variables)

---

## ⚙️ Setup

### 1. Clone the repo
```bash
git clone https://github.com/abrshDev/RSS.git
cd RSS
