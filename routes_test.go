package gowebapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func stringbody(res *http.Response, t *testing.T) string {
	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(got)
}

func TestHome(t *testing.T) {
	ts := httptest.NewServer(NewMux())
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/")
	b := stringbody(res, t)

	if strings.Index(b, "home index") == -1 {
		t.Errorf("wrong template %+v", b)
	}
}

func TestAccountIndex(t *testing.T) {
	ts := httptest.NewServer(NewMux())
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/accounts")
	b := stringbody(res, t)

	if strings.Index(b, "accounts index") == -1 {
		t.Errorf("wrong template %+v", b)
	}

	if strings.Index(b, "Main Layout") == -1 {
		t.Errorf("wrong layout %+v", b)
	}
}

func TestPopupEdit(t *testing.T) {
	ts := httptest.NewServer(NewMux())
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/accounts/1/popup_edit")
	b := stringbody(res, t)
	if strings.Index(b, "Popup Layout") == -1 {
		t.Errorf("wrong layout %+v", b)
	}
}
