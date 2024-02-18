package api

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/rand"
	"io/ioutil"
	mockdb "myclass/db/mock"
	db "myclass/db/sqlc"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateProfessionalUserAPI(t *testing.T) {
	//create a new controller
	professionalUser := randomProfessionalUserCreate()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//create a mock object for the Store interface
	mockStore := mockdb.NewMockStore(ctrl)

	req := db.CreateProfessionalUserParams{
		Name:     "Monica",
		Username: "monica",
		Password: "passwordMonica",
		Gender:   "female",
		Email:    "monica@gmail.com",
		DateOfBirth: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		Cpf:            123654,
		ImageID:        2,
		ClassHourPrice: "20",
	}

	//set expectations
	mockStore.EXPECT().
		CreateProfessionalUser(gomock.Any(), gomock.Any()).AnyTimes().
		DoAndReturn(func(ctx context.Context, arg db.CreateProfessionalUserParams) (db.ProfessionalUser, error) {
			assert.NotZero(t, arg.UpdatedAt)
			return professionalUser, nil

		})

	//start test server and send request
	server := NewServer(mockStore)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/professionalUser")
	body := getBodyReader(req)
	request, err := http.NewRequest(http.MethodPost, url, &body)

	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyProfessionalUser(t, recorder.Body, professionalUser)
}

func TestCreateProfessionalUserAPIBadRequest(t *testing.T) {
	//create a new controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//create a mock object for the Store interface
	mockStore := mockdb.NewMockStore(ctrl)

	req := gin.H{
		"Name":     "Monica",
		"Password": "passwordMonica",
		"Gender":   "female",
		"Email":    "monica@gmail.com",
		"DateOfBirth": time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		"Cpf":            8686757565,
		"ImageID":        2,
		"ClassHourPrice": "20",
	}

	//set expectations
	mockStore.EXPECT().
		CreateProfessionalUser(gomock.Any(), gomock.Any()).AnyTimes().
		DoAndReturn(func(ctx context.Context, arg db.CreateProfessionalUserParams) (db.ProfessionalUser, error) {
			return db.ProfessionalUser{}, errors.New("Bad Request")
		})

	//start test server and send request
	server := NewServer(mockStore)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/professionalUser")
	body := getBodyReader(req)
	request, err := http.NewRequest(http.MethodPost, url, &body)

	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestGetProfessionalUserAPI(t *testing.T) {
	professionalUser := randomProfessionalUser()

	testCases := []struct {
		name               string
		professionalUserId int64
		buildStubs         func(store *mockdb.MockStore)
		checkResponse      func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:               "Successful getting professional",
			professionalUserId: professionalUser.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProfessionalUser(gomock.Any(), gomock.Eq(professionalUser.ID)).AnyTimes().
					Return(professionalUser, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyProfessionalUser(t, recorder.Body, professionalUser)
			},
		},
		{
			name:               "Not found getting professional",
			professionalUserId: professionalUser.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProfessionalUser(gomock.Any(), gomock.Eq(professionalUser.ID)).AnyTimes().
					Return(db.ProfessionalUser{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		//{
		//	name:               "InternalServerError getting professional",
		//	professionalUserId: professionalUser.ID,
		//	buildStubs: func(store *mockdb.MockStore) {
		//		store.EXPECT().
		//			GetProfessionalUser(gomock.Any(), gomock.Eq(professionalUser.ID)).AnyTimes().
		//			Return(db.ProfessionalUser{}, sql.ErrConnDone)
		//	},
		//	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusInternalServerError, recorder.Code)
		//	},
		//	},
		{
			name:               "BadRequest  Invalid Id getting professional",
			professionalUserId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetProfessionalUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {

		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			testCase.buildStubs(store)

			//start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/professionalUser/%d", testCase.professionalUserId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(t, recorder)
		})
	}
}

func randomProfessionalUser() db.ProfessionalUser {
	return db.ProfessionalUser{
		ID: 1,
		CreatedAt: time.Date(
			2024, 01, 17, 20, 34, 58, 651387237, time.UTC),
		Name:     "Monica",
		Username: "monica",
		Password: "passwordMonica",
		Gender:   "female",
		Email:    "monica@gmail.com",
		DateOfBirth: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		Cpf:     123654,
		ImageID: 2,
		UpdatedAt: time.Date(
			2024, 01, 17, 20, 34, 58, 651387237, time.UTC),
		ClassHourPrice: "20",
	}
}

func randomProfessionalUserCreate() db.ProfessionalUser {
	return db.ProfessionalUser{
		ID: randomID(10, 40),
		CreatedAt: time.Date(
			2024, 01, 17, 20, 34, 58, 651387237, time.UTC),
		Name:     "Monica",
		Username: "monica",
		Password: "passwordMonica",
		Gender:   "female",
		Email:    "monica@gmail.com",
		DateOfBirth: time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
		Cpf:     123654,
		ImageID: 2,
		UpdatedAt: time.Date(
			2024, 01, 17, 20, 34, 58, 651387237, time.UTC),
		ClassHourPrice: "20",
	}
}

func requireBodyProfessionalUser(t *testing.T, body *bytes.Buffer, professionalUser db.ProfessionalUser) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotProfessionalUser db.ProfessionalUser
	err = json.Unmarshal(data, &gotProfessionalUser)
	require.NoError(t, err)
	require.Equal(t, gotProfessionalUser, professionalUser)
}

func getBodyReader(iface interface{}) bytes.Buffer {
	buffer := new(bytes.Buffer)
	_ = json.NewEncoder(buffer).Encode(iface)
	return *buffer
}

func randomID(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)

}
