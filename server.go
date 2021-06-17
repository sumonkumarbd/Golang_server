//go 1.10.4

package main
import ("fmt"
        "net"
        "os"
        )







func main() {
  
  
  nl, err :=net.Listen("tcp","localhost:8888")//ip:port
  if err != nil{
    fmt.Println(err.Error())
    os.Exit(1) //1
  }
 
  
  conn, err := nl.Accept()
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1) //1
  }
  
  romoteAddr := conn.RemoteAddr().String()
  fmt.Println(romoteAddr)
  
  //byte
  conn.Write([]byte("Welcome I'm Sumon kumar"))
  
  conn.Close()
  nl.Close()
  
  
}