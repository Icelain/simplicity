package searx

import "net/http"

type SearxEngine struct {
	
	httpclient *http.Client
	request *http.Request

}

//func NewSearxEngine() *SearxEngine {
//
//	client := &http.Client{}
//	req, err := http.NewRequest()
//
//}

//func (sx *SearxEngine) Search(query string) []byte {
//
//		
//
//}
