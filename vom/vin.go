package vin

import "fmt"

type VIN string

func (v VIN) Manufacturer() string {
	manufacturer := v[:3]
	if manufacturer[2] == '9' {
		manufacturer += v[11:14]
	}
	return string(manufacturer)
}
func NewVIN(c string) (VIN, error) {
	if len(c) != 17 {
		return "", fmt.Errorf("invalid VIN %s: more or less than 17 characters", c)
	}
	return VIN(c), nil
}

