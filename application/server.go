package application

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/http/httputil"
)

type HandleRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type HandlerResponse struct {
	ID      string            `json:"id"`
	Status  string            `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int64             `json:"length"`
}

type ProxyHandler struct {
	client *http.Client
}

func NewProxyHandler() *ProxyHandler {
	return &ProxyHandler{
		client: &http.Client{},
	}
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqDump, _ := httputil.DumpRequest(r, true)

	var inputReq HandleRequest
	err := json.NewDecoder(r.Body).Decode(&inputReq)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		p.logf("error decoding the inputReq: %v", err)
		return
	}

	outputReq, err := http.NewRequest(inputReq.Method, inputReq.URL, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		p.logf("error building the inputReq: %v", err)
		return
	}
	setRequestHeaders(outputReq, inputReq.Headers)

	resp, err := p.client.Do(outputReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		p.logf("error doing the inputReq: %v", err)
		return
	}
	respDump, _ := httputil.DumpResponse(resp, true)
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	id := uuid.New().String()
	rsp := HandlerResponse{
		ID:      id,
		Status:  resp.Status,
		Headers: headers(resp.Header),
		Length:  resp.ContentLength,
	}
	json.NewEncoder(w).Encode(rsp)

	db := DB()
	db.Save(id, reqDump, respDump)
}

func (p *ProxyHandler) logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func setRequestHeaders(destReq *http.Request, srcHeaders map[string]string) {
	for k, v := range srcHeaders {
		destReq.Header.Set(k, v)
	}
}

func headers(h http.Header) map[string]string {
	hs := make(map[string]string)
	for k := range h {
		hs[k] = h.Get(k)
	}
	return hs
}
