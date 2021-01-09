package extxt

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	got := &bytes.Buffer{}
	targetImagePath := "./testdata/image.JPG"
	want := "公開鍵暗号によって鍵配送問題は解決しましたが、公開鍵暗号に対してはman-in-the-middle攻撃が可能です。"
	var wantErr error

	gotErr := Run(got, targetImagePath)
	if gotErr != wantErr {
		t.Errorf("\nwant: \n%v\ngot: \n%v\n", wantErr, gotErr)
	}
	if !strings.Contains(got.String(), want) {
		t.Errorf("\nwant: \n%s\ngot: \n%s\n", want, got)
	}
}

func TestRunByServer(t *testing.T) {
	got := &bytes.Buffer{}
	file, err := os.Open("./testdata/image.JPG")
	if err != nil {
		t.Fatal(err)
	}

	want := "公開鍵暗号によって鍵配送問題は解決しましたが、公開鍵暗号に対してはman-in-the-middle攻撃が可能です。"
	var wantErr error

	gotErr := RunByServer(got, file)
	if gotErr != wantErr {
		t.Errorf("\nwant: \n%v\ngot: \n%v\n", wantErr, gotErr)
	}
	if !strings.Contains(got.String(), want) {
		t.Errorf("\nwant: \n%s\ngot: \n%s\n", want, got)
	}
}
