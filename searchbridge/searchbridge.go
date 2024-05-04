package searchbridge

type Engine interface {

	Search(query string) []byte

}
