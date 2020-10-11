package clients

import (
	"OMS/placeorder/config"
	"OMS/placeorder/internal/api/request"
	"OMS/placeorder/internal/api/response"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

type InventoryClient struct {
	Config *config.Config
}

func (i *InventoryClient) CheckAndBlockInventory(req *request.GetQuantity) (*response.GetQuantity, error) {
	// request body (payload)
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	log.Info().Msgf("Inventory Req : %+v", req)
	// post some data
	res, err := http.Post(
		i.Config.InventoryUrl,
		"application/json; charset=UTF-8",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return nil, err
	}

	// read response data
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// close response body
	res.Body.Close()

	//Unmarshal response
	inventoryRes := &response.GetQuantity{}
	if err := json.Unmarshal(data, inventoryRes); err != nil {
		return nil, err
	}
	log.Info().Msgf("Inventory Res : %+v", inventoryRes)
	return inventoryRes, nil
}

func (i *InventoryClient) AddItemsToInventory(req *request.AddQuantity) error {
	// request body (payload)
	requestBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	log.Info().Msgf("Add to Inventory Req : %+v", req)
	// post some data
	res, err := http.Post(
		i.Config.AddToInventoryUrl,
		"application/json; charset=UTF-8",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}

	// read response data
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// close response body
	res.Body.Close()

	return nil
}

func (i *InventoryClient) CheckAndUpdateNegativeInventory(req *request.AddQuantity) error {
	// request body (payload)
	requestBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	log.Info().Msgf("Add negative Inventory Req : %+v", req)
	// post some data
	res, err := http.Post(
		i.Config.AddToInventoryUrl,
		"application/json; charset=UTF-8",
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return err
	}

	// read response data
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// close response body
	res.Body.Close()

	return nil
}

