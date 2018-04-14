package decimal

var pool decimalPool

type decimalPool struct {
	c chan Decimal
}

// Get gets a Buffer from the BufferPool, or creates a new one if none are
// available in the pool.
func (bp *decimalPool) Get() (b Decimal) {
	select {
	case b = <-bp.c:
		// reuse existing buffer
	default:
		// create new buffer
		b = Decimal{}
	}
	return
}

// Put returns the given Buffer to the BufferPool.
func (bp *decimalPool) Put(b Decimal) {
	b.b.setZero(0, 0)
	select {
	case bp.c <- b:
	default: // Discard the buffer if the pool is full.
	}
}


func SetPoolSize(size int) {
	pool = decimalPool{c:make(chan Decimal, size)}
}

func PutDecimal(d Decimal) {
	if pool.c != nil {
		pool.Put(d)
	}
}