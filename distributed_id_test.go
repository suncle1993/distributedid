package distributedid

import (
	"testing"
)

func TestIDGenerator(t *testing.T) {
	var generator IDGenerator
	var err error

	{
		// NewIDGenerator fail
		generator, err = NewIDGenerator("10.129.3111.5")
		if err == nil || err.Error() != ErrInvalidIP.Error() {
			t.Errorf("ip validate error")
		}
	}

	{
		// generate success without option
		generator, err = NewIDGenerator("10.129.31.5")
		if err != nil {
			t.Errorf("NewIDGenerator error, error: %v", err)
		}
		id1 := generator.Generate().Int64()
		id2 := generator.Generate().Int64()
		id3 := generator.Generate().Int64()
		if id1 > id2 || id2 > id3 {
			t.Errorf("IDGenerator not auto increment")
		}
	}

	{
		// generate success with Node option
		generator, err = NewIDGenerator("10.129.31.5", Node(1))
		if err != nil {
			t.Errorf("NewIDGenerator error, error: %v", err)
		}
		if !checkGeneratedIDS(generator) {
			t.Errorf("IDGenerator not auto increment")
		}
	}

	{
		// generate success with NodeBits option
		generator, err = NewIDGenerator("10.129.31.5", NodeBits(15))
		if err != nil {
			t.Errorf("NewIDGenerator error, error: %v", err)
		}
		if !checkGeneratedIDS(generator) {
			t.Errorf("IDGenerator not auto increment")
		}
	}
}

func checkGeneratedIDS(generator IDGenerator) (pass bool) {
	id1 := generator.Generate().Int64()
	id2 := generator.Generate().Int64()
	id3 := generator.Generate().Int64()
	return id1 < id2 && id2 < id3
}
