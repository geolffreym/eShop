package products

import (
	"app/stores/interface/store"
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
	"time"
	"app/stores/interface/response"
)

func MergeProducts(threadsResponse []*stres.RThread) stres.Response {
	/**This functions get all products from stores an merge it in one only structure*/
	resultResponse := stres.Response{}
	//The threads results
	resultResponse.Results = []store.Product{}
	resultsTotal := map[string]stres.RDetails{}

	if len(threadsResponse) == 0 {
		return resultResponse
	}

	//Handle responses
	for _, response := range threadsResponse {
		//Add results for each store pages and total
		resultsTotal[response.Store] = stres.RDetails{
			Total:     response.TotalResults,
			Pages:     response.TotalPages,
			Timestamp: time.Now().Unix()}

		//Append detail to response
		resultResponse.Data = resultsTotal

		//Add store to item
		for _, item := range response.Items {
			//Append store name for each item
			item.Store = response.Store
			//Append products to product lists
			resultResponse.Results = append(resultResponse.Results, item)
		}

	}

	//Merged products
	return resultResponse
}

func FetchProductImages(imageSet amazon.Item) []store.Images {

	//List of images for item
	imageList := []store.Images{}

	//Handle images
	for _, img := range imageSet.ImageSets.ImageSet {
		//List of images type
		imageList = append(imageList, store.Images{
			img.SmallImage.URL,
			img.MediumImage.URL,
			img.LargeImage.URL,
		})
	}

	//Images for item
	return imageList
}
