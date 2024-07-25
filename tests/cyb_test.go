package tests

import (
	"fmt"
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
	socket := fmt.Sprintf("unix:%s/test.cyb.sock", os.TempDir())

	// Start GoKit CYB Server
	require.NoError(x.T(), startCybServer(&x.server, socket))

	// Start GoKit CYB Client
	require.NoError(x.T(), startCybClient(&x.client, socket))
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

	x.client.On("GET", "/test/update", func(data cyb.OutputData) {
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

func (x *CYBTestSuite) TestError() {
	_, err := cyb.MakeRequest[string](&x.client, "GET", "/test/error", nil)
	require.Error(x.T(), err)
	assert.Equal(x.T(), "Demo Error Successful", err.Error())
}

func (x *CYBTestSuite) TestStructResponse() {
	resp, err := cyb.MakeRequest[DemoResponse](&x.client, "GET", "/test/struct", nil)
	require.NoError(x.T(), err)

	assert.IsType(x.T(), DemoResponse{}, resp)
	assert.Equal(x.T(), "Christian", resp.Name)
	assert.Equal(x.T(), "demo@example.com", resp.Email)
}

func (x *CYBTestSuite) TestStructUpdate() {
	updateChan := make(chan gokit.IOData[DemoResponse], 1)

	x.client.On("GET", "/test/struct/update", func(data cyb.OutputData) {
		var update gokit.IOData[DemoResponse]

		update.Error = data.Bind(&update.Data)

		updateChan <- update
	})

	value, err := cyb.MakeRequest[string](&x.client, "GET", "/test/struct/update", nil)
	require.NoError(x.T(), err)
	assert.Equal(x.T(), "Struct Update Successful", value)

	update := <-updateChan
	require.NoError(x.T(), update.Error)

	data := update.Data
	assert.IsType(x.T(), DemoResponse{}, data)
	assert.Equal(x.T(), "Christian", data.Name)
	assert.Equal(x.T(), "demo@example.com", data.Email)
}

func (x *CYBTestSuite) TestStructRequest() {
	resp, err := cyb.MakeRequest[string](&x.client, "GET", "/test/struct/request", DemoRequest{
		Name:  "Christian",
		Email: "demo@example.com",
	})

	require.NoError(x.T(), err)
	assert.Equal(x.T(), "Success!", resp)
}

// ===============================

func TestCYB(t *testing.T) {
	suite.Run(t, new(CYBTestSuite))
}
