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
  conn, err := listener.Accept()
  if err != nil {
    fmt.Println("we gotta connection accept error boss")
  }
  fmt.Println("TCP connection accepted")
  return conn
}

func CloseConnectionLoudly(conn net.Conn) {
  conn.Close()
  fmt.Println("Connection closed")
}

func ListenAndHandleTCPShell(){
  fmt.Println("hello from the other thread")
  conn := ListenForConnection()
  defer CloseConnectionLoudly(conn)
  stdreader := bufio.NewReader(os.Stdin)
  reader := bufio.NewReader(conn)
  writer := bufio.NewWriter(conn)
  go func() {
    for{
      tcpdata, err := reader.ReadString('>')
      if err != nil {
        fmt.Println("we gotta tcp error boss")
      }
      fmt.Print(tcpdata)
    }
  }()
  go func(){
    for{
      text, _ := stdreader.ReadString('\n')
      fmt.Println("read text: " + text)
      if text == "yeet" {
        fmt.Println("quitting")
        break
      }
      writer.WriteString(text)
      writer.WriteString("\r\n")
      writer.Flush()
    }
  }()
  for {}
}
