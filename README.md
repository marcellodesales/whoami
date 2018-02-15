whoami
======

Simple HTTP docker service that prints it's container ID. In addition, it can show version and environments as response headers.

# Docker Compose

* Build and Run: `GIT_SHA=$(git rev-parse --short HEAD) docker-compose build && docker-compose up`

```
$ GIT_SHA=$(git rev-parse --short HEAD) docker-compose build && docker-compose up
Building whoami
Step 1/12 : FROM golang:alpine3.6 AS binary
...
...
Step 12/12 : CMD ["/app/http"]
 ---> Running in fde1cfcb5cfc
Removing intermediate container fde1cfcb5cfc
 ---> 7fbf711e874d

Successfully built 7fbf711e874d
Successfully tagged marcellodesales/whoami:latest
Recreating whoami_whoami_1 ... done
Attaching to whoami_whoami_1
whoami_1  | Listening on :8000
```

* Just call the service on the published port; review and change it in `docker-compose`.

```
$ curl -i localhost:8000
HTTP/1.1 200 OK
X-Application-Env: prod
X-Application-Label: master
X-Application-Port: 8000
X-Application-Version: de4f95e
Date: Thu, 15 Feb 2018 19:07:37 GMT
Content-Length: 17
Content-Type: text/plain; charset=utf-8

I'm 3b018e0f4239
```

* Just to confirm that the version is being retrieved from the local container file.

```
$ docker exec -ti 3b018e0f4239 cat /app/version
de4f95e
```

# Docker Engine


```
$ docker run -d -p 8000:8000 --name whoami -t marcellodesales/whoami
736ab83847bb12dddd8b09969433f3a02d64d5b0be48f7a5c59a594e3a6a3541
    
$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
I'm 736ab83847bb
```

