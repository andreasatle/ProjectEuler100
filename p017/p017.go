package p017

func LetterCountV1() int {
	charCnt := 0
	for i := 1; i <= 1000; i++ {
		str := ""
		subCent := i % 100
		superCent := i / 100
		if i == 1000 {
			str = "onethousand"
		} else if superCent > 0 {
			str = digitString(superCent) + "hundred"
		}
		if superCent > 0 && subCent > 0 {
			str += "and"
		}
		if subCent > 0 {
			str += subCentString(subCent)
		}
		charCnt += len(str)
	}
	return charCnt
}

func LetterCountV2() int {
	charCount := 0
	for i := 1; i <= 1000; i++ {
		subCent := i % 100
		superCent := i / 100
		if i == 1000 {
			charCount += 11 // "onethousand"
		} else if superCent > 0 {
			charCount += digitCount(superCent) + 7 // "hundred"
		}
		if superCent > 0 && subCent > 0 {
			charCount += 3 // "and"
		}
		if subCent > 0 {
			charCount += subCentCount(subCent)
		}
	}
	return charCount
}

func digitString(num int) string {
	str := ""
	switch num {
	case 1:
		str = "one"
	case 2:
		str = "two"
	case 3:
		str = "three"
	case 4:
		str = "four"
	case 5:
		str = "five"
	case 6:
		str = "six"
	case 7:
		str = "seven"
	case 8:
		str = "eight"
	case 9:
		str = "nine"
	}
	return str
}

func subCentString(num int) string {
	str := ""
	if num/10 == 1 {
		switch num {
		case 10:
			str = "ten"
		case 11:
			str = "eleven"
		case 12:
			str = "twelve"
		case 13:
			str = "thirteen"
		case 14:
			str = "fourteen"
		case 15:
			str = "fifteen"
		case 16:
			str = "sixteen"
		case 17:
			str = "seventeen"
		case 18:
			str = "eighteen"
		case 19:
			str = "nineteen"
		}
	} else {
		if num >= 20 {
			switch num / 10 {
			case 2:
				str = "twenty"
			case 3:
				str = "thirty"
			case 4:
				str = "forty"
			case 5:
				str = "fifty"
			case 6:
				str = "sixty"
			case 7:
				str = "seventy"
			case 8:
				str = "eighty"
			case 9:
				str = "ninety"
			}
		}
		if num%10 > 0 {
			str = str + digitString(num%10)
		}
	}
	return str
}

func digitCount(num int) int {
	cnt := 0
	switch num {
	case 1:
		cnt = 3 // "one"
	case 2:
		cnt = 3 // "two"
	case 3:
		cnt = 5 // "three"
	case 4:
		cnt = 4 // "four"
	case 5:
		cnt = 4 // "five"
	case 6:
		cnt = 3 // "six"
	case 7:
		cnt = 5 // "seven"
	case 8:
		cnt = 5 // "eight"
	case 9:
		cnt = 4 // "nine"
	}
	return cnt
}

func subCentCount(num int) int {
	cnt := 0
	if num/10 == 1 {
		switch num {
		case 10:
			cnt = 3 // "ten"
		case 11:
			cnt = 6 // "eleven"
		case 12:
			cnt = 6 // "twelve"
		case 13:
			cnt = 8 // "thirteen"
		case 14:
			cnt = 8 // "fourteen"
		case 15:
			cnt = 7 // "fifteen"
		case 16:
			cnt = 7 // "sixteen"
		case 17:
			cnt = 9 // "seventeen"
		case 18:
			cnt = 8 // "eighteen"
		case 19:
			cnt = 8 // "nineteen"
		}
	} else {
		if num >= 20 {
			switch num / 10 {
			case 2:
				cnt = 6 // "twenty"
			case 3:
				cnt = 6 // "thirty"
			case 4:
				cnt = 5 // "forty"
			case 5:
				cnt = 5 // "fifty"
			case 6:
				cnt = 5 // "sixty"
			case 7:
				cnt = 7 // "seventy"
			case 8:
				cnt = 6 // "eighty"
			case 9:
				cnt = 6 // "ninety"
			}
		}
		if num%10 > 0 {
			cnt = cnt + digitCount(num%10)
		}
	}
	return cnt
}
