package usersCRUD

import (
	"net/http"
	"github.com/gorilla/mux"
	"regexp"
	"log"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/satori/go.uuid"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
)

type succesMessage struct {
	Status string `json:"status"`
}

var tempID, _ = uuid.FromString("00000000-0000-0000-0000-000000000001")

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		log.Print(err, " ERROR: Can't get users")
		common.SendError(w, r, 404, "ERROR: Can't get users", err)
		return
	}
	common.RenderJSON(w, r, users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUser, err := uuid.FromString(params["id"])
	if err != nil {
		log.Print(err, "ERROR: Converting ID from URL")
		common.SendError(w, r, 400, "ERROR: Converting ID from URL", err)
		return
	}
	user, err := models.GetUser(idUser)
	if err != nil {
		log.Print(err, "ERROR: Can't find user with such ID")
		common.SendError(w, r, 404, "ERROR: Can't find user with such ID", err)
		return
	}
	common.RenderJSON(w, r, user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err, " ERROR: Can't parse POST Body")
		common.SendError(w, r, 400, "ERROR: Can't parse POST Body", err)
		return
	}
	var newUser models.User
	newUser.ID = tempID
	newUser.Login = r.Form.Get("login")
	newUser.Name = r.Form.Get("name")
	newUser.Password = r.Form.Get("password")
	valid, errMessage := IsValid(newUser)
	if !valid {
		log.Print(errMessage)
	}
	_, err = models.CreateUser(newUser)
	if err != nil {
		log.Print(err, " ERROR: Can't add this user")
		common.SendError(w, r, 400, "ERROR: Can't add this user", err)
		return
	}
	common.RenderJSON(w, r, succesMessage{Status: "success"})
}

func IsValid(user models.User) (bool, string) {
	errMessage := ""
	var checkPass = regexp.MustCompile(`^[[:graph:]]*$`)
	var checkName = regexp.MustCompile(`^[A-Z]{1}[a-z]+$`)
	var checkLogin = regexp.MustCompile(`^[[:graph:]]*$`)
	var validPass, validName, validLogin bool
	if len(user.Password) >= 8 && checkPass.MatchString(user.Password) {
		validPass = true
	} else {
		errMessage += "Invalid Password"
	}
	if checkName.MatchString(user.Name) && len(user.Name) < 15 {
		validName = true
	} else {
		errMessage += " Invalid Name"
	}
	if checkLogin.MatchString(user.Login) && len(user.Login) < 15 {
		validLogin = true
	} else {
		errMessage += " Invalid Login"
	}
	return validName && validLogin && validPass, errMessage
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUser, err := uuid.FromString(params["id"])
	if err != nil {
		log.Print(err, " ERROR: Wrong user ID (can't convert string to int)")
		common.SendError(w, r, 400, "ERROR: Wrong user ID (can't convert string to int)", err)
		return
	}
	err = models.DeleteUser(idUser)
	if err != nil {
		log.Print(err, " ERROR: Can't delete this user")
		common.SendError(w, r, 404, "ERROR: Can't delete this user", err)
		return
	}
	common.RenderJSON(w, r, succesMessage{Status: "success"})
}
