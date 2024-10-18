package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	mocks "github.com/mmfshirokan/apod/internal/handlers/mock"
	"github.com/mmfshirokan/apod/internal/model"
)

var (
	testProxyAddr = "http://proxy.test/"
	testJSON      = `{"copyright":"` + testModel.Copyright + `","date":"` + testModel.Date + `","explanation":"` + testModel.Explanation + `","hdurl":"` + testModel.UrlHD + `","media_type":"` + testModel.MediaType + `","title":"","url":"` + testModel.Url + `","service_version":"` + testModel.ServiceVersion + `","proxy_url":"` + testModel.Url + `","proxy_url":"` + testProxyAddr + testModel.Date + `.jpg"}`
	testModel     = model.ImageInfo{
		Copyright:      "\nBrennan Gilmore\n",
		Date:           "2024-10-14",
		Explanation:    "Go outside at sunset tonight and see a comet!  C/2023 A3 (Tsuchinshan–ATLAS) has become visible...",
		UrlHD:          "https://apod.nasa.gov/apod/image/2410/CometA3Dc_Gilmore_1080.jpg",
		MediaType:      "image",
		ServiceVersion: "v1",
		Title:          "Comet Tsuchinshan-ATLAS Over the Lincoln Memorial",
		Url:            "https://apod.nasa.gov/apod/image/2410/CometA3Dc_Gilmore_1080.jpg",
	}
)

func TestGet(t *testing.T) {
	testMethod := http.MethodPut
	date := "2024-01-01"
	testTarget := "/get/" + date

	geter := mocks.NewGeter(t)

	rec := httptest.NewRecorder()
	hnd := New(geter, testProxyAddr)
	req := httptest.NewRequest(testMethod, testTarget, strings.NewReader(testJSON))
	param := httprouter.Params{{Key: "date", Value: date}}

	call := geter.EXPECT().Get(req.Context(), date).Return(testModel, nil)

	hnd.Get(rec, req, param)

	call.Parent.AssertExpectations(t)

}

func TestGetAll(t *testing.T) {
	testMethod := http.MethodPut
	date := "2024-01-01"
	testTarget := "/get/" + date

	geter := mocks.NewGeter(t)

	rec := httptest.NewRecorder()
	hnd := New(geter, testProxyAddr)
	req := httptest.NewRequest(testMethod, testTarget, strings.NewReader(testJSON))
	param := httprouter.Params{{Key: "date", Value: date}}

	call := geter.EXPECT().GetAll(req.Context()).Return([]model.ImageInfo{testModel}, nil)

	hnd.GetAll(rec, req, param)

	call.Parent.AssertExpectations(t)

}
