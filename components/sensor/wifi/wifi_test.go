package wifi

import (
	"context"
	"os"
	"testing"

	"github.com/edaniels/golog"
	"go.viam.com/test"
)

func createWirelessInfoFile(t *testing.T) *os.File {
	t.Helper()

	dirPath := t.TempDir()
	file, err := os.CreateTemp(dirPath, "wireless")
	test.That(t, err, test.ShouldBeNil)

	return file
}

func TestNewSensor(t *testing.T) {
	logger := golog.NewLogger("testlog")

	file := createWirelessInfoFile(t)

	_, err := newWifi(logger, "wrong-path")
	test.That(t, err, test.ShouldNotBeNil)

	_, err = newWifi(logger, file.Name())
	test.That(t, err, test.ShouldBeNil)
}

func TestReadings(t *testing.T) {
	ctx := context.Background()
	logger := golog.NewLogger("testlog")

	file := createWirelessInfoFile(t)
	_, err := file.WriteString(`Inter-| sta-|   Quality        |   Discarded packets               | Missed | WE
 face | tus | link level noise |  nwid  crypt   frag  retry   misc | beacon | 22
XXXXXXXXX: XXXX   58.  -52.  -256        X      X      X      X  XXXXX        X`)
	test.That(t, err, test.ShouldBeNil)

	sensor, err := newWifi(logger, file.Name())
	test.That(t, err, test.ShouldBeNil)

	readings, err := sensor.Readings(ctx, nil)
	test.That(t, err, test.ShouldBeNil)
	expected := map[string]interface{}{
		"link":  int(58),
		"level": int(-52),
		"noise": int(-256),
	}
	test.That(t, readings, test.ShouldResemble, expected)
}