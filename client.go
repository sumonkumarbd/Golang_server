package main

import ("fmt"  
        "net"
)


func main() {
  conn, err := net.Dial("tcp", "91.205.173.170:8888")
  if err != nil {
    fmt.Println(err.Error())
  }
  
 // msg :="many many thanks to you sir"
  conn.Write([]byte("msg"))
  
  bs := make([]byte, 1024)
  n, _:= conn.Read(bs)
  fmt.Println(string(bs[:n]))
  
  conn.Close()
  
  
}