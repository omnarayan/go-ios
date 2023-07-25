//go:build integration
// +build integration

package accessibility_test

import (
	"testing"

	ios "github.com/izinga/go-ios/ios"
	"github.com/izinga/go-ios/ios/accessibility"
	log "github.com/sirupsen/logrus"
)

func TestIT(t *testing.T) {
	device, err := ios.GetDevice("")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := accessibility.New(device)
	if err != nil {
		t.Fatal(err)
	}

	conn.SwitchToDevice()
	if err != nil {
		t.Fatal(err)
	}
	conn.EnableSelectionMode()
	conn.GetElement()
	conn.GetElement()
	conn.TurnOff()

	// conn.EnableSelectionMode()
}
