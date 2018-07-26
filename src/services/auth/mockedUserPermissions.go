package auth

import "net/http"

func MockedCheckPermissionTrue(r *http.Request, requiredRole string, itemOwner string) bool {
	return true
}

func MockedCheckPermissionFalse(r *http.Request, requiredRole string, itemOwner string) bool {
	return false
}
