package requests

type ShortenUrlRequest struct {
	URL string `json:"url" binding:"required"`
}
