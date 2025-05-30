package mobile

import (
	"errors"
	"regexp"
	"strings"
)

func IsMobileNumberValid(mobileNumber string) bool {
	var (
		exist                    bool
		normalizedNumber, prefix string
	)
	normalizedNumber = cleanMobileNumber(mobileNumber)
	if len(normalizedNumber) != 10 || !strings.HasPrefix(normalizedNumber, "9") {
		return false
	}
	prefix = normalizedNumber[0:3]
	_, exist = operatorPrefixes[prefix]
	return exist
}
func GetMobileNumberOprator(mobileNumber string) (error, string) {
	var (
		normalizedNumber, prefix string
	)
	if !IsMobileNumberValid(mobileNumber) {
		return errors.New("This mobile number is not valid."), ""
	}
	normalizedNumber = cleanMobileNumber(mobileNumber)
	prefix = normalizedNumber[0:3]
	return nil, operatorPrefixes[prefix].OperatorName
}
func GetMobileNumberPersianOprator(mobileNumber string) (error, string) {
	var (
		normalizedNumber, prefix string
	)
	if !IsMobileNumberValid(mobileNumber) {
		return errors.New("This mobile number is not valid."), ""
	}
	normalizedNumber = cleanMobileNumber(mobileNumber)
	prefix = normalizedNumber[0:3]
	return nil, operatorPrefixes[prefix].PersianOperatorName
}
func GetMobileNumberOpratorID(mobileNumber string) (error, int) {
	var (
		normalizedNumber, prefix string
	)
	if !IsMobileNumberValid(mobileNumber) {
		return errors.New("This mobile number is not valid."), 0
	}
	normalizedNumber = cleanMobileNumber(mobileNumber)
	prefix = normalizedNumber[0:3]
	return nil, operatorPrefixes[prefix].operatorID
}
func cleanMobileNumber(mobileNumber string) string {
	var (
		cleanedMobileNumber string
		nonDigitRegex       = regexp.MustCompile(`[^\d]+`)
		prefixRemover       = regexp.MustCompile(`^(0098|\+98|0)`)
	)
	cleanedMobileNumber = nonDigitRegex.ReplaceAllString(mobileNumber, "")
	cleanedMobileNumber = cleanPhoneNumberAggressively(cleanedMobileNumber)
	cleanedMobileNumber = strings.ReplaceAll(cleanedMobileNumber, "-", "")
	cleanedMobileNumber = strings.ReplaceAll(cleanedMobileNumber, "+", "")
	cleanedMobileNumber = strings.ReplaceAll(cleanedMobileNumber, "(", "")
	cleanedMobileNumber = strings.ReplaceAll(cleanedMobileNumber, ")", "")
	cleanedMobileNumber = prefixRemover.ReplaceAllString(cleanedMobileNumber, "")
	return cleanedMobileNumber
}
func cleanPhoneNumberAggressively(phoneNumber string) string {
	var b strings.Builder
	b.Grow(len(phoneNumber)) // Optimize allocation
	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

var operatorPrefixes = map[string]mobileInfo{
	// Hamrahe Aval (MCI) - Iran's first and largest operator
	"910": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"911": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"912": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"913": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"914": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"915": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"916": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"917": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"918": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},
	"919": {operatorID: mciID, OperatorName: mci, PersianOperatorName: mciPersian},

	// Irancell - Iran's second largest operator
	"930": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"933": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"935": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"936": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"937": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"938": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"939": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"901": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"902": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"903": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"904": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"905": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},
	"941": {operatorID: irancellID, OperatorName: irancell, PersianOperatorName: irancellPersian},

	// RighTel - Iran's third operator
	"920": {operatorID: rightelID, OperatorName: rightel, PersianOperatorName: rightelPersian},
	"921": {operatorID: rightelID, OperatorName: rightel, PersianOperatorName: rightelPersian},
	"922": {operatorID: rightelID, OperatorName: rightel, PersianOperatorName: rightelPersian},

	// Other operators / MVNOs (Mobile Virtual Network Operators)
	"990": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"991": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"992": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"993": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"994": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"995": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"996": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"997": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"998": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
	"999": {operatorID: noneID, OperatorName: none, PersianOperatorName: nonePersian},
}

type mobileInfo struct {
	operatorID          int
	OperatorName        string
	PersianOperatorName string
}

const (
	mci             = "mci"
	mciPersian      = "همراه اول"
	mciID           = 1
	irancell        = "irancell"
	irancellPersian = "ایرانسل"
	irancellID      = 2
	rightel         = "rightel"
	rightelPersian  = "رایتل"
	rightelID       = 3
	none            = "none"
	nonePersian     = "مجازی"
	noneID          = 4
)
