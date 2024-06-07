package main

import "testing"

func TestGetNTPTime(t *testing.T) {
	server := "pool.ntp.org"

	t.Run("TestGetNTPTime", func(t *testing.T) {
		// Test getting NTP time
		ntpTime, err := getNTPTime(server)

		if err != nil {
			t.Errorf("error getting NTP time: %v", err)
		}

		if ntpTime.IsZero() {
			t.Error("NTP time is zero")
		}
	})

	t.Run("TestInvalidServer", func(t *testing.T) {
		// Test providing an invalid NTP server
		_, err := getNTPTime("invalid.ntp.server")

		if err == nil {
			t.Error("expected error for invalid NTP server, got nil")
		}
	})
}
