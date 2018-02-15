whoami
======

# Docker Compose

```
$ GIT_SHA=$(git rev-parse --short HEAD) docker-compose build && docker-compose up
Building whoami
Step 1/12 : FROM golang:alpine3.6 AS binary
 ---> a8f345e387a3
Step 2/12 : ADD . /app
 ---> Using cache
 ---> d249d9b5d908
Step 3/12 : WORKDIR /app
 ---> Using cache
 ---> 1fbd6030a0ac
Step 4/12 : RUN go build -o http
 ---> Using cache
 ---> c9c7d8cb8c3a

Step 5/12 : FROM alpine:3.6
 ---> 77144d8c6bdc
Step 6/12 : WORKDIR /app
 ---> Using cache
 ---> cfb43d6b41f2
Step 7/12 : ENV PORT 8000
 ---> Using cache
 ---> 8bc16fbf48d9
Step 8/12 : EXPOSE 8000
 ---> Using cache
 ---> d5e4e0ea8cc3
Step 9/12 : COPY --from=binary /app/http /app
 ---> Using cache
 ---> 45e975493933
Step 10/12 : ARG GIT_SHA
 ---> Using cache
 ---> d92e51d8f14d
Step 11/12 : RUN echo $GIT_SHA > /app/version
 ---> Running in 8b8c18834d28
Removing intermediate container 8b8c18834d28
 ---> 2dc962d97072
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

Calling the server with the port number exposed by `docker-compose`.

```
$ curl -i localhost:8000
HTTP/1.1 200 OK
X-Application-Env: prod
X-Application-Label: master
X-Application-Port: 8000
X-Application-Version: df62639
Date: Thu, 15 Feb 2018 19:05:12 GMT
Content-Length: 17
Content-Type: text/plain; charset=utf-8

I'm 03e98b84f714
```

Just to confirm that the version is being retrieved from the local container file.

```
~/dev/github/public/marcellodesales/whoami on  master! ⌚ 11:04:50
$ docker exec -ti 03e98b84f714 cat /app/version
df62639
```

# Docker Engine

Simple HTTP docker service that prints it's container ID

```
$ docker run -d -p 8000:8000 --name whoami -t marcellodesales/whoami
736ab83847bb12dddd8b09969433f3a02d64d5b0be48f7a5c59a594e3a6a3541
    
$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
I'm 736ab83847bb
```

