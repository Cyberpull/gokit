package tests

import (
	"testing"

	"github.com/Cyberpull/gokit/dbo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DBOTestSuite struct {
	suite.Suite
}

func (x *DBOTestSuite) TestSet() {
	var data dbo.Set[string]

	err := data.Scan("a,b,c,d,e")
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), []string{"a", "b", "c", "d", "e"}, data.Data)

	value, err := data.Value()
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "a,b,c,d,e", value)
}

// ===============================

func TestDBO(t *testing.T) {
	suite.Run(t, new(DBOTestSuite))
}
