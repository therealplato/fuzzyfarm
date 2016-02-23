package server
import(
    "net/http"
    "os"
    "log"
    "fmt"
    "html/template"
    "golang.org/x/net/websocket"
    "fuzzy/manager"
    "fuzzy/models"
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
func socketHandler(ws *websocket.Conn){
    fmt.Println("Received WS connection")
    s := models.FarmSocket{ws, make(chan bool)}
    go manager.NewManager(s).StartFarm()
    <-s.Done
    /*
    var s string
    fmt.Fscan(c, &s)
    fmt.Println("Received ", s)
    fmt.Fprint(c, "Hi from Go")
    */
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>Fuzzy Farm</title>
<script>
document.addEventListener('DOMContentLoaded', init)
function init(){
    var F = document.getElementById("countFuzzies")
    var A1 = document.getElementById("countKittens")
    var A2 = document.getElementById("countPuppies")
    
    var urlRX1 = /https?:\/\/(.*)/   // capture host, port, path, qs
    var urlRX2 = /\/$/;            // trailing slash
    var socketURL = window.location.href.match(urlRX1)[1];  // first capture group
    var socketURL =  socketURL.replace(urlRX2, '');         // strip trailing slash if any
    var socketURL = "wss://" + socketURL + "/socket"
    var WS = new WebSocket(socketURL);
    
    WS.onmessage = function(msg){
        var d = JSON.parse(msg.data);
        if(d.Err){
            return alert(d.Err)
        }
        F.value = d.Fuzzies.toFixed(2);
        A1.value = d.Cats;
        A2.value = d.Dogs;
    }
    WS.onopen = function(){
        console.log('JS connected!')
    }
    document.getElementById('buyCats').addEventListener('click', function(){
        WS.send("cat");
    })
    document.getElementById('buyDogs').addEventListener('click', function(){
        WS.send("dog");
    })
}
</script>
<style>
body {
    margin: 4em;
}
.leftCol {
    display: inline-block;
    width: 6em;
}
.rightCol {
    display: inline-block;
    margin-left: 6em;
    width: 10em;
}
.animal {
    margin-bottom: 0.em;
}
</style>
</head>
<body>
Spend Fuzzies to buy animals which earn Fuzzies and make more animals!
<div class="animal">
<label class="leftCol">Fuzzies</label>
<input class="rightCol" id="countFuzzies">
</div>
<div class="animal">
<label class="leftCol">Kittens</label>
<input class="rightCol" id="countKittens">
<button id="buyCats">Buy</button>
</div>
<div class="animal">
<label class="leftCol">Puppies</label>
<input class="rightCol" id="countPuppies">
<button id="buyDogs">Buy</button>
</div>
</html>
`))