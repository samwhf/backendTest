package tests

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/samwhf/backendTest/api"
	"github.com/samwhf/backendTest/common/responder"
	"github.com/samwhf/backendTest/objects"
	service "github.com/samwhf/backendTest/services/user"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	testEnvInit()
	Convey("succeed", t, func() {
		guards := []*monkey.PatchGuard{
			monkey.Patch(service.Get, getSuc),
		}
		defer bulkUnpatch(guards)
		Convey("get", func() {
			for _, req := range getUserTestRequest() {
				// mock req.xxx, ps:req.headers.Set(k, v)
				w := httptest.NewRecorder()
				testRouter.ServeHTTP(w, req)
				// t.Log(w.Code) 200
				// t.Log(w.Body.String()) {"status":0,"success":true,"message":"Resource is fetched successfully","data":{"id":"id_1","name":"test","dob":"0001-01-01T00:00:00Z","address":"","description":"","createdAt":"0001-01-01T00:00:00Z"}}
				// So(w.Result().StatusCode, ShouldEqual, http.StatusOK)
				body := &responder.ResponseBody{}
				json.NewDecoder(w.Body).Decode(body)
				So(body.Success, ShouldEqual, true)
			}
		})
	})
	Convey("error", t, func() {
		guards := []*monkey.PatchGuard{
			monkey.Patch(service.Get, getFail),
		}
		defer bulkUnpatch(guards)
		Convey("service error", func() {
			for _, req := range getUserTestRequest() {
				w := httptest.NewRecorder()
				testRouter.ServeHTTP(w, req)
				// So(w.Result().StatusCode, ShouldEqual, http.StatusOK)
				body := &responder.ResponseBody{}
				json.NewDecoder(w.Body).Decode(body)
				So(body.Success, ShouldEqual, false)
			}
		})
	})
}

func TestPost(t *testing.T) {
	testEnvInit()
	Convey("succeed", t, func() {
		guards := []*monkey.PatchGuard{
			monkey.Patch(service.Create, createSuc),
		}
		defer bulkUnpatch(guards)
		Convey("post", func() {
			for _, req := range postUserTestRequest() {
				w := httptest.NewRecorder()
				testRouter.ServeHTTP(w, req)
				body := &responder.ResponseBody{}
				json.NewDecoder(w.Body).Decode(body)
				So(body.Success, ShouldEqual, true)
			}
		})
	})
	Convey("error", t, func() {
		guards := []*monkey.PatchGuard{
			monkey.Patch(service.Create, createFail),
		}
		defer bulkUnpatch(guards)
		Convey("service error", func() {
			for _, req := range postUserTestRequest() {
				w := httptest.NewRecorder()
				testRouter.ServeHTTP(w, req)
				body := &responder.ResponseBody{}
				json.NewDecoder(w.Body).Decode(body)
				So(body.Success, ShouldEqual, false)
			}
		})
	})
}

func TestUpdate(t *testing.T) {
	testEnvInit()
	Convey("succeed", t, func() {
		guards := []*monkey.PatchGuard{
			monkey.Patch(service.Update, updateSuc),
		}
		defer bulkUnpatch(guards)
		Convey("update", func() {
			for _, req := range updateUserTestRequest() {
				w := httptest.NewRecorder()
				testRouter.ServeHTTP(w, req)
				body := &responder.ResponseBody{}
				json.NewDecoder(w.Body).Decode(body)
				So(body.Success, ShouldEqual, true)
			}
		})
	})
	Convey("error", t, func() {
		guards := []*monkey.PatchGuard{
			monkey.Patch(service.Update, updateFail),
		}
		defer bulkUnpatch(guards)
		Convey("service error", func() {
			for _, req := range updateUserTestRequest() {
				w := httptest.NewRecorder()
				testRouter.ServeHTTP(w, req)
				body := &responder.ResponseBody{}
				json.NewDecoder(w.Body).Decode(body)
				So(body.Success, ShouldEqual, false)
			}
		})
	})
}

func TestDelete(t *testing.T) {
	testEnvInit()
	Convey("succeed", t, func() {
		guards := []*monkey.PatchGuard{
			monkey.Patch(service.Delete, deleteSuc),
		}
		defer bulkUnpatch(guards)
		Convey("delete", func() {
			for _, req := range deleteUserTestRequest() {
				w := httptest.NewRecorder()
				testRouter.ServeHTTP(w, req)
				body := &responder.ResponseBody{}
				json.NewDecoder(w.Body).Decode(body)
				So(body.Success, ShouldEqual, true)
			}
		})
	})
	Convey("error", t, func() {
		guards := []*monkey.PatchGuard{
			monkey.Patch(service.Delete, deleteFail),
		}
		defer bulkUnpatch(guards)
		Convey("service error", func() {
			for _, req := range deleteUserTestRequest() {
				w := httptest.NewRecorder()
				testRouter.ServeHTTP(w, req)
				body := &responder.ResponseBody{}
				json.NewDecoder(w.Body).Decode(body)
				So(body.Success, ShouldEqual, false)
			}
		})
	})
}

// user get接口方法的测试用例
func getUserTestRequest() (testReqs []*http.Request) {
	prefix := "/api/v1/user"
	uris := []string{
		"/id_1",
	}
	for _, uri := range uris {
		request := httptest.NewRequest(http.MethodGet,
			fmt.Sprintf("http://example.com%s%s", prefix, uri), nil)
		testReqs = append(testReqs, request)
	}
	return testReqs
}

// user post接口方法的测试用例
func postUserTestRequest() (testReqs []*http.Request) {
	prefix := "/api/v1/user"
	request := httptest.NewRequest(http.MethodPost,
		fmt.Sprintf("http://example.com%s", prefix), strings.NewReader("{}"))
	request.Header.Set("Content-Type", "application/json")
	testReqs = append(testReqs, request)
	return testReqs
}

// user update接口方法的测试用例
func updateUserTestRequest() (testReqs []*http.Request) {
	prefix := "/api/v1/user"
	uris := []string{
		"/id_1",
	}
	for _, uri := range uris {
		request := httptest.NewRequest(http.MethodPut,
			fmt.Sprintf("http://example.com%s%s", prefix, uri), strings.NewReader("{}"))
		request.Header.Set("Content-Type", "application/json")
		testReqs = append(testReqs, request)
	}
	return testReqs
}

// user delete接口方法的测试用例
func deleteUserTestRequest() (testReqs []*http.Request) {
	prefix := "/api/v1/user"
	uris := []string{
		"/id_1",
	}
	for _, uri := range uris {
		request := httptest.NewRequest(http.MethodDelete,
			fmt.Sprintf("http://example.com%s%s", prefix, uri), strings.NewReader("{}"))
		request.Header.Set("Content-Type", "application/json")
		testReqs = append(testReqs, request)
	}
	return testReqs
}

var (
	testEnvInitOnce = new(sync.Once)
	testRouter      *gin.Engine
	monkeyErr       = errors.New("test error")
	getSuc          = func(_ context.Context, _ string) (*objects.User, error) {
		return &objects.User{
			ID:   "id_1",
			Name: "test",
		}, nil
	}
	getFail = func(_ context.Context, _ string) (*objects.User, error) {
		return nil, errors.New("record not found")
	}
	createSuc = func(_ context.Context, _ *objects.User) (string, error) {
		return "id_1", nil
	}
	createFail = func(_ context.Context, _ *objects.User) (string, error) {
		return "", monkeyErr
	}
	updateSuc = func(_ context.Context, _ *objects.User) error {
		return nil
	}
	updateFail = func(_ context.Context, _ *objects.User) error {
		return monkeyErr
	}
	deleteSuc = func(_ context.Context, id string) error {
		return nil
	}
	deleteFail = func(_ context.Context, id string) error {
		return monkeyErr
	}
)

func testEnvInit() {
	testEnvInitOnce.Do(func() {
		// init http router
		gin.SetMode(gin.TestMode)
		testRouter = gin.New()
		api.SetUpRoutes(testRouter)
	})
}

func bulkUnpatch(patches []*monkey.PatchGuard) {
	for _, patch := range patches {
		patch.Unpatch()
	}
}
