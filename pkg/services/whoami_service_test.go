package services

import (
	"testing"
)

func TestWhoamiSvc_Ping(t *testing.T) {
	// Inject mock DB into User service
	service := WhoamiSvc{}

	// Mock a GetAll() call
	err := service.Ping()
	if err != nil {
		t.Fatalf("Error calling Ping: %v", err)
	}
}
