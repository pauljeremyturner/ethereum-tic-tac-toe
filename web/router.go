// swagger:meta
package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var games sync.Map

func Route() {
	r := mux.NewRouter()
	// swagger:operation POST /games/{id} games gameMove
	// ---
	// summary: Start Game or Move in existing Game
	// description: If the game id is found, move in existing game, otherwise create new game.  If new game, caller can either specify empty board (computer turn first) or board with a single cross or circle (user turn first).  Caller provides game id - can be a string uuid preferred.
	// parameters:
	// - name: gameMove
	//   description: Request move in existing game or request for new game
	//   in: body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Game"
	// responses:
	//   "200":
	//     "$ref": "#/responses/gameState"
	r.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {

		rqBytes, err1 := ioutil.ReadAll(r.Body)


		if err1 != nil {
			log.Printf("Failed read post body err: %s", err1)
			panic("Could not read post body")
		}

		move := &gameMove{}

		err2 := json.Unmarshal(rqBytes, &move)

		if err2 != nil {
			log.Printf("Failed unmarshal json post body, error: %s", err2)
			w.WriteHeader(http.StatusBadRequest)
		} else {


			gs := &gameState{
				GameId:       "1234",
				CirclePlayer: "Paul",
				CrossPlayer: "Peeps",
				Circles:      576,
				Crosses:      96,
			}

			rsBytes, rr := json.Marshal(gs)
			if rr != nil {
				panic("ouch")
			}
			w.WriteHeader(http.StatusOK)
			w.Write(rsBytes)
		}

	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}


