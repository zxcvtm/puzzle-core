package routes

import (
	"fmt"
	"log"
    "net/http"
    "app/workspace/schemas"
    "app/workspace/controllers"
    "github.com/googollee/go-socket.io"
    "github.com/gorilla/mux"
)

func SocketApi(r *mux.Router) {
    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }

    server.OnConnect("/", func(s socketio.Conn) error {
        s.SetContext("")
        fmt.Println("connected:", s.ID())
        return nil
    })

    server.OnEvent("/", "join", func(s socketio.Conn, request schemas.JoinRequest) {
        s.Join("waiting-room")
        fmt.Println("----> Join waiting room")
        controllers.WaitForPlayer(s, server)
        response := schemas.JoinResponse{Id: s.ID(), Msg: "connected into waiting room user " + s.ID()}
        s.Emit("waiting-room", response)
    })

    server.OnEvent("/", "game", func(s socketio.Conn, request schemas.MoveRequest) {
        fmt.Println("----> Game Event")
        controllers.Move(request, server)
    })
    
    server.OnError("/", func(s socketio.Conn, e error) {
        fmt.Println("meet error:", e)
    })

    server.OnDisconnect("/", func(s socketio.Conn, reason string) {
        fmt.Println("closed", s.ID(), reason)
    })

    
    go server.Serve()
    //defer server.Close()

    r.Handle("/socket.io/", server)
    r.Handle("/", http.FileServer(http.Dir("./asset")))
}
