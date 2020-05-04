package controllers

import (
	"app/workspace/thirdParty"
	"app/workspace/schemas"
	"github.com/googollee/go-socket.io"
	"log"
)

type Room struct {
    Room   		string          `json:"room"`
    SortArray 	[]int   		`json:"sortArray"`
    Image 		string 			`json:"image"`
}

func searchPlayer(playerId string) (string) {
	client := thirdParty.GetRedisClient()
	playerFound, error := client.SPop("players").Result()
	if error != nil {
		log.Printf("Error %s: %s", playerId, error.Error())
		client.SAdd("players" ,playerId)
		return ""
	}
	return playerFound
}

func WaitForPlayer(s socketio.Conn, server *socketio.Server) {
	player := searchPlayer(s.ID())
	if player != "" {
		Join(player, s, server);
	}
}

func Join(player string ,s socketio.Conn, server *socketio.Server) () {
	sortArray := []int{1,2,3,4,5,6,7,8,9}
	image := "https://i.picsum.photos/id/237/500/500.jpg"
	server.ForEach("", "waiting-room", 
        func(socketConnection socketio.Conn) {
            if socketConnection.ID() == s.ID()  {
            	room := map[string]interface{}{}
				room["sortArray"]= sortArray
				room["image"]= image
            	room["opponent"]= player
                socketConnection.Emit("waiting-room", room)
            }
            if socketConnection.ID() == player {
            	room := map[string]interface{}{}
				room["sortArray"]= sortArray
				room["image"]= image
            	room["opponent"]= s.ID()
                socketConnection.Emit("waiting-room", room)
            }
        },
    )
}

func Move(request schemas.MoveRequest, server *socketio.Server) {
	server.ForEach("", "waiting-room", 
        func(socketConnection socketio.Conn) {
            if socketConnection.ID() == request.Opponent {
                socketConnection.Emit("game", request.SortArray)
            }
        },
    )
}