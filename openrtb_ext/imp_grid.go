package openrtb_ext

import "encoding/json"

// ExtImpGrid defines the contract for bidrequest.imp[i].ext.prebid.bidder.grid
type ExtImpGrid struct {
	Uid      int             `json:"uid"`
	Keywords json.RawMessage `json:"keywords,omitempty"`
}
