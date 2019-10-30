package httphelper

import (
	"testing"

	"github.com/apptio/d42cli/httphelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilsSuite struct {
	suite.Suite
}

func (suite *UtilsSuite) TestHttpDoRequest() {
	suite.T().Log("BEGIN: Test HTTP GET request")

	// func DoRequest(uri string, target string) string
	BaseURL := "https://jsonplaceholder.typicode.com/"
	uri := "posts/"
	target := "1"

	targetResponse := httphelper.DoRequestImpl("GET", BaseURL, uri, target)
	expectedResponse := "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"sunt aut facere repellat provident occaecati excepturi optio reprehenderit\",\n  \"body\": \"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto\"\n}"
	assert.Equal(suite.T(), expectedResponse, targetResponse, "they should be equal")

	suite.T().Log("END: Test HTTP GET")
}

func TestUtilsSuite(t *testing.T) {
	suite.Run(t, new(UtilsSuite))
}
