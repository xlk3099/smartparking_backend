# Ausome Parking Backend
[![Build Status](https://travis-ci.org/xlk3099/smartparking_backend.svg?branch=master)](https://travis-ci.org/xlk3099/smartparking_backend)

This is used for Autodesk ASRD 2016 Hackathon purpose, a simple server written in Golang

### TechStack:
  - Restful server using Gin framework
  - Websocket communication support by using gorilla/websocket

### To start:
- install and configure `go` on your system
- install 
- in command prompt do
```
go get github.com/xlk3099/smartparking_backend
go get gopkg.in/gin-gonic/gin.v1
go get github.com/gorilla/websocket
```
or Simply copy this repo to your machine. and do 
```
go get gopkg.in/gin-gonic/gin.v1
go get github.com/gorilla/websocket
```

- To launch the server,
```
    go run main.go
```
To play with the websocket communication, 
Modify the client.html file as you like, then open your browser, enter "localhost:8080"
