package tests

import (
	"os"
	"testing"

	"cyberpull.com/gokit"
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
	opts := cyb.Options{
		Network:    "unix",
		SocketPath: os.TempDir() + "/test.cyb.sock",
	}

	// opts := cyb.Options{
	// 	Network:    "tcp",
	// 	Host:    "127.0.0.1",
	// 	Port:    1988,
	// }

	// Start GoKit CYB Server
	require.NoError(x.T(), startCybServer(&x.server, opts))

	// Start GoKit CYB Client
	require.NoError(x.T(), startCybClient(&x.client, opts))
}

func (x *CYBTestSuite) TearDownSuite() {
	// Stop GoKit CYB Client
	require.NoError(x.T(), x.client.Stop())

	// Stop GoKit CYB Server
	require.NoError(x.T(), x.server.Stop())
}

func (x *CYBTestSuite) TestRequest() {
	value, err := cyb.MakeRequest[string](&x.client, "GET", "/test/request", nil)
	require.NoError(x.T(), err)

	assert.Equal(x.T(), "Demo Request Successful", value)
}

func (x *CYBTestSuite) TestUpdate() {
	updateChan := make(chan gokit.IOData[string], 1)

	x.client.On("GET", "/test/update", func(data cyb.Data) {
		var update gokit.IOData[string]

		update.Error = data.Bind(&update.Data)

		updateChan <- update
	})

	value, err := cyb.MakeRequest[string](&x.client, "GET", "/test/update", nil)
	require.NoError(x.T(), err)
	assert.Equal(x.T(), "Demo Update Successful", value)

	update := <-updateChan
	require.NoError(x.T(), update.Error)
	assert.Equal(x.T(), "Demo Update Successful", update.Data)
}

func TestCYB(t *testing.T) {
	suite.Run(t, new(CYBTestSuite))
}
