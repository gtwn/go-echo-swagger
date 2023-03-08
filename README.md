## Define structure
```
.
├── README.md
├── app
│   └── api
│       ├── main.go
│       ├── route
│       │   └── check.go
│       └── server
│           └── server.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── pkg
    └── route
        └── ping.go
```
---

## Declarative comments format follow
https://github.com/swaggo/swag#declarative-comments-format

---

- `swag init -g app/api/main.go` at root this command will find all of swag comment in project
- `go run app/api/*.go`