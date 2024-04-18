package proxy

import (
	"testing"

	"github.com/adityaGaneshJajpure/rpc-proxy-server/pkg/dto"
)

func TestIndex(t *testing.T) {
	t.Run("Get eth_blockNumber", func(t *testing.T) {
		request := dto.RPCRequest{
			JSONRPC: "2.0",
			Method:  "eth_blockNumber",
			Params:  []interface{}{},
			ID:      0,
		}

		response, err := ProxyHandler(request)

		if response == nil {
			t.Errorf("Expected response to be present")
		}

		if err != nil {
			t.Errorf("Expected error to be nil, but got %v", err.Error())
		}

	})

	t.Run("Get eth_getBlockByNumber", func(t *testing.T) {
		request := dto.RPCRequest{
			JSONRPC: "2.0",
			Method:  "eth_getBlockByNumber",
			Params: []interface{}{
				"0x6C59B1",
				false,
			},
			ID: 0,
		}

		response, err := ProxyHandler(request)

		if response == nil {
			t.Errorf("Expected response to be present")
		}

		if err != nil {
			t.Errorf("Expected error to be nil, but got %v", err.Error())
		}

	})

	t.Run("Run any unsupported method", func(t *testing.T) {
		request := dto.RPCRequest{
			JSONRPC: "2.0",
			Method:  "eth_getBlockReceipts",
			Params: []interface{}{
				"0x6C59B1",
			},
			ID: 0,
		}

		response, err := ProxyHandler(request)

		if response != nil {
			t.Errorf("Expected response to be nil, but is presen %v", response)
		}

		if err == nil {
			t.Errorf("Expected error to be present, but got nil")
		}

		if err.Error() != "Method not supported: eth_getBlockReceipts" {
			t.Errorf("Expected error message not found, but got %v", err.Error())
		}

	})
}
