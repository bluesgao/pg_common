package currency

type FiatCurrency struct {
	Code     string // 货币代码，如 USD
	Name     string // 货币名称，如 US Dollar
	Symbol   string // 货币符号，如 $
	Decimals int    // 小数位精度
}

var FiatCurrencies = map[string]FiatCurrency{
	"USD": {"USD", "US Dollar", "$", 2},
	"CNY": {"CNY", "Chinese Yuan", "¥", 2},
	"EUR": {"EUR", "Euro", "€", 2},
	"JPY": {"JPY", "Japanese Yen", "¥", 0},
	"GBP": {"GBP", "British Pound", "£", 2},
	"KRW": {"KRW", "Korean Won", "₩", 0},
	"HKD": {"HKD", "Hong Kong Dollar", "HK$", 2},
}

// 是否是合法法币
func IsValidFiat(code string) bool {
	_, ok := FiatCurrencies[code]
	return ok
}

// 获取精度
func GetDecimals(code string) int {
	if c, ok := FiatCurrencies[code]; ok {
		return c.Decimals
	}
	return 2 // 默认精度
}

// 获取符号，如 "$"
func GetSymbol(code string) string {
	if c, ok := FiatCurrencies[code]; ok {
		return c.Symbol
	}
	return ""
}

// 获取名称
func GetDisplayName(code string) string {
	if c, ok := FiatCurrencies[code]; ok {
		return c.Name
	}
	return ""
}
