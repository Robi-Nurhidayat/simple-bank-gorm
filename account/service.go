package account

type AccountService interface {
	CreateAccount(input CreateAccountRequest) (Account, error)
	ListAccounts() ([]Account, error)
	UpdateAccount(account Account) (Account,error)
	GetAccountByID(id int32)(Account,error)
	DeleteById(id int32) error
}

type service struct {
	repository AccountRepository
}

func NewAccountService(accountRepository AccountRepository) AccountService {
	return &service{
		repository: accountRepository,
	}
}

func (s *service) CreateAccount(input CreateAccountRequest) (Account, error) {

	account := Account{
		Owner:    input.Owner,
		Balance:  input.Balance,
		Currency: input.Currency,
	}

	newAccount, err := s.repository.CreateAccount(account)

	if err != nil {
		return newAccount, err
	}

	return newAccount, nil
}

func (s *service) ListAccounts() ([]Account, error) {

	accounts, err := s.repository.ListAccounts()

	if err != nil {
		return accounts, err
	}

	return accounts, nil
}


func (s *service) UpdateAccount(account Account) (Account, error) {

	account,err := s.repository.UpdateAccount(account)

	if err != nil {
		return account,err
	}

	return account,nil

}

func (s *service) GetAccountByID(id int32) (Account, error) {

	account,err := s.repository.GetAccountByID(id)

	if err != nil {
		return account,err
	}

	return account,nil

}

func (s *service) DeleteById(id int32) error {
	err := s.repository.DeleteById(id)

	if err != nil {
		return err
	}

	return nil
}