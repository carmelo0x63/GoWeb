# GoWeb
A simple HTTP responder, written in Go.

File `goweb.go` contains the code that is to be compiled and run by means of containers.</br>
The source code describes two functions:
1. `main()`: listens and servers any HTTP GET requests on port 8080 (`listenPort`)
2. `sayHello()`: returns a message saying `This is <hostname> running on <OS/arch> saying: <received_message>`

### Local test
Run as `go run goweb.go`. From a different terminal tab, run the following:
```
$ curl http://127.0.0.1:8080/Hello%20there\!
This is <hostname> running on linux/amd64 saying: Hello there!
```

### Build and run as container
- build
```
$ docker buildx build --tag <repository>/<name>:<tag> .
```
**NOTE**: mind the dot ("`.`") at the end of the command above

- run
```
$ docker run \
  --detach \
  --publish published=9090,target=8080 \
  --name goweb \
  <repository>/<name>:<tag>
```

- test
```
$ curl http://<docker_engine_ip_address>:9090/Hello%20there\!
This is <hostname> running on linux/amd64 saying: Hello there!
```

See also: [https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go).
