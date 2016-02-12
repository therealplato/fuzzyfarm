package models
import(
    //"golang.org/x/net/websocket"
    "io"
)
// thanks https://talks.golang.org/2012/chat.slide#32
type FarmSocket struct {
    io.ReadWriter
    Done chan bool
}
func (s FarmSocket) Close() error {
    s.Done <- true
    return nil
}