package e2e

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) Test_HealthCheck_ExpectValidResponse() {
	c := http.Client{}

	res, err := c.Get("http://localhost:8080/health")
	s.Require().NoError(err)

	s.Equal(http.StatusOK, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	s.Require().NoError(err)

	s.JSONEq(`{"message": "Server is Ok"}`, string(body))
}

func (s *EndToEndSuite) Test_HealthCheck_ExpectInvalidResponse() {
	c := http.Client{}

	res, err := c.Get("http://localhost:8080/health")
	s.Require().NoError(err)

	s.NotEqual(http.StatusInternalServerError, res.StatusCode)

}
