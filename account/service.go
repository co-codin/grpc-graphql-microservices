package account

type Service interface {
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
