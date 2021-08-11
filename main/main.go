package main

import (
  "fmt"
  "os"
  "time"
  "log"
  "strconv"
  miniredis "github.com/alicebob/miniredis/v2"
)

func forever() {
    for {
        time.Sleep(time.Second)
    }
}

func main() {
  m := miniredis.NewMiniRedis()
  port, _ := strconv.Atoi(os.Args[1])
  err := m.StartAddr(fmt.Sprintf("0.0.0.0:%d", port))

  if err != nil {
    log.Printf("Error %s", err)
    os.Exit(1)
  }

  defer m.Close()

  go forever()
  select {}
}
