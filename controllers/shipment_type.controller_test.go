package controllers

import (
	"net/http"
	"testing"
)

func TestCreateShipmentType(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateShipmentType(tt.args.w, tt.args.r)
		})
	}
}

func TestGetShipmentTypes(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetShipmentTypes(tt.args.w, tt.args.r)
		})
	}
}

func TestUpdateShipmentType(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateShipmentType(tt.args.w, tt.args.r)
		})
	}
}