package usecase

type (
	SendMoneyUsecase interface {
		Do(in SendMoneyInput) (*SendMoneyOutput, error)
	}
	SendMoneyInput struct {
		SendingAccountNumber   string
		RecipientAccountNumber string
		Amount                 float64
	}
	SendMoneyOutput struct {
		TransactionID int64
	}
	sendMoneyUsecase struct {
	}
)

func NewSendMoneyUsecase() *SendMoneyUsecase {
	return &sendMoneyUsecase{}
}

func (u *sendMoneyUsecase) Do(in SendMoneyInput) (*SendMoneyOutput, error) {

}
