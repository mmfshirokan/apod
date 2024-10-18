package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mmfshirokan/apod/internal/model"
)

type Geter interface {
	Get(ctx context.Context, date string) (model.ImageInfo, error)
	GetAll(ctx context.Context) ([]model.ImageInfo, error)
}

type Handlers struct {
	sv       Geter
	proxyAdr string
}

func New(sv Geter, proxyAdr string) *Handlers {
	return &Handlers{sv: sv, proxyAdr: proxyAdr}
}

// Get godoc
//
// @Summary      Get Image Info
// @Description  Gets an image info on specific date
// @Tags         imageInfo
// @Produce      json
// @Success      200  {object}  model.ImageInfoResponse  "image info"
// @Failure      404  {object}  string  "error message"
// @Router       /get/{date} [get]
func (h *Handlers) Get(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := r.Context()

	date := param.ByName("date")
	if _, err := time.Parse(time.DateOnly, date); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ii, err := h.sv.Get(ctx, date)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	iiResp := response(ii, h.proxyAdr)

	if err := json.NewEncoder(w).Encode(iiResp); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAll godoc
//
// @Summary      Get All Images Info
// @Description  Gets all recorded Data
// @Tags         imageInfo
// @Produce      json
// @Success      200  {object}  []model.ImageInfoResponse  "image infos"
// @Failure      404  {object}  string  "error message"
// @Router       /get [get]
func (h *Handlers) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	iis, err := h.sv.GetAll(ctx)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	iiResps := make([]model.ImageInfoResponse, 0, len(iis))
	for _, ii := range iis {
		iiResp := response(ii, h.proxyAdr)
		iiResps = append(iiResps, iiResp)
	}

	if err := json.NewEncoder(w).Encode(iiResps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func response(ii model.ImageInfo, proxyAdress string) model.ImageInfoResponse {
	return model.ImageInfoResponse{
		Copyright:      ii.Copyright,
		Date:           ii.Date,
		Explanation:    ii.Explanation,
		UrlHD:          ii.UrlHD,
		MediaType:      ii.MediaType,
		ServiceVersion: ii.ServiceVersion,
		Title:          ii.Title,
		Url:            ii.Url,
		ProxyURL:       proxyAdress + ii.Date + ".jpg",
	}
}
