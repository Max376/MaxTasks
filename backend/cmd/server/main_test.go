package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakePinger struct {
	err   error
	calls int
}

func (pinger *fakePinger) PingContext(context.Context) error {
	pinger.calls++
	return pinger.err
}

func TestHealthEndpointReturnsHealthyProcessStatus(t *testing.T) {
	pinger := &fakePinger{err: errors.New("database is unavailable")}
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	newMux(pinger).ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("health status = %d, want %d", recorder.Code, http.StatusOK)
	}
	response := decodeStatus(t, recorder)
	if response.Service != serviceName || response.Status != "ok" {
		t.Fatalf("health response = %#v, want service %q and status %q", response, serviceName, "ok")
	}
	if response.Checks != nil {
		t.Fatalf("health response unexpectedly contains checks: %#v", response.Checks)
	}
	if pinger.calls != 0 {
		t.Fatalf("health ping calls = %d, want 0", pinger.calls)
	}
}

func TestReadinessEndpointReturnsReadyWhenDatabaseResponds(t *testing.T) {
	pinger := &fakePinger{}
	request := httptest.NewRequest(http.MethodGet, "/ready", nil)
	recorder := httptest.NewRecorder()

	newMux(pinger).ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("readiness status = %d, want %d", recorder.Code, http.StatusOK)
	}
	response := decodeStatus(t, recorder)
	if response.Service != serviceName || response.Status != "ready" {
		t.Fatalf("readiness response = %#v, want service %q and status %q", response, serviceName, "ready")
	}
	if response.Checks == nil || response.Checks.Database != "ok" {
		t.Fatalf("readiness database check = %#v, want ok", response.Checks)
	}
	if pinger.calls != 1 {
		t.Fatalf("readiness ping calls = %d, want 1", pinger.calls)
	}
}

func TestReadinessEndpointReturnsUnavailableWhenDatabaseFails(t *testing.T) {
	pinger := &fakePinger{err: errors.New("dial tcp: password=secret")}
	request := httptest.NewRequest(http.MethodGet, "/ready", nil)
	recorder := httptest.NewRecorder()

	newMux(pinger).ServeHTTP(recorder, request)

	if recorder.Code != http.StatusServiceUnavailable {
		t.Fatalf("readiness status = %d, want %d", recorder.Code, http.StatusServiceUnavailable)
	}
	response := decodeStatus(t, recorder)
	if response.Status != "not_ready" {
		t.Fatalf("readiness status field = %q, want not_ready", response.Status)
	}
	if response.Checks == nil || response.Checks.Database != "unavailable" {
		t.Fatalf("readiness database check = %#v, want unavailable", response.Checks)
	}
	if strings.Contains(recorder.Body.String(), "secret") || strings.Contains(recorder.Body.String(), "dial tcp") {
		t.Fatalf("readiness response leaks internal error: %s", recorder.Body.String())
	}
	if pinger.calls != 1 {
		t.Fatalf("readiness ping calls = %d, want 1", pinger.calls)
	}
}

func TestReadinessEndpointReturnsUnavailableWithoutDatabase(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/ready", nil)
	recorder := httptest.NewRecorder()

	newMux(nil).ServeHTTP(recorder, request)

	if recorder.Code != http.StatusServiceUnavailable {
		t.Fatalf("readiness status = %d, want %d", recorder.Code, http.StatusServiceUnavailable)
	}
	response := decodeStatus(t, recorder)
	if response.Checks == nil || response.Checks.Database != "unavailable" {
		t.Fatalf("readiness database check = %#v, want unavailable", response.Checks)
	}
}

func decodeStatus(t *testing.T, recorder *httptest.ResponseRecorder) serviceStatus {
	t.Helper()
	var response serviceStatus
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	return response
}
