package thread

import (
	"app/stores/interface/store"
	"fmt"
	"time"
	"app/stores/interface/response"
)

func  NewChannel(i []store.IStore) chan *stres.RThread {
	var ch = make(chan *stres.RThread, len(i)) // buffered

	//Make thread call
	for _, storeList := range i {
		go func(st store.IStore) {
			//The client Store
			client := st.GetStore()
			//The operation to exec
			operation := st.GetOperation()

			//Get search result items
			fmt.Printf("Fetching %s \n", client.Store)
			result := operation(client, st.GetParameters())

			//Add result to channel
			ch <- &stres.RThread{
				result.Result,
				result.Total,
				result.Pages,
				client.Store,
				len(i)}
		}(storeList)
	}

	return ch

}

func ProcessChannelResponse(ch chan *stres.RThread) []*stres.RThread {
	//Array of responses coming from channel
	responses := []*stres.RThread{}

	//For each channel response
	for {
		select {
		case response := <-ch:
			fmt.Printf("\nFetched %s search", response.Store)
			//Append channels response
			responses = append(responses, response)

			//Verify each request by correspondent response
			if len(responses) == response.ChannelSize {
				close(ch)
				return responses
			}

		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}

	}

	//Return joined channel responses
	return responses
}
