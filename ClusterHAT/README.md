This directory holds a few files backported from my ClusterHAT.

----

https://docs.docker.com/network/drivers/overlay/

docker network \
  create \
  --driver overlay \
  --attachable \
  goweb_backend

docker stack \
  deploy \	
  --compose-file compose_web.yaml \
  goweb

docker compose \
  --file compose_lb.yaml \
  up -d



docker network inspect --format "{{range .Peers}}{{println .IP}}{{end}}" goweb_backend
192.0.2.82
192.0.2.83
192.0.2.81
192.0.2.84
192.0.2.80


$ docker service inspect --format "{{.Endpoint}}" goweb_web
{{vip [{ tcp 8080 8888 ingress}]} [{ tcp 8080 8888 ingress}] [{z432jvyrusib41qsabqs12cp3 10.0.0.218/24} {mhxqt6qzvp3m5jflunvnhbtq1 10.0.29.2/24}]}


$ docker container inspect --format "{{println .NetworkSettings.Networks.goweb_backend.IPAddress}}{{.NetworkSettings.Ports}}" goweb-lb-1
10.0.29.13
map[80/tcp:[] 9999/tcp:[{0.0.0.0 9090} {:: 9090}]]


pi@zero1:~$ docker container inspect --format "{{println .NetworkSettings.Networks.goweb_backend.IPAddress}}{{.NetworkSettings.Ports}}" goweb_web.1.f4i81ic0jk01on9puumzktsma
10.0.29.3
map[]

