/**
If your application is trying to submit requests that exceed the maximum request for your account,
you may receive error messages from Product Advertising API. The request limit for each account is
calculated based on revenue performance. Each account used to access the Product Advertising API
is allowed an initial usage limit of 1 request per second*. Each account will receive an
additional 1 request per second (up to a maximum of 10) for every $4,600 of shipped item
revenue driven in a trailing 30-day period (about $0.11 per minute). You can verify that
your sales are being attributed to your calls to the Product Advertising API by checking
for the following:
*/

package amazon

import (
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
	"app/stores/interface/store"
	"app/libs/confhandler"
	"log"
	"time"
	timer "app/helpers/time"
)

//Json conf struct
type ConfigurationApiAmazon struct {
	Amazon struct {
		ApiEndPoint    string
		ApiRoot        string
		AWSAccessKeyId string
		AWSSecretKeyId string
		AssociateTag   string
		Service        string
		SearchIndex    string
		ResponseGroups string
		Version        string
		Region         string
	}
}

//Parameters struct for amazon
type Store struct {
	Keywords  string
	Operation string
}

func (i Store) GetStore() store.Client {
	//Configurations for stores
	//Pass by reference
	var conf ConfigurationApiAmazon
	confhandler.SetConf(&conf, "stores.json")

	//Conf amazon region..
	region := amazon.Region(conf.Amazon.Region)
	client, err := amazon.New(conf.Amazon.AWSAccessKeyId,
		conf.Amazon.AWSSecretKeyId, conf.Amazon.AssociateTag, region)

	//handle error
	if err != nil {
		log.Fatal(err)
	}

	//Return url ready
	return store.Client{Client: client, Store: "Amazon"}

}

//func (i Store) SearchItem(client amazon.Client) store.Result {
//
//	//Return results
//	//return store.Result{Result: product, Extra: i.Keywords}
//}

func (i Store) GetOperation() func(amazon.Client) store.Result {
	//Function lists
	functions := map[string]func(amazon.Client) store.Result{
		"SearchItem": i.SearchItem,
	}

	//SwitchOperations
	if i.Operation != "" {
		return functions[i.Operation]
	} else {
		return functions["SearchItem"]
	}

}
