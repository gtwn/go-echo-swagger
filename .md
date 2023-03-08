## Define structure
```
├── app
│   └── api
│       ├── docs
│       │   ├── docs.go
│       │   ├── swagger.json
│       │   └── swagger.yaml
│       ├── main.go
│       ├── route
│       │   └── check.go
│       └── server
│           └── server.go
├── go.mod
└── go.sum
```
---

## Declarative comments format follow
https://github.com/swaggo/swag#declarative-comments-format

---

- `swag init --md ./` at main root
- `go run app/api/*.go`