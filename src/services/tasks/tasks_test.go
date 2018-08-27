package tasks_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"

	"encoding/json"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/auth"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/tasks"
	"github.com/satori/go.uuid"
)

var router = services.NewRouter()
var testUUID = "00000000-0000-0000-0000-000000000001"

type tasksCRUDTestCase struct {
	name             string
	url              string
	want             int
	mockedGetTask    models.Task
	mockedGetTasks   []models.Task
	mockedDeleteTask uuid.UUID
	mockedCreateTask models.Task
	mockedTasksError error
	permission       bool
	mock             func()
	userId           string
	testTime         string
}

func TestGetTasks(t *testing.T) {

	tests := []tasksCRUDTestCase{
		{
			name:             "Get_Tasks_200",
			url:              "/v1/tasks",
			want:             200,
			mockedGetTasks:   []models.Task{},
			mockedTasksError: nil,
			permission:       true,
		},
		{
			name:             "Get_Tasks_404",
			url:              "/v1/tasks",
			want:             404,
			mockedGetTasks:   []models.Task{},
			mockedTasksError: http.ErrBodyNotAllowed,
			permission:       true,
		},
		{
			name:             "Get_Tasks_403",
			url:              "/v1/tasks",
			want:             403,
			mockedGetTasks:   []models.Task{},
			mockedTasksError: http.ErrBodyNotAllowed,
			permission:       false,
		},
	}
	auth.MockedGetSession("", nil)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			auth.MockedCheckPermission(tc.permission)
			models.MockedGetTasks(tc.mockedGetTasks, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			tasks.GetTasksHandler(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestGetTasksByID(t *testing.T) {
	tests := []tasksCRUDTestCase{
		{
			name: "Get_TaskById_200",
			url:  "/v1/tasks/" + testUUID,
			want: 200,
		},
		{
			name: "Get_TaskByWithInvalidTaskID",
			url:  "/v1/tasks/asdasd",
			want: 400,
		},
		{
			name:             "Get_TaskByWithWrongTaskID",
			url:              "/v1/tasks/00000000-0000-0000-0000-000000000002",
			want:             404,
			mockedTasksError: http.ErrLineTooLong,
		},
	}
	auth.MockedGetSession("", nil)
	auth.MockedCheckPermission(true)
	models.MockedGetUserByID(models.User{}, nil)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			models.MockedGetTask(tc.mockedGetTask, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestDeleteTasks(t *testing.T) {
	id, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	tests := []tasksCRUDTestCase{
		{
			name: "Delete_Task_204",
			url:  "/v1/tasks/" + testUUID,
			want: 204,
		},
		{
			name:             "Delete_Task_400",
			url:              "/v1/tasks/asda",
			want:             400,
			mockedDeleteTask: id,
			mockedTasksError: nil,
			mock: func() {
			},
		},
		{
			name:             "Delete_Task_404",
			url:              "/v1/tasks/00000000-0000-0000-0000-000000000001",
			want:             404,
			mockedDeleteTask: id,
			mockedTasksError: http.ErrAbortHandler,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedDeleteTask(tc.mockedDeleteTask, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodDelete, tc.url, nil)

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestCreateTasks(t *testing.T) {
	tests := []tasksCRUDTestCase{
		{
			name:     "Add_Task_201",
			url:      "/v1/tasks",
			want:     201,
			userId:   "00000000-0000-0000-0000-000000000001",
			testTime: "Mon Jan 2 15:04:05 MST 2006",
		},
		{
			name:     "Add_Task_WrongID",
			url:      "/v1/tasks",
			want:     400,
			userId:   "sadsadsad",
			testTime: "Mon Jan 2 15:04:05 MST 2006",
		},

		{name: "Add_Task_WrongTime",
			url:      "/v1/tasks",
			want:     415,
			userId:   "00000000-0000-0000-0000-000000000001",
			testTime: "sdadasd",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var data tasks.JsonTask
			data.UserID = tc.userId
			data.Name = "JustUser"
			data.Time = tc.testTime
			data.Desc = "Desc of my task"
			body, _ := json.Marshal(data)

			models.MockedCreateTask(tc.mockedCreateTask, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestGetUserTasks(t *testing.T) {
	auth.MockedCheckPermission(true)
	tests := []tasksCRUDTestCase{
		{
			name:             "GetUserTasks",
			url:              "/v1/users/tasks/00000000-0000-0000-0000-000000000001",
			want:             200,
			mockedTasksError: nil,
			mockedGetTasks:   []models.Task{},
		},
		{
			name:             "GetUserTasks",
			url:              "/v1/users/tasks/00000000-0000-0000-0000-000000000001",
			want:             404,
			mockedTasksError: http.ErrLineTooLong,
			mockedGetTasks:   []models.Task{},
		},
		{
			name:             "GetUserTasks",
			url:              "/v1/users/tasks/wrongid",
			want:             404,
			mockedTasksError: http.ErrLineTooLong,
			mockedGetTasks:   []models.Task{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
			models.MockedGetUserTasks(tc.mockedGetTasks, tc.mockedTasksError)
			router.ServeHTTP(rec, req)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
