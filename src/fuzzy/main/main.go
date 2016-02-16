package main
import(
    "fuzzy/server"
)

func main(){
    go server.StartServing()
    quit := make(chan bool)
    _ = <-quit
}