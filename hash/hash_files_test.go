package hash

import (
	"fmt"
	"testing"
)

func TestComputeDigest(t *testing.T) {
	t.Log("Given we want to calculate a SHA-256 Hash")
	{
		var digest []byte
		digestName := "SHA-256"
		t.Log("When calculate the Hash for 'Hello World!'")
		{
			digest = ComputeDigest(digestName, []byte("Hello World!"))

		}
		expectedDigest := "7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069"
		t.Logf("Then the result should be %s", expectedDigest)

		result := fmt.Sprintf("%x", digest)
		//result := string(digest)
		if result == expectedDigest {
			t.Log("OK")
		} else {
			t.Errorf("ERROR, obtained %s instead of %s", result, expectedDigest)
		}
	}
}
