package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coseo12/nomacoin/blockchain"
	"github.com/coseo12/nomacoin/utils"
	"github.com/gorilla/mux"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}
type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type addTxPayload struct {
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

func documentation(w http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "List a block",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "See a block",
		},
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "Get TxOuts for an Address",
		},
		{
			URL:         url("/mempool"),
			Method:      "GET",
			Description: "Get Mempool",
		},
		{
			URL:         url("/transactions"),
			Method:      "POST",
			Description: "Set Transactions",
		},
	}
	json.NewEncoder(w).Encode(data)
}

func blocks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(blockchain.Blockchain().Blocks())
	case "POST":
		blockchain.Blockchain().AddBlock()
		w.WriteHeader(http.StatusCreated)
	}
}

func block(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(w)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.Blockchain())
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func balance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount := blockchain.Blockchain().BalanceByAddress(address)
		json.NewEncoder(w).Encode(balanceResponse{address, amount})
	default:
		utils.HandleErr(json.NewEncoder(w).Encode(blockchain.Blockchain().TxOutsByAddress((address))))
	}
}

func mempool(w http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(w).Encode(blockchain.Mempool.Txs))
}

func transactions(w http.ResponseWriter, r *http.Request) {
	var payload addTxPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		json.NewEncoder(w).Encode(errorResponse{"not enough founds"})
	}
	w.WriteHeader(http.StatusCreated)
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/balance/{address}", balance).Methods("GET")
	router.HandleFunc("/mempool", mempool).Methods("GET")
	router.HandleFunc("/transactions", transactions).Methods("POST")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
