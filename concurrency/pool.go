package concurrency

import (
	"bytes"
	"io"
	"sync"
	"time"
)

var bufPool =sync.Pool{
	New:func() interface{} {
		return new(bytes.Buffer)
	},
}


func Log(w io.Writer, key, val string) {
	b :=bufPool.Get().(*bytes.Buffer)
	b.Reset()

	b.WriteString(time.Now().String())
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	b.WriteString("\n")
	w.Write(b.Bytes())
	bufPool.Put(b)
}