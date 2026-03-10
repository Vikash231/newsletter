# InShorts Clone

A news platform clone inspired by InShorts — delivering news in 60 words or less.

**Stack:** Go (backend) + React + Vite (frontend)

## Features

- News cards with 60-word summaries
- Category filtering: Technology, Sports, Business, Science, Entertainment, Health
- Infinite scroll with intersection observer
- Bookmark & share articles
- Skeleton loading states
- Fully responsive layout
- REST API with pagination

## Project Structure

```
newsletter/
├── backend/          # Go HTTP server (stdlib only)
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── frontend/         # React + Vite SPA
│   ├── src/
│   │   ├── App.jsx
│   │   ├── components/
│   │   │   ├── Header.jsx / .css
│   │   │   ├── CategoryTabs.jsx / .css
│   │   │   ├── NewsCard.jsx / .css
│   │   │   └── NewsFeed.jsx / .css
│   │   └── hooks/
│   │       └── useNews.js
│   ├── index.html
│   ├── vite.config.js
│   └── Dockerfile
└── docker-compose.yml
```

## Quick Start

### Backend (Go)

```bash
cd backend
go run main.go
# API runs at http://localhost:8080
```

### Frontend (React)

```bash
cd frontend
npm install
npm run dev
# App runs at http://localhost:3000
```

### Docker Compose

```bash
docker-compose up --build
# App at http://localhost:3000, API at http://localhost:8080
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/health` | Health check |
| GET | `/api/categories` | List all categories |
| GET | `/api/news` | Get news (supports `?category=`, `?page=`, `?limit=`) |
| GET | `/api/news/:id` | Get single article |

### Example Requests

```bash
# All news
curl http://localhost:8080/api/news

# Filter by category with pagination
curl "http://localhost:8080/api/news?category=technology&page=1&limit=5"

# Single article
curl http://localhost:8080/api/news/1

# Categories
curl http://localhost:8080/api/categories
```