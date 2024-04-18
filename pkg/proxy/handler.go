package proxy

import (
	"encoding/json"
	"fmt"

	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/constants"
	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/dto"
	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/utils"
)

func ProxyHandler(req dto.RPCRequest) (*dto.RPCResponse, error) {
	var response *dto.RPCResponse

	if utils.StringContains(constants.SUPPORTED_METHODS, req.Method) {
		bytesResponse, err := proxyRequest(req)
		if err != nil {
			fmt.Println("Error obtained from proxy:", err)
			return nil, err
		}

		err = json.NewDecoder(bytesResponse).Decode(&response)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return nil, fmt.Errorf("Error decoding JSON")
		}

		return response, nil
	} else {
		return nil, fmt.Errorf("Method not supported: %s", req.Method)
	}
}
