package server
import(
    "net/http"
    "os"
    "log"
    "fmt"
    "html/template"
    "golang.org/x/net/websocket"
)
    var listenAddr string

func StartServing(){
    fmt.Println(os.Args)
    if(len(os.Args) < 3){
        log.Fatal("Pass listenInterface port on command line")
    }
    listenAddr = os.Args[1]
    listenAddr += ":"
    listenAddr += os.Args[2]
    http.HandleFunc("/", handler)
    http.Handle("/socket", websocket.Handler(socketHandler))
    err := http.ListenAndServe(listenAddr, nil)
    if(err != nil){
        log.Fatal(err)
    }
}

func handler(w http.ResponseWriter, r *http.Request){
    rootTemplate.Execute(w, listenAddr)
}
func socketHandler(c *websocket.Conn){
    var s string
    fmt.Fscan(c, &s)
    fmt.Println("Received ", s)
    fmt.Fprint(c, "Hi from Go")
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>Fuzzy Farm</title>
<script>
document.addEventListener('DOMContentLoaded', init)
function init(){
    var WS = new WebSocket('wss://go-plato-platocambrian.c9users.io/socket')
    WS.onmessage = function(msg){
        console.log(msg)
    }
    WS.onopen = function(){
        console.log('JS connected!')
        WS.send('JS connected!');
    }
}
</script>
<body>
<p>Hello Web!</p>
</html>
`))