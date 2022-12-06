package test

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"gogs.mirlowz.com/x/gf-x-elasticSearch/es"
	"testing"
)

func TestEs(t *testing.T) {
	_ = es.Elastic.Init()

	data, err := es.SendRequest(es.ElasticSearchRequest[esapi.SearchRequest]{
		Request: esapi.SearchRequest{
			Index: []string{"testa"},
			Body: bytes.NewBuffer([]byte(`
				{"query":{
					"match_all": {}
				}}
			`)),
		},
	}.Create())

	fmt.Println(data.Interface(), err)
}
