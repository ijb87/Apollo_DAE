package iosplitter

import "io"

type splitReader struct {
	id       int
	ch       chan spinT
	donechan chan bool
}

type spinT struct {
	id int
	b  []byte
	ch chan spoutT
}

type spoutT struct {
	n   int
	err error
}

type routT struct {
	d   []byte
	err error
}

func Splitn(r io.Reader, n int) []io.Reader {

	rout := make(chan routT)
	rin := make(chan int)
	//read Direct from input io.Reader and send out on channel
	go func() {
		var dbuff []byte
		for n := range rin {
			if len(dbuff) != n {
				dbuff = make([]byte, n)
			}
			n, err := r.Read(dbuff)
			rout <- routT{d: dbuff[:n], err: err}
			if err != nil {
				return
			}
		}
	}()

	spin := make(chan spinT)

	donechan := make(chan bool)
	go func() {
		buff := []byte{}
		pos := make([]int, n)
		rwaiting := false
		done := 0

		var e error = nil
		for {
			select {
			case sp := <-spin:
				p := pos[sp.id]
				if len(buff) < p+len(sp.b) && !rwaiting && e == nil {
					rin <- p + len(sp.b) - len(buff)
					rwaiting = true
				}
				nc := copy(sp.b, buff[p:])
				pos[sp.id] += nc
				sp.ch <- spoutT{n: nc, err: e}
			case ro := <-rout:
				buff = drop(pos, buff)
				buff = append(buff, ro.d...)
				if ro.err != nil {
					e = ro.err
				}
				rwaiting = false
			case <-donechan:
				done++
				if done >= n {
					return
				}
			}
		}
	}()

	res := []io.Reader{}
	for i := 0; i < n; i++ {
		res = append(res, &splitReader{id: i, ch: spin, donechan: donechan})
	}
	return res
}

func (sp *splitReader) Read(b []byte) (int, error) {
	ch := make(chan spoutT)
	sp.ch <- spinT{b: b, id: sp.id, ch: ch}

	res := <-ch

	if res.err != nil {
		sp.donechan <- true
	}
	return res.n, res.err
}

func drop(p []int, b []byte) []byte {
	lowest := p[0]
	for _, v := range p {
		if v == 0 {
			return b
		}
		if v < lowest {
			lowest = v
		}
	}

	for k := range p {
		p[k] -= lowest
	}

	return b[lowest:]
}