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
	timer "app/helpers/time"
	"app/helpers/products"
	"log"
	"app/helpers/functions"
	"strconv"
	"time"
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
	Keywords   string
	Operation  string
	Parameters map[string]string
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

func (i Store) SearchItem(client store.Client, params ...map[string]string) store.Result {
	/**	The ItemSearch operation searches for items on Amazon. The Product Advertising API
	returns up to ten items per search results page.*/
	var pageIndex int32 = 1
	var product []store.Product = []store.Product{}
	//Configuration for type of result in item
	var storeGroup []amazon.ItemSearchResponseGroup = []amazon.ItemSearchResponseGroup{
		amazon.ItemSearchResponseGroupImages,
		amazon.ItemSearchResponseGroupItemAttributes,
	}

	//N bitSize
	//-(2^N-1) a (2 ^N-1)-1
	page, err := functions.ExtractParam("Page", params)

	//If param exists
	if err == nil {
		pageIndex64, _ := strconv.ParseInt(page, 10, 32)
		pageIndex = int32(pageIndex64)
	}

	//Request Search
	res, err := client.Client.ItemSearch(amazon.ItemSearchParameters{
		SearchIndex:    amazon.SearchIndexAll,
		OnlyAvailable:  true,
		Keywords:       i.Keywords,
		ResponseGroups: storeGroup,
		ItemPage:       int(pageIndex), //Pagination for items results
	}).Do()

	//If error
	if err != nil {
		//Each account used to access the Product Advertising API
		//is allowed an initial usage limit of 1 request per second
		log.Print(err)
		log.Println("Retrying request")
		timer.Wait(time.Second)
		//TODO here need to pass `params` as argument
		//TODO handle switch exceptions to know what to do!!
		return i.SearchItem(client)
	}

	//Standardized structure for product
	for _, item := range res.Items.Item {
		//Price
		amount := item.ItemAttributes.ListPrice.FormattedPrice
		//Handle images
		imageList := products.FetchProductImages(item)
		//Pack product data
		product = append(product, store.Product{
			Id:    item.ASIN,
			Title: item.ItemAttributes.Title,
			Price: amount,
			Image: imageList})
	}

	//Return results
	return store.Result{
		Result: product,
		Extra:  i.Keywords,
		Total:  res.Items.TotalResults,
		Pages:  res.Items.TotalPages}
}

func (i Store) ItemLookUp(client store.Client, params ...map[string]string) store.Result {
	/** Given an Item identifier, the ItemLookup operation returns some or all of the
	item attributes, depending on the response group specified in the request. By default,
	ItemLookup returns an item’s ASIN, Manufacturer, ProductGroup, and Title of the item.*/

	var product []store.Product = []store.Product{}
	//Configuration for type of result in item
	var storeGroup []amazon.ItemLookupResponseGroup = []amazon.ItemLookupResponseGroup{
		amazon.ItemLookupResponseGroupLarge,
	}

	//Request item lookup
	res, err := client.Client.ItemLookup(amazon.ItemLookupParameters{
		IDType:         amazon.IDTypeASIN,
		ResponseGroups: storeGroup,
		ItemIDs: []string{
			i.Keywords,
		},
	}).Do()

	//If error
	if err != nil {
		//Each account used to access the Product Advertising API
		//is allowed an initial usage limit of 1 request per second
		log.Print(err)
		log.Println("Retrying request itemlookup")
		timer.Wait(time.Second)
		return i.ItemLookUp(client)
	}

	//Standardized structure for product
	item := res.Items.Item[len(res.Items.Item)-1] // Pop
	//Price of product
	amount := item.ItemAttributes.ListPrice.FormattedPrice
	//Handle images
	imageList := products.FetchProductImages(item)
	//The product list
	product = append(product, store.Product{
		Id:    item.ASIN,
		Title: item.ItemAttributes.Title,
		Price: amount,
		Image: imageList,
	})

	return store.Result{
		Result: product,
		Extra:  i.Keywords,
		Total:  res.Items.TotalResults,
		Pages:  res.Items.TotalPages}
}

func (i Store) GetOperation() store.Operation {
	//Operation to handle function to execute
	var operation string = i.Operation

	//If not operations
	if len(operation) == 0 {
		operation = "SearchItem"
	}

	//Function lists
	functionList := map[string]store.Operation{
		"SearchItem": i.SearchItem,
		"ItemLookup": i.ItemLookUp,
	}

	//SwitchOperations
	return functionList[operation]
}

func (i Store) GetParameters() map[string]string {
	//Return parameters passed
	return i.Parameters
}
