package smoke_test

import (
	"flag"
	"testing"
	"time"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var (
	Builder string

	config struct {
		Procfile string `json:"procfile"`
	}
)

func init() {
	flag.StringVar(&Builder, "name", "", "")
}

func TestSmoke(t *testing.T) {
	format.MaxLength = 0
	Expect := NewWithT(t).Expect

	flag.Parse()

	Expect(Builder).NotTo(Equal(""))

	SetDefaultEventuallyTimeout(60 * time.Second)

	suite := spec.New("Buildpackless Smoke", spec.Parallel(), spec.Report(report.Terminal{}))
	suite("Java Native Image", testJavaNativeImage)
	suite("Java", testJava)
	suite("Procfile", testProcfile)
	suite.Run(t)
}
