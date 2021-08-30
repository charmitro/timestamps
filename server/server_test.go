package server

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPeriod1h(t *testing.T) {
	req, err := http.Get("http://localhost:8080/ptlist?period=1h&tz=UTC&t1=20210714T204603Z&t2=20210715T123456Z")

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))
	if req.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", req.StatusCode)
	}

	if string(body) != `["20210714T210000Z","20210714T220000Z","20210714T230000Z","20210715T000000Z","20210715T010000Z","20210715T020000Z","20210715T030000Z","20210715T040000Z","20210715T050000Z","20210715T060000Z","20210715T070000Z","20210715T080000Z","20210715T090000Z","20210715T100000Z","20210715T110000Z","20210715T120000Z","20210715T130000Z"]` {
		t.Fatalf(`expected
["20210714T210000Z","20210714T220000Z","20210714T230000Z","20210715T000000Z","20210715T010000Z","20210715T020000Z","20210715T030000Z","20210715T040000Z","20210715T050000Z","20210715T060000Z","20210715T070000Z","20210715T080000Z","20210715T090000Z","20210715T100000Z","20210715T110000Z","20210715T120000Z","20210715T130000Z"]
got
%s`, string(body))
	}

}

func TestPeriod1d(t *testing.T) {
	req, err := http.Get("http://localhost:8080/ptlist?period=1d&tz=Europe/Athens&t1=20211010T204603Z&t2=20211115T123456Z")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))
	if req.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", req.StatusCode)
	}

	if string(body) != `["20211010T210000Z","20211011T210000Z","20211012T210000Z","20211013T210000Z","20211014T210000Z","20211015T210000Z","20211016T210000Z","20211017T210000Z","20211018T210000Z","20211019T210000Z","20211020T210000Z","20211021T210000Z","20211022T210000Z","20211023T210000Z","20211024T210000Z","20211025T210000Z","20211026T210000Z","20211027T210000Z","20211028T210000Z","20211029T210000Z","20211030T210000Z","20211031T210000Z","20211101T210000Z","20211102T210000Z","20211103T210000Z","20211104T210000Z","20211105T210000Z","20211106T210000Z","20211107T210000Z","20211108T210000Z","20211109T210000Z","20211110T210000Z","20211111T210000Z","20211112T210000Z","20211113T210000Z","20211114T210000Z","20211115T210000Z"]` {
		t.Fatalf(`expected
["20211010T210000Z","20211011T210000Z","20211012T210000Z","20211013T210000Z","20211014T210000Z","20211015T210000Z","20211016T210000Z","20211017T210000Z","20211018T210000Z","20211019T210000Z","20211020T210000Z","20211021T210000Z","20211022T210000Z","20211023T210000Z","20211024T210000Z","20211025T210000Z","20211026T210000Z","20211027T210000Z","20211028T210000Z","20211029T210000Z","20211030T210000Z","20211031T210000Z","20211101T210000Z","20211102T210000Z","20211103T210000Z","20211104T210000Z","20211105T210000Z","20211106T210000Z","20211107T210000Z","20211108T210000Z","20211109T210000Z","20211110T210000Z","20211111T210000Z","20211112T210000Z","20211113T210000Z","20211114T210000Z","20211115T210000Z"]
got
%s`, string(body))
	}
}

func TestPeriod1mo(t *testing.T) {
	req, err := http.Get("http://localhost:8080/ptlist?period=1mo&tz=Europe/Athens&t1=20210214T204603Z&t2=20211115T123456Z")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))
	if req.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", req.StatusCode)
	}

	if string(body) != `["20210214T210000Z","20210314T210000Z","20210414T210000Z","20210514T210000Z","20210614T210000Z","20210714T210000Z","20210814T210000Z","20210914T210000Z","20211014T210000Z","20211114T210000Z","20211214T210000Z"]` {
		t.Fatalf(`expected
["20210214T210000Z","20210314T210000Z","20210414T210000Z","20210514T210000Z","20210614T210000Z","20210714T210000Z","20210814T210000Z","20210914T210000Z","20211014T210000Z","20211114T210000Z","20211214T210000Z"]
got
%s`, string(body))
	}
}

func TestPeriod1y(t *testing.T) {
	req, err := http.Get("http://localhost:8080/ptlist?period=1y&tz=Europe/Athens&t1=20180214T204603Z&t2=20211115T123456Z")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))
	if req.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", req.StatusCode)
	}

	if string(body) != `["20181231T220000Z","20191231T220000Z","20201231T220000Z","20211231T220000Z"]` {
		t.Fatalf(`expected
["20181231T220000Z","20191231T220000Z","20201231T220000Z","20211231T220000Z"]
got
%s`, string(body))
	}
}

func TestError(t *testing.T) {
	req, err := http.Get("http://localhost:8080/ptlist?period=1w&tz=Europe/Athens&t1=20180214T204603Z&t2=20211115T123456Z")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))
	if req.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", req.StatusCode)
	}

	if string(body) != `{"status":"error","desc":"Unsupported period"}` {
		t.Fatalf(`expected
{"status":"error","desc":"Unsupported period"}
got
%s`, string(body))
	}
}

//
