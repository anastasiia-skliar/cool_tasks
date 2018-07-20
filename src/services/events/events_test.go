package events_test

import (
	"bytes"
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services"
	"github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var router = services.NewRouter()

type EventsTestCase struct {
	name            string
	url             string
	want            int
	mockedGetEvents []models.Event
}

func TestGetByRequestHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name:            "Get_Events_200",
			url:             "/v1/events",
			want:            200,
			mockedGetEvents: []models.Event{},
		},
		{
			name:            "Get_Events_404",
			url:             "/v1/events?mock=890",
			want:            404,
			mockedGetEvents: []models.Event{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetEvents()
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
			if tc.name == "Get_Events_404" {
				mock := func() {
					var err = http.ErrBodyNotAllowed
					models.GetEventsByRequest = func(values url.Values) ([]models.Event, error) {
						return []models.Event{}, err
					}
				}
				mock()
			}
			router.ServeHTTP(rec, req)
			fmt.Println(rec.Code)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}

func TestAddToTripHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name: "Add_Events_200",
			url:  "/v1/events",
			want: 200,
		},
		{
			name: "Add_Events_400",
			url:  "/v1/events",
			want: 400,
		},
		{
			name: "Add_Events_400_2",
			url:  "/v1/events",
			want: 400,
		},
		{
			name: "Add_Events_400_3",
			url:  "/v1/events",
			want: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "Add_Events_400":
				data := url.Values{}
				data.Add("event_id", "asdsad")
				data.Add("trip_id", "00000000-0000-0000-0000-000000000001")
				models.MockedAddEventToTrip()
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				router.ServeHTTP(rec, req)
				fmt.Println(rec.Code)
				if rec.Code != tc.want {
					t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
				}
			case "Add_Events_400_2":
				data := url.Values{}
				data.Add("event_id", "00000000-0000-0000-0000-000000000001")
				data.Add("trip_id", "asdasd")
				models.MockedAddEventToTrip()
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				router.ServeHTTP(rec, req)
				fmt.Println(rec.Code)
				if rec.Code != tc.want {
					t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
				}

			case "Add_Events_400_3":
				data := url.Values{}
				data.Add("event_id", "00000000-0000-0000-0000-000000000001")
				data.Add("trip_id", "00000000-0000-0000-0000-000000000001")
				mock := func() {
					var err = error(new(http.ProtocolError))
					models.AddEventToTrip = func(event_id uuid.UUID, trip_id uuid.UUID) error {
						return err
					}
				}
				mock()
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				router.ServeHTTP(rec, req)
				fmt.Println(rec.Code)
				if rec.Code != tc.want {
					t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
				}

			case "Add_Events_200":
				data := url.Values{}
				data.Add("event_id", "00000000-0000-0000-0000-000000000001")
				data.Add("trip_id", "00000000-0000-0000-0000-000000000001")
				models.MockedAddEventToTrip()
				rec := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodPost, tc.url, bytes.NewBufferString(data.Encode()))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				router.ServeHTTP(rec, req)
				fmt.Println(rec.Code)
				if rec.Code != tc.want {
					t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
				}
			}
		})
	}
}

func TestGetByTripHandler(t *testing.T) {
	tests := []EventsTestCase{
		{
			name:            "Get_Events_200",
			url:             "/v1/events/trip/00000000-0000-0000-0000-000000000001",
			want:            200,
			mockedGetEvents: []models.Event{},
		},
		{
			name:            "Get_Events_200",
			url:             "/v1/events/trip/sadsad",
			want:            400,
			mockedGetEvents: []models.Event{},
		},
		{
			name:            "Get_Events_404",
			url:             "/v1/events/trip/00000000-0000-0000-0000-000000000009",
			want:            404,
			mockedGetEvents: []models.Event{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			models.MockedGetEventsByTrip()
			if tc.name == "Get_Events_404" {
				mock := func() {
					var err = http.ErrLineTooLong
					models.GetEventsByTrip = func(trip_id uuid.UUID) ([]models.Event, error) {
						return []models.Event{}, err
					}
				}
				mock()
			}
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.url, nil)

			router.ServeHTTP(rec, req)
			fmt.Println(rec.Code)
			if rec.Code != tc.want {
				t.Errorf("Expected: %d , got %d", tc.want, rec.Code)
			}
		})
	}
}
