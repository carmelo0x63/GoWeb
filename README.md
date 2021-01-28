# GoWeb
Simple HTTP responder, written in Go.</br>
</br>
[Automated builds](https://docs.docker.com/docker-hub/builds/) are set up in Docker Hub. To pull/run simply issue the following command:</br>
`docker run -d -p 8080:8080 carmelo0x99/go_web:<tag>`
</br>
**NOTE**: port 8080 must be acessible on the host running the container</br>
- Sample output
```
$ curl http://<ip_address>:8080/Hello
This is 10bd525fa558 running on linux/amd64 saying: Hello
```

To build ARM images, the appropriate Dockerfile shall be explicitly used. For instance:
```
$ docker build -t <repo>/<image>:<tag> -f Dockerfile.armv6 .
```

