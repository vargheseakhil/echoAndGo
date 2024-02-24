package handlers

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPostIndexHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1213/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndToEndSuite) TestNoPostResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:1213/post/1234")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status": "ok", "data": []}`, string(b))
}
