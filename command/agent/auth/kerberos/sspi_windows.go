// +build windows

package kerberos

import (
	"encoding/base64"
	"github.com/alexbrainman/sspi/negotiate"
)

const haveSSPI = true

func sspiAuthHeader(serviceName string) (string, error) {
	cred, err := negotiate.AcquireCurrentUserCredentials()
	if err != nil {
		return "", err
	}
	defer cred.Release()

	clientCtx, token, err := negotiate.NewClientContext(cred, serviceName)
	if err != nil {
		return "", err
	}
	defer clientCtx.Release()

	return "Negotiate "+base64.StdEncoding.EncodeToString(token), nil
}
