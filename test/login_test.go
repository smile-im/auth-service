package test

import (
	"log"
	"testing"

	"github.com/smile-im/microkit-client/client/auth"
	"github.com/smile-im/microkit-client/proto/authpb"
)

var (
	cl authpb.AuthClient
)

func TestMain(m *testing.M) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var err error
	cl, err = auth.NewClient()
	if err != nil {
		log.Panicln(err)
	}
	m.Run()
}
