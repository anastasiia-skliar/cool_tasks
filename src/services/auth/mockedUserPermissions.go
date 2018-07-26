package auth

import "net/http"

func mockedCheckPermissionTrue(r *http.Request, requiredRole string, itemOwner string) bool {
	return true
}

func mockedCheckPermissionFalse(r *http.Request, requiredRole string, itemOwner string) bool {
	return false
}