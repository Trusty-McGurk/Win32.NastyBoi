package main

import (
  "fmt"
  "net"
  "bufio"
  //"time"
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

func ListenAndHandleTCPShell(stdreader *bufio.Reader){
  fmt.Println("hello from the other thread")
  conn := ListenForConnection()
  scanner := bufio.NewScanner(conn)
  writer := bufio.NewWriter(conn)
  defer fmt.Println("handle thead died")
  go func() {
    defer fmt.Println("reader thread died")
    for scanner.Scan() {
      fmt.Println("scanned data")
      tcpdata := scanner.Text()
      fmt.Println(tcpdata)
    }
  }()
  for{
    fmt.Println("Reading text: ")
    text, _ := stdreader.ReadString('\n')
    fmt.Println("read text: " + text)
    if text == "yeet" {
      fmt.Println("quitting")
      break
    }
    fmt.Println("writing text")
    _, err := writer.WriteString(text)
    writer.WriteString("\r\n")
    if err != nil {
      fmt.Println("write error: " + err.Error())
    }
    fmt.Println("Flushing")
    writer.Flush()
    fmt.Println("wrote text")
  }
  CloseConnectionLoudly(conn)
}
