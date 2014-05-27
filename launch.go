package main

import (
  "github.com/danhigham/pty"
  "io"
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
  "os"
  "os/signal"
  "os/exec"
  "regexp"
  "net"
)

func main() {

  bin := "tmate"
  cmd := exec.Command(bin)

  log.Print("Starting tmate...")

  f, err := pty.Start(cmd)
  pty.Setsize(f, 1000, 1000)

  row, col, _ := pty.Getsize(f)

  log.Print(col)
  log.Print(row)

  if err != nil {
    log.Print(err)
  }

  quit := make(chan int)

  go func(r io.Reader) {
    sessionRegex, _ := regexp.Compile(`Remote\ssession\:\sssh\s([^\.]+\.tmate.io)`)

    // select {
    //   case <- quit:
    //     log.Print("Stopping tmate")
    //     return
    // }

    for {

      buf := make([]byte, 1024)
      _, err := r.Read(buf[:])

      if err != nil {
        return
      }

      matches := sessionRegex.FindSubmatch(buf)

      if len(matches) > 0 {
        log.Print(string(matches[1]))
      }
    }

  }(f)

  serverUrl, _ := url.Parse("http://127.0.0.1:8080")
  reverseProxy := httputil.NewSingleHostReverseProxy(serverUrl)

  http.Handle("/", reverseProxy)

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)

  l, _ := net.Listen("tcp", ":" + os.Getenv("PORT"))

  go func(){
    for _ = range c {
      // sig is a ^C, handle it
      log.Print("Stopping tmate...")

      close(quit)

      l.Close()
      f.Close()
    }
  }()

  http.Serve(l, nil)
  log.Print(err)
}
