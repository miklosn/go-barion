// Code generated by "enumer -type=PaymentStatus -json"; DO NOT EDIT.

//
package barion

import (
	"encoding/json"
	"fmt"
)

const _PaymentStatusName = "PreparedStartedInProgressWaitingReservedAuthorizedCanceledSucceededFailedPartiallySucceededExpired"

var _PaymentStatusIndex = [...]uint8{0, 8, 15, 25, 32, 40, 50, 58, 67, 73, 91, 98}

func (i PaymentStatus) String() string {
	if i < 0 || i >= PaymentStatus(len(_PaymentStatusIndex)-1) {
		return fmt.Sprintf("PaymentStatus(%d)", i)
	}
	return _PaymentStatusName[_PaymentStatusIndex[i]:_PaymentStatusIndex[i+1]]
}

var _PaymentStatusValues = []PaymentStatus{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

var _PaymentStatusNameToValueMap = map[string]PaymentStatus{
	_PaymentStatusName[0:8]:   0,
	_PaymentStatusName[8:15]:  1,
	_PaymentStatusName[15:25]: 2,
	_PaymentStatusName[25:32]: 3,
	_PaymentStatusName[32:40]: 4,
	_PaymentStatusName[40:50]: 5,
	_PaymentStatusName[50:58]: 6,
	_PaymentStatusName[58:67]: 7,
	_PaymentStatusName[67:73]: 8,
	_PaymentStatusName[73:91]: 9,
	_PaymentStatusName[91:98]: 10,
}

// PaymentStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PaymentStatusString(s string) (PaymentStatus, error) {
	if val, ok := _PaymentStatusNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PaymentStatus values", s)
}

// PaymentStatusValues returns all values of the enum
func PaymentStatusValues() []PaymentStatus {
	return _PaymentStatusValues
}

// IsAPaymentStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PaymentStatus) IsAPaymentStatus() bool {
	for _, v := range _PaymentStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for PaymentStatus
func (i PaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentStatus
func (i *PaymentStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PaymentStatus should be a string, got %s", data)
	}

	var err error
	*i, err = PaymentStatusString(s)
	return err
}
