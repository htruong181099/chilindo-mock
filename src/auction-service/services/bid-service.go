package services

import (
	"chilindo/src/auction-service/dtos"
	"chilindo/src/auction-service/models"
	"chilindo/src/auction-service/repository"
	"errors"
	"log"
)

type IBidService interface {
	GetBidsOfAuction(dto *dtos.AuctionIdDTO) (*[]models.Bid, error)
	GetBidById(dto *dtos.BidIdDTO) (*models.Bid, error)
	CreateBid(dto *dtos.CreateBidDTO) (*models.Bid, error)
}

type BidService struct {
	BidRepository     repository.IBidRepository
	AuctionRepository repository.IAuctionRepository
}

func NewBidService(bidRepository repository.IBidRepository, auctionRepository repository.IAuctionRepository) *BidService {
	return &BidService{BidRepository: bidRepository, AuctionRepository: auctionRepository}
}

func (b BidService) GetBidsOfAuction(dto *dtos.AuctionIdDTO) (*[]models.Bid, error) {
	bid, err := b.BidRepository.GetBidsOfAuction(dto)
	if err != nil {
		log.Println("GetBidsOfAuction: Error to get bid in package service", err)
	}
	return bid, nil
} //Done

func (b BidService) GetBidById(dto *dtos.BidIdDTO) (*models.Bid, error) {
	bids, err := b.BidRepository.GetBidById(dto)
	if err != nil {
		log.Println("GetBidById: Error to get bid in package service", err)
	}
	return bids, nil
} //Done

func (b BidService) CreateBid(dto *dtos.CreateBidDTO) (*models.Bid, error) {
	auction, aErr := b.AuctionRepository.GetAuctionById(&dtos.AuctionIdDTO{
		AuctionId: dto.Bid.AuctionId,
	})
	if aErr != nil {
		log.Println("CreateBid: Error to get auction in package service", aErr)
		return nil, aErr
	}
	if auction == nil {
		return nil, nil
	}

	if dto.Bid.Amount < auction.LowestBid {
		return nil, errors.New("CreateBid: Invalid amount")
	}

	lastBid, prevErr := b.BidRepository.UpdateLastBid(&dtos.AuctionIdDTO{
		AuctionId: dto.Bid.AuctionId,
	})
	if prevErr != nil {
		log.Println("CreateBid: Error to get auction in package service", aErr)
		return nil, aErr
	}
	log.Println("lb: ", lastBid)
	if lastBid != nil {
		log.Println("bid ", dto.Bid.Amount)
		log.Println("lastbid ", lastBid.Id, lastBid.Amount)

		if dto.Bid.Amount <= lastBid.Amount {
			return nil, errors.New("CreateBid: Invalid amount")
		}

		if dto.Bid.BidderId == lastBid.BidderId {
			return nil, errors.New("CreateBid: Invalid bidder")
		}
	}

	bid, err := b.BidRepository.CreateBid(dto)
	if err != nil {
		log.Println("CreateBid: Error to create bid in package service", err)
		return nil, err
	}
	return bid, nil
} //Done
