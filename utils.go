// Utility functions
package gocaptcha

import "strings"

// Converts boolean into a valid string for request params.
func boolConv(boolean bool) string {
	if boolean {
		return "1"
	}
	return "0"
}

// Checks ERROR in api response
func checkApiResponse(body string) bool {
	return !strings.Contains(body, "ERROR")
}

// Gets the result from api
func parseResult(body string) string {
	return strings.Split(body, "OK|")[1]
}
