package style

type Color string

// ANSI 转义序列
const (
	FgReset     Color = "\033[0m"
	FgRed             = "\033[31m"
	FgHiRed           = "\033[91m"
	FgGreen           = "\033[32m"
	FgHiGreen         = "\033[92m"
	FgYellow          = "\033[33m"
	FgHiYellow        = "\033[93m"
	FgBlue            = "\033[34m"
	FgHiBlue          = "\033[94m"
	FgMagenta         = "\033[35m"
	FgHiMagenta       = "\033[95m"
	FgCyan            = "\033[36m"
	FgHiCyan          = "\033[96m"
	FgWhite           = "\033[37m"
	FgHiWhite         = "\033[97m"

	BgReset     Color = "\033[0m"
	BgRed             = "\033[41m"
	BgHiRed           = "\033[101m"
	BgGreen           = "\033[42m"
	BgHiGreen         = "\033[102m"
	BgYellow          = "\033[43m"
	BgHiYellow        = "\033[103m"
	BgBlue            = "\033[44m"
	BgHiBlue          = "\033[104m"
	BgMagenta         = "\033[45m"
	BgHiMagenta       = "\033[105m"
	BgCyan            = "\033[46m"
	BgHiCyan          = "\033[106m"
	BgWhite           = "\033[47m"
	BgHiWhite         = "\033[107m"
)

const (
	TextBold      Color = "\033[1m"
	TextItalic          = "\033[3m"
	TextUnderline       = "\033[4m"
	TextReset           = "\033[0m"
)
