package auth

func MockedCheckPermission(res bool) {
	CheckPermission = func(userSession string, requiredRole string, itemOwner string) bool {
		return res
	}
}
