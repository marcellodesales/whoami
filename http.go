package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    versionFile, err := ioutil.ReadFile("/app/version")
    version := "latest"
    if err == nil {
      version = string(versionFile)
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
        w.Header().Set("X-Application-Port", port)
        w.Header().Set("X-Application-Version", version)
        w.Header().Set("X-Application-Env", env)
        w.Header().Set("X-Application-Label", label)
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
 	      fmt.Fprintf(w, "I'm %s\n", hostname)
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}

