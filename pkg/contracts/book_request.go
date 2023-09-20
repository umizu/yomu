package contracts

type BookRequest struct {
	Title     string `json:"title"`
	MediaType string `json:"mediaType"`
	Length    int    `json:"length"`
}
