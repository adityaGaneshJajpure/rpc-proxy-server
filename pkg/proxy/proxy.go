package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/constants"
	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/dto"
)

func proxyRequest(req dto.RPCRequest) (*bytes.Buffer, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return nil, fmt.Errorf("Error marshaling struct to JSON")
	}

	// Create a new HTTP request with POST method and request body
	httpReq, err := http.NewRequest("POST", os.Getenv(constants.RPC_PROVIDER), bytes.NewBuffer([]byte(reqData)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, fmt.Errorf("Error proxy request")
	}

	// TODO: set additional headers if required (based on RPC provider)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	// Send the request
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, fmt.Errorf("Error sending request to RPC provider")
	}

	defer resp.Body.Close()

	// check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response status code:", resp.StatusCode)
		return nil, fmt.Errorf("Error obtained from RPC provider")
	}

	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, fmt.Errorf("Error obtained from RPC provider")
	}

	return responseBody, nil
}
