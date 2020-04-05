package iosplitter

import (
	"bytes"
	"io"
	"testing"
)

var pdata []byte = []byte("There we are I'm writing  test case for having lots of bytes to copy from a reader to another reader, in fact I'm really hoping we can split a single reader into a hundred different readers and each of them will have something call to talk about")

func Test_Serial(t *testing.T) {

	r := bytes.NewBuffer(pdata)

	rr := Splitn(r, 5)

	if len(rr) < 5 {
		t.Errorf("Where are my readers")
	}

	ress := make([]bytes.Buffer, 5)
	for i := 0; i < 5; i++ {
		io.Copy(&ress[i], rr[i])
	}

	for k, v := range ress {
		if !bytes.Equal(v.Bytes(), pdata) {
			t.Errorf("Unexpected Read: %d : %s", k, v.Bytes())
		}
	}
}

func Test_Concurr(t *testing.T) {
	r := bytes.NewBuffer(pdata)
	nsplit := 50

	rr := Splitn(r, nsplit)

	dchan := make(chan bool)
	ress := make([]bytes.Buffer, nsplit)
	for i := 0; i < nsplit; i++ {
		go func(w io.Writer, r io.Reader) {
			io.CopyBuffer(w, r, make([]byte, 20))
			dchan <- true
		}(&ress[i], rr[i])
	}

	for i := 0; i < nsplit; i++ {
		<-dchan
	}

	for k, v := range ress {
		if !bytes.Equal(v.Bytes(), pdata) {
			t.Errorf("Unexpected Read: %d : %s", k, v.Bytes())
		}
	}

}