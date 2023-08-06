package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	s "github.com/Slvr-one/bookmaker/app/structs"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

// fetch horses available (number)
func GetHorses(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	// return the list of horses
	json.NewEncoder(ctx.Writer).Encode(Horses)
	// fmt.Fprintf(ctx.Writer, "%v", horses)

	// hLen := len(horses)
	// json.NewEncoder(ctx.Writer).Encode(hLen)
}

func CreateHorse(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(ctx.Request)
	horseName := params["name"]

	exist := IfHorseExist(horseName, Horses)

	if !exist {
		msg := fmt.Sprintf("Cannot create, Horse named %s does not exist", horseName)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

}

// func DeleteHorse(ctx *gin.Context) {
// 	ctx.Writer.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(ctx.Request)
// 	horseName := params["name"]

// 	//querry horses available from db:
// 	// horses := db.GetAvHorses()
// 	horses, exist := ctx.Get("horses") //?
// 	// exist := IfHorseExist(horseName)

// 	if !exist {
// 		msg := fmt.Sprintf("Cannot delete, Horse named %s does not exist", horseName)
// 		fmt.Fprint(ctx.Writer, msg)
// 		LogToFile(msg)
// 		return
// 	}

// 	for i, horse := range Horses {
// 		if horse.Name == horseName {
// 			/* append() adds elements to a slice, but in this case, its used to remove an element:
// 			The [:i] notation denotes that all the elements before the i-th index are included,
// 			while [i+1:] denotes that all the elements after the i-th index are included.
// 			By combining these two notations with the ellipsis(...), the horses slice is effectively modified to exclude the i-th element.
// 			note that the original slice is not modified; rather, a new slice is created missing the i-th element. */
// 			Horses = append(Horses[:i], Horses[i+1:]...)
// 			break
// 		}
// 	}
// 	ctx.Writer.WriteHeader(http.StatusNoContent)
// }

// fetch horse by name
func GetHorse(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(ctx.Request)
	horseName := params["name"]

	// check if the horse exists
	exist := IfHorseExist(horseName, Horses)

	if !exist {
		msg := fmt.Sprintf("Cannot get, Horse named %s does not exist", horseName)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

	for _, h := range Horses {

		if h.Name == horseName {
			// var horse Horse
			json.NewEncoder(ctx.Writer).Encode(h)
			return
		}
	}
	// json.NewEncoder(ctx.Writer).Encode(&Horse{})

	// // return the horse
	// fmt.Fprintf(ctx.Writer, "%v", horseName)
	// return
}

// Update horse name
func UpdateHorse(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(ctx.Request)
	horseName := params["name"]
	horseColor := params["color"]
	horseAge, _ := strconv.Atoi(params["age"])

	// check if the horse exists
	exist := IfHorseExist(horseName, Horses)

	if !exist {
		msg := fmt.Sprintf("Cannot update, Horse named %s does not exist", horseName)
		fmt.Fprint(ctx.Writer, msg)
		LogToFile(msg)
		return
	}

	newHorse := s.Horse{Name: horseName, Color: horseColor, Age: horseAge}

	// newHorse := Horse{color = horseColor, name = horseName}

	// get the body of the request
	// body, err := ioutil.ReadAll(r.Body)
	_, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		err := json.NewDecoder(ctx.Request.Body).Decode(&newHorse)
		if err != nil {
			msg := fmt.Sprintf("Error decoding / reading json body - %v", err)
			fmt.Fprint(ctx.Writer, msg)
			Check(err, msg)
			return
		}

		// check if the new horse info is valid
		if newHorse.Name == "" {
			fmt.Fprintf(ctx.Writer, "Horse name cannot be empty")
			return
		}
		if newHorse.Color == "" {
			fmt.Fprintf(ctx.Writer, "Horse color cannot be empty")
			return
		}

		// TODO here probably insert new horse to db
		// update the horse info

		for _, h := range Horses {
			if h.Name == horseName {
				h.Name = horseName   //newHorse.Name
				h.Color = horseColor //newHorse.Color
				h.Age = horseAge     //newHorse.Age
				h.Record.Losses = 0
				h.Record.Wins = 0

				// h.Record = &s.Record{0, 0}

				break
			}
		}
		// return the updated horse

		fmt.Fprintf(ctx.Writer, "%v", Horses)
		return
	}

	// for i, item := range horses {
	// 	if item.Name == horseName {
	// 		horses = append(horses[:i], horses[i+1:]...) //redundent
	// 		var horse Horse
	// 		_ = json.NewDecoder(r.Body).Decode(&horse)
	// 		horse.Name = horseName
	// 		horses = append(horses, horse)
	// 		json.NewEncoder(ctx.Writer).Encode(horse)
	// 		return
	// 	}
	// }
}
