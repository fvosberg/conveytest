package main

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCustomerRegistration(t *testing.T) {
	Convey("Given a running elasticsearch server", t, func() {
		_, err := http.Get("http://elasticsearch:9200")
		So(err, ShouldBeNil)
	})
}
