package theme

import (
	"errors"
	"image/color"
	"strconv"
)

type Property struct {
	Name  string
	Value string
}

var errInvalidFormat = errors.New("invalid format")

func (prop Property) AsPX() (int, error) {
	strLen := len(prop.Value)
	if prop.Value[strLen-2:] != "px" {
		return 0, errors.New("value is not a pixel value")
	}

	pxStr := prop.Value[:strLen-2]

	i, err := strconv.Atoi(pxStr)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (prop Property) AsColor() (c color.RGBA, err error) {
	s := prop.Value
	c.A = 0xff

	if len(s) == 0 {
		return c, errInvalidFormat
	}

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 9:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
		c.A = hexToByte(s[7])<<4 + hexToByte(s[8])
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
