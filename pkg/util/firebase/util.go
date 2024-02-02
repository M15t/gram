package firebase

import (
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
)

func fetchFirebaseCredentials(url string) ([]byte, error) {
	// Send an HTTP GET request to fetch the credentials
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}

	// Read the response body
	credentials, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func mapToStruct(input map[string]interface{}, output interface{}) error {
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &output,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	decoder.Decode(input)

	return nil
}

func mapToFireStoreUpdate(input map[string]interface{}) []firestore.Update {
	return nil
}
