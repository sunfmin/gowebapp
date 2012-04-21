package accounts

type Account struct {
	Name       string
	Gender     string
	Department string
}

func (a *Account) FormName() string {
	return "Account"
}

func NewAccount() (a *Account) {
	return &Account{}
}
