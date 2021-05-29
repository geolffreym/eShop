package functions

import "errors"

func ExtractParam(searchIndex string, params []map[string]string) (string, error) {
	//Extract and validate params
	for _, v := range params {
		if z, ok := v[searchIndex]; ok {
			return z, nil
		}
	}

	// Return error
	return "", errors.New("bad index")
}
