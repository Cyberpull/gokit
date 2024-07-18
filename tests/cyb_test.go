package tests

import (
	"os"
	"testing"

	"cyberpull.com/gokit/cyb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CYBTestSuite struct {
	suite.Suite

	server cyb.Server
	client cyb.Client
}

func (x *CYBTestSuite) SetupSuite() {
	address := os.TempDir() + "/demo.cyb.sock"

	// Start GoKit CYB Server
	require.NoError(x.T(), startCybServer(&x.server, address))

	// Start GoKit CYB Client
	require.NoError(x.T(), startCybClient(&x.client, address))
}

func (x *CYBTestSuite) TearDownSuite() {
	// Stop GoKit CYB Client
	require.NoError(x.T(), x.client.Stop())

	// Stop GoKit CYB Server
	require.NoError(x.T(), x.server.Stop())
}

func (x *CYBTestSuite) TestRequest() {
	value, err := x.client.Request("GET", "/test/request", nil)
	require.NoError(x.T(), err)
	assert.Equal(x.T(), "Demo Request Successful", value.Content)
}

func (x *CYBTestSuite) TestUpdate() {
	//
}

func TestCYB(t *testing.T) {
	suite.Run(t, new(CYBTestSuite))
}
