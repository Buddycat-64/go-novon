package core

// Define the message structure
type GetTotalPointsMessage struct {
	Type      string `json:"type"`
	Timestamp int64  `json:"timestamp"`
}

// Define options structure (equivalent to the JS options object)
type SendOptions struct {
	NoReply         bool `json:"noReply"`
	ResponseTimeout int  `json:"responseTimeout"`
}

// Define a struct that matches the JSON structure
type TotalPointsResponse struct {
	Type                string `json:"type"`
	TotalPoints         int    `json:"totalPoints"`
	BalanceMultiplier   int    `json:"balanceMultiplier"`
	NextTierMultiplier  int    `json:"nextTierMultiplier"`
	NextTierNknRequired int    `json:"nextTierNknRequired"`
	Timestamp           int64  `json:"timestamp"`
}
