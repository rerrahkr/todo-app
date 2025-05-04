# Todo-App

A simple todo list web app.

Try to make an app with React (TypeScript + Vite) + Go + PostgreSQL.

You need to prepare an .env file with the following environment variables:

```sh
POSTGRES_DB=dbname
POSTGRES_USER=username
POSTGRES_PASSWORD=password
POSTGRES_PORT=5432
POSTGRES_HOST=host
POSTGRES_URI=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}

BACKEND_PORT=8080
VITE_BACKEND_URL=http://localhost:${BACKEND_PORT}

FRONTEND_PORT=5173
FRONTEND_URI=http://localhost:${FRONTEND_PORT}
```
