package index
//
import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//Interface Store

	w.Write([]byte("Welcome to eShop Api."))

	//fmt.Print(amazonParams.GetStoreUri())
}
