package openrtb_ext

// ExtImpGamma defines the contract for bidrequest.imp[i].ext.prebid.bidder.gamma
type ExtImpGamma struct {
	PartnerID string `json:"id"`
	ZoneID    string `json:"zid"`
	WebID     string `json:"wid"`
}
