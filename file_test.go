package enu_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuemori/enu"
)

func TestFileReader(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	f, err := os.CreateTemp(os.TempDir(), "enu-filetest-")
	if err != nil {
		t.Fatal(err)
	}

	for _, s := range []string{"foo\n", "bar\n", "baz"} {
		if _, err = f.Write([]byte(s)); err != nil {
			t.Fatal(err)
		}
	}

	reader := enu.FromFile(f.Name())
	r1, ok := reader.First()
	is.True(ok)
	is.Equal("foo", r1)
	is.NoError(reader.Err())

	r2 := reader.ToSlice()
	is.Equal([]string{"foo", "bar", "baz"}, r2)
	is.NoError(reader.Err())

	var r3 []string
	err = reader.Result(&r3).Err()
	is.Equal([]string{"foo", "bar", "baz"}, r3)
	is.NoError(err)

	err = reader.Each(func(line string, index int) {
		switch index {
		case 0:
			is.Equal("foo", line)
		case 1:
			is.Equal("bar", line)
		case 2:
			is.Equal("baz", line)
		default:
			t.Fatalf("unknown index %d", index)
		}
	}).Err()
	is.NoError(err)

	if err := os.Remove(f.Name()); err != nil {
		t.Fatal(err)
	}

	r4 := reader.ToSlice()
	is.Equal([]string{}, r4)
	is.ErrorContains(reader.Err(), "file already closed")

	reader2 := enu.FromFile(f.Name())
	_, ok = reader2.First()
	is.False(ok)
	is.ErrorContains(reader2.Err(), "no such file or directory")
}
