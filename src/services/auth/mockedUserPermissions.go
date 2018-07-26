package auth

import "net/http"

func MockedCheckPermission(res bool){
	CheckPermission= func(r *http.Request, requiredRole string, itemOwner string) bool {
		return res
	}
}
