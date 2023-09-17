package payload

type ParingCommit struct {
	SessionNo *string `json:"sessionNo"`
	ItemNo    *string `json:"itemNo"`
}

type ParingCommitResponse struct {
	Matched     *bool   `json:"matched"`
	ForwardLink *string `json:"forwardLink"`
}
