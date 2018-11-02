package interrupter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestSigInt_WaitNoSignal(t *testing.T) {
	c := New(func() {
		t.Fail()
	})
	c.Close()
	i := c.(*interrupter)
	_, ok := <-i.c
	assert.False(t, ok)
}

func TestSigInt_WaitSignal(t *testing.T) {
	b := make(chan bool)
	c := New(func() {
		b <- true
	})
	i := c.(*interrupter)
	i.c <- os.Interrupt
	time.Sleep(time.Second)
	select {
	case <-b:
	default:
		t.Fail()
	}
	c.Close()
	_, ok := <-i.c
	assert.False(t, ok)
}
