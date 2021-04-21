// +build !windows

package kerberos

import "errors"

const haveSSPI = false

func sspiAuthHeader(serviceName string) (string, error) {
	return "", errors.New("platform does not support Kerberos SSPI")
}
