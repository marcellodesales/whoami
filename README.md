whoami
======

# Docker Compose

```
$ docker-compose up
Recreating whoami_whoami_1 ... done
Attaching to whoami_whoami_1
whoami_1  | Listening on :8000
Killing whoami_whoami_1 ... done
```

Get the port mapped on the host

```
$ docker ps
CONTAINER ID        IMAGE                                                         
 COMMAND                  CREATED                  STATUS   
 PORTS                     NAMES
0f192f4af208        marcellodesales/whoami                                        
              "/app/http"              Less than a second ago   Up 6 seconds       
             0.0.0.0:32768->8000/tcp   whoami_whoami_1
```

Now that you can see port `32768` was randomly assigned, you can call the service and see all the details:

```
$ curl localhost:32768
I'm 0f192f4af208

$ curl  -i localhost:32768
HTTP/1.1 200 OK
X-Application-Env: prod
X-Application-Label: master
X-Application-Port: 8000
X-Application-Version: 1.0.0
Date: Thu, 15 Feb 2018 18:29:45 GMT
Content-Length: 17
Content-Type: text/plain; charset=utf-8

I'm 0f192f4af208
```

# Docker Engine

Simple HTTP docker service that prints it's container ID

```
$ docker run -d -p 8000:8000 --name whoami -t marcellodesales/whoami
736ab83847bb12dddd8b09969433f3a02d64d5b0be48f7a5c59a594e3a6a3541
    
$ curl $(hostname --all-ip-addresses | awk '{print $1}'):8000
I'm 736ab83847bb
```

