package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"google.golang.org/appengine"

	"github.com/darthfrazier/vanity/internal/pkg/vanity"
)

func main() {
	registerHandlers()
	appengine.Main()
}

func registerHandlers() {
	r := mux.NewRouter()

	// Basic CRUD endpoints
	r.Methods("POST").Path("/piece").Handler(appHandler(createPieceHandler))
	r.Methods("POST").Path("/outfit").Handler(appHandler(createOutfitHandler))

	r.Methods("GET").Path("/piece/{id:[0-9]+}").Handler(appHandler(readPieceHandler))
	r.Methods("GET").Path("/outfit/{id:[0-9]+}").Handler(appHandler(readOutfitHandler))
	// TODO:: implement
	// r.Methods("GET").Path("/pieces").Handler(appHandler(readPiecesHandler))
	// r.Methods("GET").Path("/outfit").Handler(appHandler(readOutfitsHandler))

	r.Methods("PUT").Path("/piece/{id:[0-9]+}").Handler(appHandler(updatePieceHandler))
	r.Methods("PUT").Path("/outfit/{id:[0-9]+}").Handler(appHandler(updateOutfitHandler))

	r.Methods("DELETE").Path("/piece/{id:[0-9]+}").Handler(appHandler(deletePieceHandler))
	r.Methods("DELETE").Path("/outfit/{id:[0-9]+}").Handler(appHandler(deleteOutfitHandler))

	// Application logic endpoints
	r.Methods("PUT").Path("/piece/checkin/{id:[0-9]+}").Queries("checkinSource", "{checkinSource}").
		Handler(appHandler(checkInHandler))
	// r.Methods("GET").Path("/dress-me").Handler(appHandler(dressMeHandler))

	// Respond to App Engine and Compute Engine health checks.
	// Indicate the server is healthy.
	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

func checkInHandler(w http.ResponseWriter, r *http.Request) *appError {
	checkinSource := mux.Vars(r)["checkinSource"]
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "bad piece id: %v", err)
	}
	piece, err := vanity.DB.GetPiece(id)
	if err != nil {
		return appErrorf(err, "could not find piece: %v", err)
	}
	fmt.Println(checkinSource)
	fmt.Println(id)
	piece.ID = id

	// Perform state change if necessary
	err = piece.CheckIn(checkinSource, vanity.DB)
	if err != nil {
		return appErrorf(err, "Error updating piece state: %v", err)
	}

	return nil
}

// Handlers
// func dressMeHandler(w http.ResponseWriter, r *http.Request) *appError {

// }

func createPieceHandler(w http.ResponseWriter, r *http.Request) *appError {
	// Create piece
	var piece vanity.Piece
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&piece)
	if err != nil {
		return appErrorf(err, "%v", err)
	}

	id, err := vanity.DB.AddPiece(&piece)
	if err != nil {
		return appErrorf(err, "could not save piece: %v", err)
	}
	json.NewEncoder(w).Encode(id)
	return nil
}

func readPieceHandler(w http.ResponseWriter, r *http.Request) *appError {
	// Read piece
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "bad piece id: %v", err)
	}
	piece, err := vanity.DB.GetPiece(id)
	if err != nil {
		return appErrorf(err, "could not find piece: %v", err)
	}
	json.NewEncoder(w).Encode(piece)
	return nil
}

// func readPiecesHandler(w http.ResponseWriter, r *http.Request) *appError {
// 	// Read piece
// }

func updatePieceHandler(w http.ResponseWriter, r *http.Request) *appError {
	var piece vanity.Piece
	decoder := json.NewDecoder(r.Body)

	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "bad piece id: %v", err)
	}

	err = decoder.Decode(&piece)
	if err != nil {
		return appErrorf(err, "could not parse piece from json: %v", err)
	}
	piece.ID = id

	err = vanity.DB.UpdatePiece(&piece)
	if err != nil {
		return appErrorf(err, "could not save piece: %v", err)
	}
	return nil
}

func deletePieceHandler(w http.ResponseWriter, r *http.Request) *appError {
	// Delete piece
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "bad piece id: %v", err)
	}

	err = vanity.DB.DeletePiece(id)
	if err != nil {
		return appErrorf(err, "could not delete piece: %v", err)
	}
	return nil
}

func createOutfitHandler(w http.ResponseWriter, r *http.Request) *appError {
	// Create Outfit
	var outfit *vanity.Outfit
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(outfit)
	if err != nil {
		return appErrorf(err, "%v", err)
	}

	outfit.InitializeOutfit(vanity.DB)
	id, err := vanity.DB.AddOutfit(outfit)
	if err != nil {
		return appErrorf(err, "could not save outfit: %v", err)
	}
	json.NewEncoder(w).Encode(id)
	return nil
}

func readOutfitHandler(w http.ResponseWriter, r *http.Request) *appError {
	// Read Outfit
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "bad outfit id: %v", err)
	}
	outfit, err := vanity.DB.GetOutfit(id)
	if err != nil {
		return appErrorf(err, "could not find outfit: %v", err)
	}
	json.NewEncoder(w).Encode(outfit)
	return nil
}

// func readOutfitsHandler(w http.ResponseWriter, r *http.Request) *appError {
// 	// TODO:: read state and class as query params
// 	outfits, err := vanity.DB.GetOutfits("", "Blue")
// 	if err != nil {
// 		return appErrorf(err, "could not find outfit: %v", err)
// 	}
// 	json.NewEncoder(w).Encode(outfits)
// 	return nil
// }

func updateOutfitHandler(w http.ResponseWriter, r *http.Request) *appError {
	// Update Outfit
	decoder := json.NewDecoder(r.Body)

	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "bad outfit id: %v", err)
	}

	outfit, err := vanity.DB.GetOutfit(id)
	if err != nil {
		return appErrorf(err, "could not find outfit: %v", err)
	}

	err = decoder.Decode(outfit)
	if err != nil {
		return appErrorf(err, "could not parse outfit from json: %v", err)
	}

	outfit.InitializeOutfit(vanity.DB)
	err = vanity.DB.UpdateOutfit(outfit)
	if err != nil {
		return appErrorf(err, "could not save outfit: %v", err)
	}
	return nil
}

func deleteOutfitHandler(w http.ResponseWriter, r *http.Request) *appError {
	// Delete Outfit
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "bad outfit id: %v", err)
	}

	err = vanity.DB.DeleteOutfit(id)
	if err != nil {
		return appErrorf(err, "could not delete outfit: %v", err)
	}
	return nil
}

// Error handling
type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}
