package num2word

var uaRepl = [][]string{
	// t - тисячі; m - мільони; M - мільярди;
	{",,,,,,", "eM"},
	{",,,,,", "em"},
	{",,,,", "et"},
	// e - одиниці; d - десятки; c - сотні;
	{",,,", "e"},
	{",,", "d"},
	{",", "c"},
	{"0c0d0et", ""},
	{"0c0d0em", ""},
	{"0c0d0eM", ""},
	// --
	{"0c", ""},
	{"1c", "сто "},
	{"2c", "двісті "},
	{"3c", "триста "},
	{"4c", "чотириста "},
	{"5c", "п'ятсот "},
	{"6c", "шістсот "},
	{"7c", "сімсот "},
	{"8c", "вісімсот "},
	{"9c", "дев'ятсот "},
	{"1d0e", "десять "},
	{"1d1e", "одинадцять "},
	{"1d2e", "дванадцять "},
	{"1d3e", "тринадцять "},
	{"1d4e", "чотирнадцять "},
	{"1d5e", "п'ятнадцять "},
	{"1d6e", "шістнадцять "},
	{"1d7e", "сімнадцять "},
	{"1d8e", "вісімнадцять "},
	{"1d9e", "дев'ятнадцять "},
	// --
	{"0d", ""},
	{"2d", "двадцять "},
	{"3d", "тридцять "},
	{"4d", "сорок "},
	{"5d", "п'ятдесят "},
	{"6d", "шістдесят "},
	{"7d", "сімдесят "},
	{"8d", "вісімдесят "},
	{"9d", "дев'яносто "},
	// --
	{"0e", ""},
	{"5e", "п'ять "},
	{"6e", "шість "},
	{"7e", "сім "},
	{"8e", "вісім "},
	{"9e", "дев'ять "},
	// --
	{"1e.", "одна гривня "},
	{"2e.", "дві гривні "},
	{"3e.", "три гривні "},
	{"4e.", "чотири гривні "},
	{"1et", "одна тисяча "},
	{"2et", "дві тисячі "},
	{"3et", "три тисячі "},
	{"4et", "чотири тисячі "},
	{"1em", "один мільйон "},
	{"2em", "два мільйони "},
	{"3em", "три мільйони "},
	{"4em", "чотири мільйони "},
	{"1eM", "один мільярд "},
	{"2eM", "два мільярди "},
	{"3eM", "три мільярди "},
	{"4eM", "чотирі мільярди "},
	//  блок для написання копійок без скорочення "коп"
	{"11k", "11 копійок"},
	{"12k", "12 копійок"},
	{"13k", "13 копійок"},
	{"14k", "14 копійок"},
	{"1k", "1 копійка"},
	{"2k", "2 копійки"},
	{"3k", "3 копійки"},
	{"4k", "4 копійки"},
	{"k", " копійок"},
	// --
	{".", "гривень "},
	{"t", "тисяч "},
	{"m", "мільйонів "},
	{"M", "мільярдів "},
}

// UaMoney - гроші прописом українською
func UaAmount(number float64, upperFirstChar bool) string {

	return Num2word(number, upperFirstChar, uaRepl)
}

