package usersCRUD

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"regexp"
	"log"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
)

type Users struct {
	Id int
	Login string
	Password string
	Name string
}

var users=[]Users {{1,"Karim","1234qwer","Karim"}}
var user=Users{1,"Karim","1234qwer","Karim"}

func SendUsers() ([]Users, error){
	return users, nil
}

func ByID (id int)(user Users, err error) {
	user=Users{1,"Karim","1234qwer","Karim"}
	return user, err
}

func Add (user Users) error{
	return nil
}

func Delete(id int) error {
	return nil
}

func GetUsers (w http.ResponseWriter, r *http.Request) {
	users, err:=SendUsers()
	if err!=nil {
		log.Print(err, " ERROR: Can't get users")
		common.SendError(w, r, 404, "ERROR: Can't get users", err)
		return
	}
	common.RenderJSON(w, r, users)
}

func GetUserByID (w http.ResponseWriter, r*http.Request) {
	params:=mux.Vars(r)
	idUser,readError:=strconv.Atoi(params["id"])
	if readError!=nil {
		log.Print(readError, "ERROR: Converting ID from URL")
		common.SendError(w, r, 400, "ERROR: Converting ID from URL", readError)
		return
	}
	user, err:=ByID(idUser)
	if err!=nil {
		log.Print(err, "ERROR: Can't find user with such ID")
		common.SendError(w, r, 404, "ERROR: Can't find user with such ID", err)
		return
	}
	common.RenderJSON(w, r, user)
}

func AddUser (w http.ResponseWriter, r*http.Request) {
	parsingError:=r.ParseForm()
	if parsingError!=nil {
		log.Print(parsingError, " ERROR: Can't parse POST Body")
		common.SendError(w, r, 400, "ERROR: Can't parse POST Body", parsingError)
		return
	}
	var newUser Users
	newUser.Id=0
	newUser.Login = r.Form.Get("login")
	newUser.Name = r.Form.Get("name")
	newUser.Password = r.Form.Get("password")
	valid, err:=IsValid(newUser)
	if !valid {
		log.Print(err)
	}
	addingError:=Add(newUser)
	if addingError!=nil {
		log.Print(addingError, " ERROR: Can't add this user")
		common.SendError(w, r, 400, "ERROR: Can't add this user", addingError)
		return
	}
}

func IsValid (user Users) (bool, string) {
	err:=""
	var checkPass = regexp.MustCompile(`^[[:graph:]]*$`)
	var checkName = regexp.MustCompile(`^[A-Z]{1}[a-z]+$`)
	var checkLogin = regexp.MustCompile(`^[[:graph:]]*$`)
	var validPass, validName, validLogin bool
	if len(user.Password)>=8 && checkPass.MatchString(user.Password) {
		validPass = true
	} else {
		err+="Invalid Password"
	}
	if checkName.MatchString(user.Name) && len(user.Name)<15 {
		validName = true
	} else {
		err+= " Invalid Name"
	}
	if checkLogin.MatchString(user.Login) && len(user.Login)<15 {
		validLogin = true
	} else {
		err+=" Invalid Login"
	}
	return validName && validLogin && validPass, err
}

func DeleteUser(w http.ResponseWriter, r*http.Request) {
	params:=mux.Vars(r)
	idUser,readError:=strconv.Atoi(params["id"])
	if readError!=nil {
		log.Print(readError, " ERROR: Wrong user ID (can't convert string to int)")
		common.SendError(w, r, 400, "ERROR: Wrong user ID (can't convert string to int)", readError)
		return
	}
	err:=Delete(idUser)
	if err!=nil {
		log.Print(err, " ERROR: Can't delete this user")
		common.SendError(w, r, 404, "ERROR: Can't delete this user", err)
		return
	}
}



//func main () {
//	r:= mux.NewRouter()
//	r.HandleFunc("/users", GetUsers).Methods("GET")
//	r.HandleFunc("/users/{id}", GetUserByID).Methods("GET")
//	r.HandleFunc("/users", AddUser).Methods("POST")
//	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
//	http.ListenAndServe(":8080", r)
//	//fmt.Println(IsValid(user))
//
//}
