# Todo-App

This is a simple web application for managing a todo list.

The application consists of a frontend built with React, a backend implemented in Go, and a PostgreSQL database.

To run the application, create an `.env` file with the following environment variables:

```sh
POSTGRES_DB=dbname
POSTGRES_USER=username
POSTGRES_PASSWORD=password
POSTGRES_PORT=5432
POSTGRES_HOST=database
POSTGRES_URI=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}

BACKEND_PORT=8080
VITE_BACKEND_URI=/api

FRONTEND_PORT=80
FRONTEND_URI=http://frontend:${FRONTEND_PORT}
```

These 3 parts of the web app run in Docker containers and are managed with Docker Compose.

```sh
docker compose build
docker compose up -d
```

After starting the containers, the frontend is exposed at `http://localhost:8081`.
