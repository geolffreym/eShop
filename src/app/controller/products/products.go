package products

import (
	"net/http"
	"app/stores/amazon"
	"app/stores/interface/store"
	"app/libs/thread"
	"gopkg.in/unrolled/render.v1"
	"app/helpers/products"
	"github.com/gorilla/mux"
	"fmt"
)

func IndexProducts(w http.ResponseWriter, req *http.Request) {
	//Interface Store
	r := render.New()
	q := mux.Vars(req)
	fmt.Print(q)
	//Check if valid param for search
	if val, ok := q["q"]; ok {
		//Stores
		amazonParams := amazon.Store{}
		amazonParams.Keywords = val
		amazonParams.Parameters = map[string]string{}

		//Params
		//If param page
		if r, oki := q["page"]; oki {
			amazonParams.Parameters["Page"] = r
		}

		//Slice of stores
		//amazonParams.Operation = "SearchItem"
		storeList := []store.IStore{amazonParams}

		//Handle channels response
		channel := thread.NewChannel(storeList)
		threadsResponse := thread.ProcessChannelResponse(channel)

		//Parallelism and Concurrence
		//Get merged results
		resultMap := products.MergeProducts(threadsResponse)

		//Response with json
		r.JSON(w, http.StatusOK, resultMap)
		return
	}

	//404
	http.Error(w, "No products were found", 404)

}

func IndexProductsByID(w http.ResponseWriter, req *http.Request) {
	//Interface Store
	r := render.New()
	vars := mux.Vars(req)
	//Check if valid param
	if val, ok := vars["id"]; ok {

		//Stores
		amazonParams := amazon.Store{}
		//"B0719R4WYK"
		amazonParams.Keywords = val
		amazonParams.Operation = "ItemLookup"

		//amazonParams.Operation = "SearchItem"
		storeList := []store.IStore{amazonParams}

		//Handle channels response
		channel := thread.NewChannel(storeList)
		//Return and array of responses
		threadsResponse := thread.ProcessChannelResponse(channel)

		//Get merged results from multi-responses
		resultMap := products.MergeProducts(threadsResponse)
		//Response with json
		r.JSON(w, http.StatusOK, resultMap)
		return
	}

	//404
	http.Error(w, "Invalid product id", 404)

}
