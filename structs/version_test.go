package structs_test

import (
	"os"
	"testing"

	pb "github.com/Tackem-org/Global/pb/registration"
	"github.com/Tackem-org/Global/structs"
	"github.com/stretchr/testify/assert"
)

func TestStringToVersion(t *testing.T) {
	v1, err1 := structs.StringToVersion("v1.2.3")
	assert.Equal(t, structs.Version{1, 2, 3}, v1)
	assert.Nil(t, err1)
	v2, err2 := structs.StringToVersion("V1.2.4")
	assert.Equal(t, structs.Version{1, 2, 4}, v2)
	assert.Nil(t, err2)
	v3, err3 := structs.StringToVersion("V1.2.4.5")
	assert.Equal(t, structs.Version{0, 0, 0}, v3)
	assert.ErrorIs(t, err3, structs.ErrBadVersion)
}

func stringToFile(s string, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()
	file.WriteString(s)
}

func TestFileToVersion(t *testing.T) {
	fileName := "test.version"
	stringToFile("0.0.1", fileName)
	v1, _ := structs.FileToVersion(fileName)
	assert.Equal(t, structs.Version{0, 0, 1}, v1)
	os.Remove(fileName)
}

func TestVersionGreaterThan(t *testing.T) {
	tests := []struct {
		left   structs.Version
		right  structs.Version
		result bool
	}{
		{structs.Version{1, 0, 0}, structs.Version{0, 0, 0}, true},
		{structs.Version{0, 1, 0}, structs.Version{0, 0, 0}, true},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 0}, true},
		{structs.Version{1, 0, 0}, structs.Version{1, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 1, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 1}, false},
		{structs.Version{1, 0, 0}, structs.Version{2, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 2, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 2}, false},
	}

	for i, test := range tests {
		assert.Equal(t, test.result, test.left.GreaterThan(test.right), i)
	}
}

func TestVersionGreaterThanOrEqualTo(t *testing.T) {
	tests := []struct {
		left   structs.Version
		right  structs.Version
		result bool
	}{
		{structs.Version{1, 0, 0}, structs.Version{0, 0, 0}, true},
		{structs.Version{0, 1, 0}, structs.Version{0, 0, 0}, true},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 0}, true},
		{structs.Version{1, 0, 0}, structs.Version{1, 0, 0}, true},
		{structs.Version{0, 1, 0}, structs.Version{0, 1, 0}, true},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 1}, true},
		{structs.Version{1, 0, 0}, structs.Version{2, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 2, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 2}, false},
	}

	for i, test := range tests {
		assert.Equal(t, test.result, test.left.GreaterThanOrEqualTo(test.right), i)
	}
}

func TestVersionLessThan(t *testing.T) {
	tests := []struct {
		left   structs.Version
		right  structs.Version
		result bool
	}{
		{structs.Version{1, 0, 0}, structs.Version{0, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 0, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 0}, false},
		{structs.Version{1, 0, 0}, structs.Version{1, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 1, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 1}, false},
		{structs.Version{1, 0, 0}, structs.Version{2, 0, 0}, true},
		{structs.Version{0, 1, 0}, structs.Version{0, 2, 0}, true},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 2}, true},
	}

	for i, test := range tests {
		assert.Equal(t, test.result, test.left.LessThan(test.right), i)
	}
}

func TestVersionLessThanOrEqualTo(t *testing.T) {
	tests := []struct {
		left   structs.Version
		right  structs.Version
		result bool
	}{
		{structs.Version{1, 0, 0}, structs.Version{0, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 0, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 0}, false},
		{structs.Version{1, 0, 0}, structs.Version{1, 0, 0}, true},
		{structs.Version{0, 1, 0}, structs.Version{0, 1, 0}, true},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 1}, true},
		{structs.Version{1, 0, 0}, structs.Version{2, 0, 0}, true},
		{structs.Version{0, 1, 0}, structs.Version{0, 2, 0}, true},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 2}, true},
	}

	for i, test := range tests {
		assert.Equal(t, test.result, test.left.LessThanOrEqualTo(test.right), i)
	}
}

func TestVersionEqualTo(t *testing.T) {
	tests := []struct {
		left   structs.Version
		right  structs.Version
		result bool
	}{
		{structs.Version{1, 0, 0}, structs.Version{0, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 0, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 0}, true},
		{structs.Version{1, 0, 0}, structs.Version{1, 0, 0}, true},
		{structs.Version{0, 1, 0}, structs.Version{0, 1, 0}, true},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 1}, true},
		{structs.Version{1, 0, 0}, structs.Version{2, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 2, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 2}, false},
	}

	for i, test := range tests {
		assert.Equal(t, test.result, test.left.EqualTo(test.right), i)
	}
}

func TestVersionEqualToHotfix(t *testing.T) {
	tests := []struct {
		left   structs.Version
		right  structs.Version
		result bool
	}{
		{structs.Version{1, 0, 0}, structs.Version{0, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 0, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 0}, false},
		{structs.Version{1, 0, 0}, structs.Version{1, 0, 0}, true},
		{structs.Version{0, 1, 0}, structs.Version{0, 1, 0}, true},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 1}, true},
		{structs.Version{1, 0, 0}, structs.Version{2, 0, 0}, false},
		{structs.Version{0, 1, 0}, structs.Version{0, 2, 0}, false},
		{structs.Version{0, 0, 1}, structs.Version{0, 0, 2}, false},
	}

	for i, test := range tests {
		assert.Equal(t, test.result, test.left.EqualToHotfix(test.right), i)
	}
}

// func (v Version) String() string {
// 	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Hotfix)
// }

// func (v Version) ToProto() *pb.Version {
// 	return &pb.Version{
// 		Major:  uint32(v.Major),
// 		Minor:  uint32(v.Minor),
// 		Hotfix: uint32(v.Hotfix),
// 	}
// }

func TestVersionToString(t *testing.T) {
	tests := []struct {
		v structs.Version
		e string
	}{
		{structs.Version{0, 0, 0}, "0.0.0"},
		{structs.Version{1, 0, 0}, "1.0.0"},
		{structs.Version{0, 1, 0}, "0.1.0"},
		{structs.Version{0, 0, 1}, "0.0.1"},
	}

	for i, test := range tests {
		assert.Equal(t, test.e, test.v.String(), i)
	}
}

func TestVersionToProto(t *testing.T) {
	v := structs.Version{0, 0, 0}
	e := pb.Version{Major: 0, Minor: 0, Hotfix: 0}
	p := v.ToProto()
	assert.Equal(t, e.Major, p.Major)
	assert.Equal(t, e.Minor, p.Minor)
	assert.Equal(t, e.Hotfix, p.Hotfix)
}
