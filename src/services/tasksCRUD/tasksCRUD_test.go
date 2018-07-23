package tasksCRUD_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"

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
	mock 			func()
	userId          string
	testTime		string
}

func TestGetTasks(t *testing.T) {

	tests := []tasksCRUDTestCase{
		{
			name:             "Get_Tasks_200",
			url:              "/v1/tasks",
			want:             200,
			mockedGetTasks:   []models.Task{},
			mockedTasksError: nil,
		},
		{
			name:             "Get_Tasks_404",
			url:              "/v1/tasks",
			want:             404,
			mockedGetTasks:   []models.Task{},
			mockedTasksError: http.ErrBodyNotAllowed,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetTasks(tc.mockedGetTasks, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)

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
			name: "Get_TaskByWithWrongTaskID",
			url:  "/v1/tasks/00000000-0000-0000-0000-000000000001",
			want: 404,
			mockedTasksError:http.ErrLineTooLong,
		},
	}

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
			name: "Delete_Task_200",
			url:  "/v1/tasks/" + testUUID,
			want: 200,
		},
		{
			name:              "Delete_Task_400",
			url:               "/v1/tasks/asda",
			want:              400,
			mockedDeleteTask: id,
			mockedTasksError:   nil,
			mock: func() {
			},
		},
		{
			name:              "Delete_Task_404",
			url:               "/v1/tasks/00000000-0000-0000-0000-000000000001",
			want:              404,
			mockedDeleteTask: id,
			mockedTasksError:   http.ErrAbortHandler,
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
			name: "Add_Task_200",
			url:  "/v1/tasks",
			want: 200,
			userId:"00000000-0000-0000-0000-000000000001",
			testTime:"Mon Jan 2 15:04:05 MST 2006",
		},
		{
			name: "Add_Task_WrongID",
			url:  "/v1/tasks",
			want: 400,
			userId:"sadsadsad",
			testTime:"Mon Jan 2 15:04:05 MST 2006",
		},

		{name: "Add_Task_WrongTime",
		url:  "/v1/tasks",
		want: 415,
		userId:"00000000-0000-0000-0000-000000000001",
		testTime: "sdadasd",
	},
	}



	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := url.Values{}
			data.Add("user_id", tc.userId)
			data.Add("name", "JustUser")
			data.Add("time", tc.testTime)
			data.Add("desc", "Desc of my task")
			models.MockedCreateTask(tc.mockedCreateTask, tc.mockedTasksError)
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

			router.ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
			data.Del("user_id")
			data.Del("name")
			data.Del("time")
			data.Del("desc")
		})
	}
}

func TestGetUserTasks(t *testing.T) {
	tests := []tasksCRUDTestCase{
		{
			name: "GetUserTasks",
			url:  "/v1/users/tasks/00000000-0000-0000-0000-000000000001",
			want: 200,
			mockedTasksError:nil,
			mockedGetTasks:[]models.Task{},
		},
		{
			name: "GetUserTasks",
			url:  "/v1/users/tasks/00000000-0000-0000-0000-000000000001",
			want: 404,
			mockedTasksError:http.ErrLineTooLong,
			mockedGetTasks:[]models.Task{},
		},
		{
			name: "GetUserTasks",
			url:  "/v1/users/tasks/wrongid",
			want: 404,
			mockedTasksError:http.ErrLineTooLong,
			mockedGetTasks:[]models.Task{},
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
