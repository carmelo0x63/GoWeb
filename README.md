# GoWeb
A simple HTTP responder, written in Go.

File `goweb.go` contains the code that is to be compiled and run by means of containers.</br>
The source code describes two functions:
1. `main()`: listens and servers any HTTP GET requests on port 8888 (`listenPort`)
2. `sayHello()`: returns a message saying `This is <hostname> running on <OS/arch> saying: <received_message>`

### Local test
<img src="assets/images/stand-alone.png">

Run as `go run goweb.go`. From a different terminal tab, run the following:
```
$ curl http://127.0.0.1:8888/Hello%20GoWeb\!
This is <hostname> running on linux/amd64 saying: Hello GoWeb!
```

----

### Build and run as container
<img src="assets/images/docker-basic.png">

- build
```
$ docker buildx build --tag <repository>/<name>:<tag> .
```
**NOTE**: mind the dot ("`.`") at the end of the command above

- run
```
$ docker run \
  --detach \
  --publish published=8080,target=8888 \
  --name responder \
  <repository>/<name>:<tag>
```

- test
```
$ curl http://<private_subnet_ip_address>:8080/Hello%20GoWeb\!
This is <hostname> running on linux/amd64 saying: Hello GoWeb!
```
**NOTE**: the published port (`8080`) is different from the listening port (`8888`)

### Run with Docker Swarm
<img src="assets/images/docker-swarm.png">

- run
```
$ docker service create \
  --replicas 3 \
  --name responder \
  --publish published=8080,target=8888 \
  <repository>/<name>:<tag>
```

- test
```
$ curl http://<private_subnet_ip_address>:8080/Hello%20GoWeb\!
This is <hostname> running on linux/amd64 saying: Hello GoWeb!
```

#### To scale up/down
```
$ docker service scale goweb=5
```
**NOTE**: the numbers `3` and `5` above are only examples

----

### Run as a `Stack`
```
$ docker stack deploy \
  --compose-file compose.yaml
 responder 
```
**NOTE**: what was before stated through the command line is now described in the `compose.yaml` file

----

### Add a load balancer (`nginx`)
<img src="assets/images/docker+lb.png">

Take a look and edit as appropriate file `data/loadbalancer/default.conf.ori`. The file defines:</br>
- `backend`: the servers (private subnet IP addresses) and ports on which the service stack is offered
- `listen`: the port on which the LB operates (within the container)
</br>

- run
```
docker run \
  --detach \
  --name loadbalancer \
  --mount type=bind,source=./data/loadbalancer,target=/etc/nginx/conf.d \
  --publish published=9090,target=9999 \
  nginx
```

- test
```
$ curl http://<private_subnet_ip_address>:9090/Hello%20GoWeb\!
This is <hostname> running on linux/amd64 saying: Hello GoWeb!
```

----

See also: [https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go)
