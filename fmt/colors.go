package fmt

import "github.com/Cyberpull/gokit/color"

var (
	Green   xColorFmt
	Red     xColorFmt
	Cyan    xColorFmt
	Magenta xColorFmt
	Black   xColorFmt
	Blue    xColorFmt
	White   xColorFmt
	Yellow  xColorFmt

	HiBlack   xColorFmt
	HiBlue    xColorFmt
	HiCyan    xColorFmt
	HiGreen   xColorFmt
	HiMagenta xColorFmt
	HiRed     xColorFmt
	HiWhite   xColorFmt
	HiYellow  xColorFmt

	BgGreen   xColorFmt
	BgRed     xColorFmt
	BgCyan    xColorFmt
	BgMagenta xColorFmt
	BgBlack   xColorFmt
	BgBlue    xColorFmt
	BgWhite   xColorFmt
	BgYellow  xColorFmt

	BgHiBlack   xColorFmt
	BgHiBlue    xColorFmt
	BgHiCyan    xColorFmt
	BgHiGreen   xColorFmt
	BgHiMagenta xColorFmt
	BgHiRed     xColorFmt
	BgHiWhite   xColorFmt
	BgHiYellow  xColorFmt
)

func init() {
	Green.color = *color.New(color.FgGreen)
	Red.color = *color.New(color.FgRed)
	Cyan.color = *color.New(color.FgCyan)
	Magenta.color = *color.New(color.FgMagenta)
	Black.color = *color.New(color.FgBlack)
	Blue.color = *color.New(color.FgBlue)
	White.color = *color.New(color.FgWhite)
	Yellow.color = *color.New(color.FgYellow)

	HiBlack.color = *color.New(color.FgHiBlack)
	HiBlue.color = *color.New(color.FgHiBlue)
	HiCyan.color = *color.New(color.FgHiCyan)
	HiGreen.color = *color.New(color.FgHiGreen)
	HiMagenta.color = *color.New(color.FgHiMagenta)
	HiRed.color = *color.New(color.FgHiRed)
	HiWhite.color = *color.New(color.FgHiWhite)
	HiYellow.color = *color.New(color.FgHiYellow)

	// Background ======================

	BgGreen.color = *color.New(color.BgGreen)
	BgRed.color = *color.New(color.BgRed)
	BgCyan.color = *color.New(color.BgCyan)
	BgMagenta.color = *color.New(color.BgMagenta)
	BgBlack.color = *color.New(color.BgBlack)
	BgBlue.color = *color.New(color.BgBlue)
	BgWhite.color = *color.New(color.BgWhite)
	BgYellow.color = *color.New(color.BgYellow)

	BgHiBlack.color = *color.New(color.BgHiBlack)
	BgHiBlue.color = *color.New(color.BgHiBlue)
	BgHiCyan.color = *color.New(color.BgHiCyan)
	BgHiGreen.color = *color.New(color.BgHiGreen)
	BgHiMagenta.color = *color.New(color.BgHiMagenta)
	BgHiRed.color = *color.New(color.BgHiRed)
	BgHiWhite.color = *color.New(color.BgHiWhite)
	BgHiYellow.color = *color.New(color.BgHiYellow)
}
