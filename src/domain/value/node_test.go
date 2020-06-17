package value

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NodeTestSuite struct {
	suite.Suite
}

func (s *NodeTestSuite) SetupSuite() {
}

func TestNodeTestSuite(t *testing.T) {
	suite.Run(t, new(NodeTestSuite))
}

// NewNode :: for success scenarios :: creates a new node.
func (s *NodeTestSuite) TestNewNode() {
	tcs := []struct{
		title string
		nodeName string
	}{
		{
			title: "ForAValidName",
			nodeName: "node name",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		s.Run(tc.title, func() {

			node, err := NewNode(WithName(tc.nodeName))

			s.NoError(err)
			s.Equal(tc.nodeName, node.Name(), "Node name must match")
		})
	}
}

// NewNode :: without any options :: returns an error.
func (s *NodeTestSuite) TestNewNodeWithoutFunctionalOptions() {

	node, err := NewNode()

	s.Empty(node, "Node value object must be empty")
	s.Error(err)
}

// New Node :: for an empty name :: returns an error.
func (s *NodeTestSuite) TestNewNodeForAnEmptyName() {

	node, err := NewNode(WithName(""))

	s.Empty(node, "Node value object must be empty")
	s.Error(err)
}