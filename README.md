# ThaiBrave API

A Go/Gin HTTP API backed by PostgreSQL. The API serves quiz questions and answers and accepts submitted user answers.

## Requirements

- Docker with Docker Compose (recommended), or Go 1.27 and PostgreSQL 15 or compatible
- The PostgreSQL database and credentials configured in `app.env`

## Run with Docker Compose

From this directory, start the API and PostgreSQL:

```sh
docker compose up --build
```

The API listens on [http://localhost:8080](http://localhost:8080). Compose starts PostgreSQL with database `thaibrave`, user `postgres`, and password `postgres`. PostgreSQL is published to the host on port `54321`; the API container connects to the Compose service on port `5432`. The API source directory is mounted into the development container, and CompileDaemon polls it every second to rebuild and restart the API when Go files change. Polling is used because filesystem notifications can be unreliable for Windows-hosted Docker mounts. After changing the development image, recreate it with `docker compose up --build`.

Stop the services with:

```sh
docker compose down
```

The Compose file does not mount a persistent PostgreSQL data volume, so database contents may be lost when the PostgreSQL container is removed.

## Run locally

1. Start PostgreSQL and create a database and user matching the settings in `app.env`. The default local settings expect PostgreSQL at `localhost:54321`.
2. From this directory, run:

   ```sh
   go run ./cmd/main.go
   ```

   The process reads `app.env` from its current working directory, connects to PostgreSQL, and listens on port `8080`.

The repositories automatically create the `questions`, `answer_choices`, and `users` tables on startup. The project does not currently seed question or answer data; add records to those tables before expecting the question endpoint to return quiz content.

### Database configuration

`app.env` contains the local database settings:

```dotenv
DB_HOST=localhost
DB_NAME=thaibrave
DB_PORT=54321
DB_USER=postgres
DB_PASSWORD=postgres
```

When using Compose, the API service overrides `DB_HOST` and `DB_PORT` to `postgres` and `5432`, respectively. Adjust the other values in `app.env` if you change the PostgreSQL credentials or database name.

## HTTP API

For additional request/response examples, import the Postman collection at [ThaiBev.postman_collection.json](./ThaiBev.postman_collection.json). It includes sample calls for health checks, creating questions, fetching questions, and submitting user answers.

All routes are under `/api`. Requests that include a body should use `Content-Type: application/json`.

### Get questions by ID

```http
POST /api/question/getQuestionsbyIDList
Content-Type: application/json
```

The request body is an array of question IDs:

```json
[1, 2]
```

Example:

```powershell
curl.exe -X POST http://localhost:8080/api/question/getQuestionsbyIDList -H "Content-Type: application/json" -d "[1,2]"
```

On success, returns `200 OK` with an array of questions. Each question contains its ID, text, and answer choices:

```json
[
  {
    "id": 1,
    "question": "Example question?",
    "answers": [
      {
        "id": 10,
        "answer": "Example answer",
        "is_correct": false
      }
    ]
  }
]
```

The answer service currently includes extra empty answer objects in this list; the example shows the intended question/answer shape. A malformed JSON body returns `400 Bad Request`. A database/service error returns `500 Internal Server Error`.

### Get user information

```http
GET /api/user/{id}
```

Example:

```powershell
curl.exe http://localhost:8080/api/user/1
```

This route currently returns `200 OK` with the JSON string `"nil"` for any ID; user lookup is not implemented yet.

### Save user information

```http
POST /api/user/save
Content-Type: application/json
```

The request body accepts a name and a list of submitted answer IDs:

```json
{
  "name": "Jane Doe",
  "answer_list": [
    { "id": 10 },
    { "id": 11 }
  ]
}
```

Example:

```powershell
curl.exe -X POST http://localhost:8080/api/user/save -H "Content-Type: application/json" -d '{"name":"Jane Doe","answer_list":[{"id":10},{"id":11}]}'
```

The score is calculated from the answer IDs in the database; the submitted `is_correct` value, if present, is not used. On success, returns `200 OK`:

```json
{
  "message": "User info saved successfully"
}
```

Currently, the implementation saves the name `"John Doe"` rather than the submitted name. Invalid JSON returns `400 Bad Request`; a database/service error (for example, an answer ID that cannot be found) returns `500 Internal Server Error`.
