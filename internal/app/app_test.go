package app

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ameteiko/mindnet/internal/platform/testdata"
)

type AppTestSuite struct {
	suite.Suite

	// Expectations.
	validNodeName string
}

func (s *AppTestSuite) SetupSuite() {
	s.validNodeName = testdata.Slug("node name")
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}

func (s *AppTestSuite) TestCreateNode() {
	app := App{}

	node, err := app.CreateNode(s.validNodeName)

	s.NoError(err)
	s.Equal(s.validNodeName, node.Name())
}
