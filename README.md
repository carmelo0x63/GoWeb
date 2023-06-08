# GoWeb
A simple HTTP responder, written in Go. For a detailed description read [https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go).</br>
</br>
[Automated builds](https://docs.docker.com/docker-hub/builds/) are set up in Docker Hub for AMD64/x86_64 architecture. To pull/run simply issue the following command:</br>
`docker run -d -p 8080:8080 carmelo0x99/goweb:latest`
</br>
Sample output:
```
$ curl http://<ip_address>:8080/Hello\!
This is 10bd525fa558 running on linux/amd64 saying: Hello!
```
**NOTE**: port 8080 must be acessible on the host running the container</br>
</br>
To build ARM images, the appropriate Dockerfile shall be explicitly used. For instance:
```
$ docker build -t carmelo0x99/goweb:armv6 -f Dockerfile.armv6 .
```

