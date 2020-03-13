# GoWeb
Simple HTTP responder, written in Go.</br>
[Automated builds](https://docs.docker.com/docker-hub/builds/) are set up in Docker Hub. To pull/run simply issue the following command:</br>
`docker run -d -p 8080:8080 ccarmelo/goweb:latest`
</br>
**NOTE**: port 8080 must be acessible on the host running the container</br>
- Sample output
```
$ curl http://<ip_address>:8080/Hello
This is 10bd525fa558 running on linux/amd64 saying: Hello
```
