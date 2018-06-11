package auth

import (
	"time"
	"log"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"net/http"
)

type login struct {
	id string
	login string
	pass string
}
func Login(w http.ResponseWriter, r *http.Request) {

	var newLogin *login=new(login)

	r.ParseForm()
	timeNow := time.Now()
	newLogin.login = r.Form.Get("Login")
	newLogin.pass=r.Form.Get("Pass")

	if atoiErr != nil {
		log.Print(atoiErr, " ERROR: Wrong user ID (can't convert string to int)")
		common.SendError(w,r,400,"ERROR: Wrong user ID (can't convert string to int)",atoiErr)
		return
	}

	newTask.User_id = userId
	newTask.Name = r.Form.Get("name")
	newTime := r.Form.Get("time")
	newTask.Created_at = timeNow
	newTask.Updated_at = timeNow //When we create new task updated time = created time
	newTask.Desc = r.Form.Get("desc")

	parsedTime, parsedErr := time.Parse(time.UnixDate, newTime)

	if parsedErr != nil {
		log.Print(parsedErr, " ERROR: Wrong date (can't convert string to int)")
		common.SendError(w,r,415,"ERROR: Wrong date(can't convert string to int)",parsedErr)
		return
	}

	newTask.Time = parsedTime

	task, taskErr := testAddTask(newTask)

	if taskErr != nil {
		log.Print(taskErr, " ERROR: Can't add new task")
		common.SendError(w,r,400,"ERROR: Can't add new task",taskErr)
		return
	}

	common.RenderJSON(w, r, task)
}

func tryLogin (loginUser login){
	userInDB:=db.search("login",loginUser.login)
	if  userData ==userData.pass{
		redis.CreateSession(userInDB.id)
		return userInDB.id
	}
}

func Logout(logoutUser login)  {
	redis.FinishSession(logoutUser.id)
}