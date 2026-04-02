package googleipmi

import (
	"testing"
)

// Ensure that Open doesn't crash when called from a machine that doesn't have
// the `/dev/ipmi0` device path.
func TestGetTitanIpmiTPM(t *testing.T) {
	tpm, err := Open()
	if tpm != nil {
		t.Fatalf("Expected a nil TPM.")
	}
	if err == nil {
		t.Fatalf("Expected a non-nil error.")
	}
}
