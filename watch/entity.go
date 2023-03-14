package watch

type watchRequest struct {
	Id uint64 `json:"id"`
}

type auction struct {
	AuctionInfo []watchRequest `json:"auctionInfo"`
}
