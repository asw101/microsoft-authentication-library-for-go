// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package msalbase

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"
)

//ConvertStrUnixToUTCTime converts a string representation of unix time to a UTC timestamp
func ConvertStrUnixToUTCTime(unixTime string) (time.Time, error) {
	timeInt, err := strconv.ParseInt(unixTime, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(timeInt, 0).UTC(), nil
}

//ConcatenateScopes combines all scopes into one space-separated string
func ConcatenateScopes(scopes []string) string {
	return strings.Join(scopes, DefaultScopeSeparator)
}

//SplitScopes splits a space-separated string of scopes to a list
func SplitScopes(scopes string) []string {
	return strings.Split(scopes, DefaultScopeSeparator)
}

//ExtractStringPointerForCache checks a map to see if the key required exists
//If it does, returns a pointer to the string value, if not, returns nil
func ExtractStringPointerForCache(j map[string]interface{}, key string) *string {
	if val, ok := j[key]; ok {
		if str, ok := val.(string); ok {
			delete(j, key)
			return &str
		}
	}
	delete(j, key)
	return nil
}

//GetStringFromPointer checks if a pointer to a string is nil
//If it's not, the pointer is dereferenced; otherwise, an empty string is returned
func GetStringFromPointer(pointer *string) string {
	if pointer == nil {
		return ""
	}
	return *pointer
}

//DecodeJWT decodes a JWT and converts it to a byte array representing a JSON object
//Adapted from MSAL Python and https://stackoverflow.com/a/31971780
func DecodeJWT(data string) ([]byte, error) {
	if i := len(data) % 4; i != 0 {
		data += strings.Repeat("=", 4-i)
	}
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}
