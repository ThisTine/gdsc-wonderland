package payload

type ParingCommit struct {
	SessionNo *string `json:"sessionNo" validate:"required"`
	ItemNo    *string `json:"itemNo" validate:"required"`
}

type ParingCommitResponse struct {
	Matched     *bool   `json:"matched"`
	ForwardLink *string `json:"forwardLink"`
	PairedWith  *string `json:"pairedWith"`
}
