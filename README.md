# E-Commerce Microservices with Golang

This is a project just for hobby

## UI

All UI components will be here and served by Fiber. So, you need to run it using this command

### Running

```
go run .\ui\cmd\web\main.go
```

## Broker Service

This is just a basic broker service for now. There are two endpoints, one is optional.

### Registered Routes

- GET: `/` -> Returns JSON for index
- GET: `/ping` -> Returns 200 and text/plain result. This route works when you use `HeartBeat` middleware.

### Routes

Routes can be found under the `routes` folder. `routes.go` file contains middleware configs and handlers. Routes use these handlers.

### Middleware

All middleware can be found under the `middleware` folder. I created a middleware to demonstrate how you can create your own. Shortly, this middleware adds an ability to show service status. It would be necessary If you use a health check service and need to know your service's status.

### Running

```
go run .\broker-service\cmd\web\main.go
```