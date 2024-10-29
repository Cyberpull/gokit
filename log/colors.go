package log

import "github.com/Cyberpull/gokit/color"

var (
	Cyan    Logger
	Red     Logger
	Green   Logger
	Magenta Logger
	Yellow  Logger
	Black   Logger
	Blue    Logger
	White   Logger

	HiCyan    Logger
	HiRed     Logger
	HiGreen   Logger
	HiMagenta Logger
	HiYellow  Logger
	HiBlack   Logger
	HiBlue    Logger
	HiWhite   Logger

	BgCyan    Logger
	BgRed     Logger
	BgGreen   Logger
	BgMagenta Logger
	BgYellow  Logger
	BgBlack   Logger
	BgBlue    Logger
	BgWhite   Logger

	BgHiCyan    Logger
	BgHiRed     Logger
	BgHiGreen   Logger
	BgHiMagenta Logger
	BgHiYellow  Logger
	BgHiBlack   Logger
	BgHiBlue    Logger
	BgHiWhite   Logger
)

func init() {
	setDefault(&Cyan, color.FgCyan)
	setDefault(&Red, color.FgRed)
	setDefault(&Green, color.FgGreen)
	setDefault(&Magenta, color.FgMagenta)
	setDefault(&Yellow, color.FgYellow)
	setDefault(&Black, color.FgBlack)
	setDefault(&Blue, color.FgBlue)
	setDefault(&White, color.FgWhite)

	setDefault(&HiCyan, color.FgHiCyan)
	setDefault(&HiRed, color.FgHiRed)
	setDefault(&HiGreen, color.FgHiGreen)
	setDefault(&HiMagenta, color.FgHiMagenta)
	setDefault(&HiYellow, color.FgHiYellow)
	setDefault(&HiBlack, color.FgHiBlack)
	setDefault(&HiBlue, color.FgHiBlue)
	setDefault(&HiWhite, color.FgHiWhite)

	// Background ===============

	setDefault(&BgCyan, color.BgCyan)
	setDefault(&BgRed, color.BgRed)
	setDefault(&BgGreen, color.BgGreen)
	setDefault(&BgMagenta, color.BgMagenta)
	setDefault(&BgYellow, color.BgYellow)
	setDefault(&BgBlack, color.BgBlack)
	setDefault(&BgBlue, color.BgBlue)
	setDefault(&BgWhite, color.BgWhite)

	setDefault(&BgHiCyan, color.BgHiCyan)
	setDefault(&BgHiRed, color.BgHiRed)
	setDefault(&BgHiGreen, color.BgHiGreen)
	setDefault(&BgHiMagenta, color.BgHiMagenta)
	setDefault(&BgHiYellow, color.BgHiYellow)
	setDefault(&BgHiBlack, color.BgHiBlack)
	setDefault(&BgHiBlue, color.BgHiBlue)
	setDefault(&BgHiWhite, color.BgHiWhite)
}
