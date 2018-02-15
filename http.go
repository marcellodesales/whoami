package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    version := os.Getenv("VERSION")
    if version == "" {
        version = "latest"
    }
  
    env := os.Getenv("ENV")
    if env == "" {
        env = "default"
    }

    label := os.Getenv("LABEL")
    if label == "" {
        label = "master"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        rw.Header().Set("X-Application-Port", port)
        rw.Header().Set("X-Application-Version", version)
        rw.Header().Set("X-Application-Env", env)
        rw.Header().Set("X-Application-Label", label)
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
 	      fmt.Fprintf(w, "I'm %s\n", hostname)
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}

