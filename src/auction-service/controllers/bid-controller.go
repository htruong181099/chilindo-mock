package controllers

import "chilindo/src/auction-service/services"

type IBidController interface {
}

type BidController struct {
	BidService services.IBidService
}
