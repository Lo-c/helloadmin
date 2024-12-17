package acme

type FindRequest struct {
	Id string `json:"id"`
}

type UpdateRequest struct {
	Id       uint   `json:"id" example:"1"`
	AcmeData string `json:"acme_path" example:"/root/.acme.sh"`
	Email    string `json:"email" example:"foo@bar.com"`
}

type AcmeData struct {
	Id       uint   `json:"id" example:"1"`
	AcmePath string `json:"acme_path" example:"/root/.acme.sh"`
	Email    string `json:"email" example:"foo@bar.com"`
}

type Response struct {
	Acme []AcmeData `json:"acme"`
}
