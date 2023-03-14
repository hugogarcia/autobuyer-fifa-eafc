package entity

type Trade struct {
	TradeId     uint64
	TradeState  string
	BuyNowPrice uint32
	CurrentBid  uint32
	Watched     bool
	BidState    string
	StartingBid uint32
	Expires     int32
	TradeOwner  bool
	ItemData    ItemData
}

type ItemData struct {
	ResourceId uint64
	AssetId uint64
}

type TradeResponse struct {
	Credits     uint64
	AuctionInfo []Trade
}

