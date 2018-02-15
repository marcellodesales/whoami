package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
  "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Set the version as it depends in build
    versionFile, err := ioutil.ReadFile("/app/version")
    version := "latest"
    if err == nil {
      version = strings.Replace(string(versionFile), "\n", "", -1)
    }

    // Set the label as it depends in build
    labelFile, err := ioutil.ReadFile("/app/label")
    label := "master"
    if err == nil {
      label = strings.Replace(string(labelFile), "\n", "", -1)
    }

    // Set the port as it depends on deploy
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Set the env as it depends in deployment
    env := os.Getenv("ENV")
    if env == "" {
        env = "default"
    }

    hostname, _ := os.Hostname()
    fmt.Fprintf(os.Stdout, "WhoAmI Server version=%s label=%s env=%s\n", version, label, env)
    fmt.Fprintf(os.Stdout, "Listening on %s:%s\n\n", hostname, port)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Set the response headers
        w.Header().Set("X-Application-Port", port)
        w.Header().Set("X-Application-Version", version)
        w.Header().Set("X-Application-Env", env)
        w.Header().Set("X-Application-Label", label)

        var rHeaders []string; 
        // Loop through headers
        for name, headers := range r.Header {
            name = strings.ToLower(name)
            for _, h := range headers {
              rHeaders = append(rHeaders, fmt.Sprintf("%v: %v", name, h))
            }
        }

        allHeaders := strings.Join(rHeaders, ", ")

        // Log the server-side
        fmt.Fprintf(os.Stdout, "This is server %s:%s serving request %s\n", hostname, port, allHeaders)
        fmt.Fprintf(w, "I'm %s serving request %s \n", hostname, allHeaders)
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}

