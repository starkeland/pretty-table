package style

type Color int

// ANSI 转义序列
const (
	FgReset   = "\033[0m"
	FgRed     = "\033[31m"
	FGGreen   = "\033[32m"
	FGYellow  = "\033[33m"
	FGBlue    = "\033[34m"
	FGMagenta = "\033[35m"
	FGCyan    = "\033[36m"
	FGWhite   = "\033[37m"

	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

const (
	TextBold      = "\033[1m"
	TextItalic    = "\033[3m"
	TextUnderline = "\033[4m"
)
