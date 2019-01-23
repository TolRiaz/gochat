package main


func main() {
    tcpAddr, err := net.ResolveTCPAddr("tcp", "www.google.co.kr:ssh")
    if err != nil {
        fmt.Println("Error : ", err.Error())
        os.Exit(1)
    }
    fmt.Println("IP ", tcpAddr.IP.String(),
        "Port ", tcpAddr.Port)
    os.Exit(0)
}
