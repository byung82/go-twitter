package twitter

type Twitter struct {
	Data     Data     `json:"data"`
	Includes *Include `json:"includes"`
}
