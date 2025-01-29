# Go User API

A simple Go web server that manages user data in an in-memory cache with thread safety.

## Features
- Add a user (`POST /users`)
- Get a user by ID (`GET /users/{id}`)
- Get all users (`GET /users`)
- Delete a user (`DELETE /users/{id}`)
- Thread-safe cache implementation
- Abstraction of storage logic for decoupled database

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/go-user-api.git
   cd go-user-api
   ```
2. Run server:
   ```sh
   make run
   ```
3. Navigate to `http://localhost:4545` and call an endpoint

## API Endpoints

| Method   | Endpoint        | Description          |
|----------|----------------|----------------------|
| `GET`    | `/`            | Welcome message/health check     |
| `POST`   | `/users`       | Create a new user   |
| `GET`    | `/users`       | Get all users       |
| `GET`    | `/users/{id}`  | Get user by ID      |
| `DELETE` | `/users/{id}`  | Delete a user       |

### Example Usage

#### Create a User
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name": "Alice"}' http://localhost:4545/users
```

#### Get All Users
```sh
curl http://localhost:4545/users
```

## Contributing
Feel free to fork and submit PRs!

## License:
`MIT`

This should work for GitHub! Let me know if you need any tweaks. 
