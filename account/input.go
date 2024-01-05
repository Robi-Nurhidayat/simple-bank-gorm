package account

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}


type UpdateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}


type IdFromUri struct {
	ID int32 `uri:"id" binding:"required"`
}