package models
import(
    "golang.org/x/net/websocket"
)
// thanks https://talks.golang.org/2012/chat.slide#32
type FarmSocket struct {
    Conn *websocket.Conn
    Done chan bool
}
func (s FarmSocket) Close() error {
    s.Done <- true
    return nil
}