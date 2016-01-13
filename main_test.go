package mobitool
import (
	"testing"
	"log"
)

func TestDecode(t *testing.T) {
	content, err := Decode("BookOfTexas.mobi")
	if err != nil {
		t.Error(err)
	}
	log.Print(content)
}
