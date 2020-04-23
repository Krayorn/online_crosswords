package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	"encoding/json"
)

type Game struct {
	Grid        []Row `json:"grid"`
}

type Row struct {
	Cells	[]Cell `json:"cells"`
}

type Cell struct {
	Value string `json:"value"`
	Kind string `json:"kind"`
}

var StateGame = Game{}

var socketList []*websocket.Conn

func main() {
	log.Println("listening on 80")

	width, height := 10, 10

	row := make([]Cell, width)
	grid := make([]Row, height)

	for i := 0; i < width; i++ {
		row[i] = Cell{Value:"", Kind:"fillable"}
	}

	for j := 0; j < height; j++ {
		copiedRow := make([]Cell, width)
		copy(copiedRow, row)
		grid[j] = Row{Cells:copiedRow}
	}

	StateGame = Game{
		Grid: grid,
	}

	http.Handle("/", websocket.Handler(HandleClient))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func HandleClient(ws *websocket.Conn) {
	var checkSocket bool
	for _, socket := range socketList {
		if ws == socket {
			checkSocket = true
		}
	}
	if checkSocket == false {
		socketList = append(socketList, ws)
		ws.Write(sendGameState())
	}

	var info map[string]interface{}
	var data string
	for {
		_ = websocket.Message.Receive(ws, &data)
		json.Unmarshal([]byte(data), &info)

		if info["kind"] == "update" {
			rowIndex := int(info["row"].(float64))
			cellIndex := int(info["cell"].(float64))

			StateGame.Grid[rowIndex].Cells[cellIndex].Value = info["value"].(string)

			sendGameUpdate()
		}
	}
}

func sendGameState() []byte {
	message, err := json.Marshal(StateGame)
	if err != nil {
		fmt.Println("Something wrong with JSON Marshal init")
	}
	return message
}

func sendGameUpdate() {
	for _, socket := range socketList {
		socket.Write(sendGameState())
	}
}
