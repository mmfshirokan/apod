package repository

import (
	"os"
	"testing"
)

var (
	testImgName     = "gilmore"
	testImgLocation = "../../www/html/"
	testImgDest     = "./"
)

func TestAddImage(t *testing.T) {
	img := NewImage(testImgDest)

	f, err := os.Open(testImgLocation + testImgName + ".jpg")
	if err != nil {
		t.Fatal(err)
	}

	if err = img.Add(f, testImgName); err != nil {
		t.Fatal(err)
	}

	if err = os.Remove(testImgDest + testImgName + ".jpg"); err != nil {
		t.Fatal(err)
	}

	t.Log("TestAddImage finished sucssesfully")
}
