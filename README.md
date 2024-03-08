# Go API

This is an API developed in Go 1.22.1

## Env variables
- PORT: Port in which the API will be listening
- DATABASE_URL: Connection URL to Postgres database

## Endpoints
- POST /user. Create user
    - Payload: name string, surname string
- GET /user/{id}. Get user by id
- POST /project. Create post
    - Payload: name string, description string
- POST /bug. Create bug
    - Payload: user int, project int, description string
- GET /bug/{id}. Get bug by id
- GET /bug. List bugs filtered by user_id, project_id, start_date and end_date query params
