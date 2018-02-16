whoami
======

Simple HTTP docker service that prints it's container ID. In addition, it can show version and environments as response headers.
This is useful to verify kubernetes deployments with versions and rollouts.

# Server - Build & Run

* Execute [build-and-run.sh](build-and-run.sh)

```
$ ./build-and-run.sh
...
...
Step 14/14 : CMD ["/app/http"]
 ---> Running in c1644aedd2bc
Removing intermediate container c1644aedd2bc
 ---> 42ff86ea56a1

Successfully built 42ff86ea56a1
Successfully tagged marcellodesales/whoami:latest
Recreating whoami_whoami_1 ... done
Attaching to whoami_whoami_1
whoami_1  | WhoAmI Server version=95dd08c label=develop env=prodListening on d5aca8c3ba93:8000
whoami_1  |
```

# Client Requests 

* Just call the service on the published port; review and change it in `docker-compose`.

```
$ curl -i localhost:8000
HTTP/1.1 200 OK
X-Application-Env: prod
X-Application-Label: develop
X-Application-Port: 8000
X-Application-Version: 95dd08c
Date: Thu, 15 Feb 2018 21:54:16 GMT
Content-Length: 71
Content-Type: text/plain; charset=utf-8

I'm d5aca8c3ba93 serving request user-agent: curl/7.54.0, accept: */*
```

* The server-side output will show the current client request info:

```
Successfully built 42ff86ea56a1
Successfully tagged marcellodesales/whoami:latest
Recreating whoami_whoami_1 ... done
Attaching to whoami_whoami_1
whoami_1  | WhoAmI Server version=95dd08c label=develop env=prodListening on d5aca8c3ba93:8000
whoami_1  |
whoami_1  | This is server d5aca8c3ba93:8000, version=95dd08c label=develop env=prod, serving request user-agent: curl/7.54.0, accept: */*
```

* Just to confirm that the version is being retrieved from the local container file.

```
$ docker exec -ti d5aca8c3ba93 cat /app/label
develop

$ docker exec -ti d5aca8c3ba93 cat /app/version
95dd08c
```

# Docker Engine

* Simpler example with defaults.

```
$ docker run -d -p 8000:8000 --name whoami -t marcellodesales/whoami
736ab83847bb12dddd8b09969433f3a02d64d5b0be48f7a5c59a594e3a6a3541
    
$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
I'm 736ab83847bb
```

