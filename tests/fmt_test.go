package tests

import (
	"testing"

	"github.com/Cyberpull/gokit/color"
	"github.com/Cyberpull/gokit/fmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FmtTestSuite struct {
	suite.Suite
}

func (x *FmtTestSuite) TestColor() {
	require.NotPanics(x.T(), func() {
		logger := fmt.Color(color.BgCyan)
		logger.Println("Testing")
	})
}

func (x *FmtTestSuite) TestRedColor() {
	require.NotPanics(x.T(), func() {
		fmt.Red.Println("Testing")
	})
}

// ===============================

func TestFmt(t *testing.T) {
	suite.Run(t, new(FmtTestSuite))
}
