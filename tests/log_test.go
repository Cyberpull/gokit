package tests

import (
	"testing"

	"github.com/Cyberpull/gokit/color"
	"github.com/Cyberpull/gokit/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LogTestSuite struct {
	suite.Suite
}

func (x *LogTestSuite) TestColor() {
	require.NotPanics(x.T(), func() {
		logger := log.Color(color.BgCyan)
		logger.Println("Testing")
	})
}

// ===============================

func TestLog(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}
