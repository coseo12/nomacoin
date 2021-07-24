package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/coseo12/nomacoin/blockchain"
)

const (
	templateDir string = "explorer/templates/"
)

var port string
var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	data := &homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(w, "home", data)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
