package service

import (
	"OMS/inventory/commons"
	"OMS/inventory/internal/api/request"
	"OMS/inventory/internal/api/response"
	"OMS/inventory/internal/cache"
	"OMS/inventory/internal/db"
	"encoding/binary"

	"github.com/rs/zerolog/log"
)

type Inventory struct {
	AppCache *cache.AppCache
	Mysql    *db.MysqlDB
}

func (i *Inventory) GetProductQuantity(req *request.GetQuantity) (*response.GetQuantity, error) {
	if err := commons.ValidateGetQuantityReq(req); err != nil {
		return nil, err
	}
	if req.BlockQuantity > 0 {
		q, err := i.Mysql.GetAndBlockProductQuantity(req.ProductId, req.BlockQuantity)
		if err != nil {
			return nil, err
		}
		return &response.GetQuantity{Quantity: q}, nil
	}

	if req.FromCache {
		q := 0
		qByte, err := i.AppCache.Get(req.ProductId)
		if err != nil {
			log.Error().Err(err).Msgf("Unable to fetch quantity of Item : %v", req.ProductId)
			q, _ = i.Mysql.GetProductQuantity(req.ProductId)
		} else {
			q = int(binary.LittleEndian.Uint64(qByte))
		}
		res := &response.GetQuantity{
			Quantity: q,
		}
		log.Info().Interface("Get Q Qes - ", res).Msg("GetProductQuantity API")
		return res, nil
	}
	return nil, nil
}

func (i *Inventory) AddProductQuantity(req *request.AddQuantity) error {
	if err := commons.ValidateAddQuantityReq(req); err != nil {
		return err
	}
	if req.AddQuantity > 0 {
		return i.Mysql.UpdateProductQuantity(req.ProductId, req.AddQuantity)
	}
	return nil
}

func (i *Inventory) AddNegativeProductQuantity(req *request.AddQuantity) error {
	if err := commons.ValidateAddQuantityReq(req); err != nil {
		return err
	}
	if req.AddQuantity > 0 {
		return i.Mysql.AddNegativeProductQuantity(req.ProductId, req.AddQuantity)
	}
	return nil
}
