package main

import (
  "fmt"
  "net"
  "bufio"
  "time"
  //"os"
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
  go func() {
    for scanner.Scan() {
      fmt.Println("scanned data")
      tcpdata := scanner.Text()
      fmt.Println(tcpdata)
    }
  }()
  for{
   // text, _ := stdreader.ReadString('\n')
   // fmt.Println("read text: " + text)
   // if text == "yeet" {
   //   fmt.Println("quitting")
   //   break
   // }
    fmt.Println("we got here")
    time.Sleep(3000 * time.Millisecond)
    _, err := writer.WriteString("systeminfo\n")
    if err != nil {
      fmt.Println("write error: " + err.Error())
    }
    writer.Flush()
    fmt.Println("wrote text")
  }
  CloseConnectionLoudly(conn)
}
