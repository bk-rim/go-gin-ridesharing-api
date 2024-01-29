# Golang ridesharing restful api

## Tech Stack
- golang
- gin 
- postgres
- testify

## Project structure

```md
├── controller
│   ├── allUser.go
│   ├── allUser_test.go
│   ├── authentication.go
│   ├── authentication_test.go
│   ├── travel.go
│   └── travel_test.go
├── database
│   └── config.go
├── go.mod
├── go.sum
├── middleware
│   └── jwtAuth.go
├── model
│   ├── travel.go
│   └── user.go
├── README.md
├── server.go
└── utils
    ├── jwt.go
    └── loadEnv.go
```

## endpoints

- /api/v1/auth (publics routes)
    - /register
    - /login

- /api/v1/admin (privates routes)
    - /users (get all users)

- /api/v1 (privates routes)
    - /travels (create travel)
    - /travels/idDriver/:idDriver (get travels by idDriver)
    - /travels/idTravel/:idTravel (delete travel)

- /api/v1 (publics routes)
    - /travels (get all travels)
    - /travels/start/:start/end/:end (filter by departure and arrival)
    - /travels/start/:start/end/:end/date/:date (filter by departure arrival and date)