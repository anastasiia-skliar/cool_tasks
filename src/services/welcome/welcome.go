package welcome

import (
	"net/http"

	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
)

type welcomeStruct struct {
	Message string `json:"message"`
}

// GetWelcomeHandler get policies for partner
func GetWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	common.RenderJSON(w, r, welcomeStruct{Message: "Hello World"})

}
