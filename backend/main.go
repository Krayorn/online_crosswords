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

var StateGame = Game{
	Grid: []Row{
		Row{Cells: []Cell{
				Cell{Value:"rb: Force Mâle br: Recouvrir de glace", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rb: Déclaration de rupture bb: Est vexant", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rb: Rend brun bb: Il sert à lier", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rb: Candide, Ingénu", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rb: Division de varois bb: Sali de nouveau", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"bb: Sérieux, Calme", Kind:"definition"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rb: Perroquets bb: Décelai, Découvris", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"rr: Réexaminera br: Attacha", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Navires du moyen âge bb: Caverne, Repaire", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"bb: Agent de Maîtrise", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"rr: Hissait, Dressait br: Titane de labo", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Point gagnant au tennis bb: Règlement", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Un intellectuel bb: Mit de l'encre", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"bb: Marches derrière", Kind:"definition"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"rr: Essai, Expérience br: Monnaie du Japon", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Angles bb: Telle une ellipse", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Cuit à feu vif bb S'éternise", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Pronom d'intiles bb: Comme, Pareil à", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"br: Enduit de cirage", Kind:"definition"},
				Cell{Value:"rr: Perça bb: Encourager", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Le fait de trier bb: Réaction à chaud", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Faciles bb: Adorai", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"rr: Du nez br: Le petit coin (les)", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Appareil volant bb: Stimulée", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"bb: Courte annéee lumière", Kind:"definition"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Pas trés chaud bb:Café décaféiné", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Eau du pas-de-calais bb: Abri de sioux", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"rr: La même chose br: Saison de vacances", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Marteau, Tournevis bb: Du tonus (du)", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Connut, Sut bb: Voie rapide", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"bb: Matière d'alliance", Kind:"definition"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"rr: A dit en criant (s'est) br: Etat d'Asie", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Un professionnel (Un)", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
		Row{Cells: []Cell{
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"rr: Crier tel un poussin", Kind:"definition"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
				Cell{Value:"", Kind:"fillable"},
			},
		},
	},
}

var socketList []*websocket.Conn

func main() {
	log.Println("listening on 80")

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
