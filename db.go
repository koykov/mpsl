package mpsl

import (
	"sync/atomic"

	"github.com/koykov/bytealg"
	"github.com/koykov/fastconv"
	"github.com/koykov/hash"
	"github.com/koykov/policy"
)

const (
	FullURL = "https://raw.githubusercontent.com/publicsuffix/list/master/public_suffix_list.dat"
)

const (
	statusNil = iota
	statusActive
)

type DB struct {
	policy.RWLock
	status uint32
	hasher hash.BHasher
	index  index
	buf    []byte
}

var (
	bSpace      = []byte(" ")
	bDot        = []byte(".")
	bMaskAll    = []byte("*.")
	bBeginICANN = []byte("// ===BEGIN ICANN DOMAINS===")
	bEndICANN   = []byte("// ===BEGIN ICANN DOMAINS===")
)

func New(hasher hash.BHasher) (*DB, error) {
	if hasher == nil {
		return nil, ErrNoHasher
	}
	db := &DB{
		status: statusActive,
		hasher: hasher,
		index:  make(index),
	}
	return db, nil
}

func (db *DB) Get(hostname []byte) (ps []byte) {
	if err := db.checkStatus(); err != nil {
		return nil
	}
	ps, _ = db.GetWP(hostname)
	return
}

func (db *DB) GetWP(hostname []byte) ([]byte, int) {
	if err := db.checkStatus(); err != nil {
		return nil, -1
	}
	hl := len(hostname)
	if hl < 2 {
		return nil, -1
	}

	var off, dc int
	if dc = dcOf(hostname) - 1; dc < 0 {
		return nil, -1
	}
	db.RLock()
	defer db.RUnlock()
	for {
		if off = bytealg.IndexAt(hostname, bDot, off); off == -1 {
			break
		}
		off++
		p := hostname[off:]
		h := db.hasher.Sum64(p)
		if e, ok := db.index[h]; ok {
			lo, hi := e.decode()
			eb := db.buf[lo:hi]
			return eb, off
		}
		if dc -= 1; dc == -1 {
			break
		}
	}

	return nil, -1
}

func (db *DB) GetStr(hostname string) (ps string) {
	ps, _ = db.GetStrWP(hostname)
	return
}

func (db *DB) GetStrWP(hostname string) (ps string, pos int) {
	x, p := db.GetWP(fastconv.S2B(hostname))
	if p == -1 {
		return
	}
	ps, pos = fastconv.B2S(x), p
	return
}

func (db *DB) entryBytes(e entry) []byte {
	lo, hi := e.decode()
	if hi >= uint32(len(db.buf)) || lo > hi {
		return nil
	}
	return db.buf[lo:hi]
}

func (db *DB) Reset() {
	if err := db.checkStatus(); err != nil {
		return
	}
	db.SetPolicy(policy.Locked)
	db.Lock()
	db.index.reset()
	db.buf = db.buf[:0]
	db.Unlock()
	db.SetPolicy(policy.LockFree)
}

func (db *DB) checkStatus() error {
	if atomic.LoadUint32(&db.status) == statusNil {
		return ErrBadDB
	}
	return nil
}

func psMustSkip(line []byte) bool {
	if len(line) == 0 || line[0] == '/' || line[0] == '!' {
		return true
	}
	return false
}

func dcOf(p []byte) (dc int) {
	off := 0
loop:
	if off = bytealg.IndexAt(p, bDot, off); off != -1 {
		dc++
		off++
		goto loop
	}
	return
}
