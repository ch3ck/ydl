package lame

// #cgo LDFLAGS: -lmp3lame
// #include <lame/lame.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Handle *C.struct_lame_global_struct

const (
	STEREO        = C.STEREO
	JOINT_STEREO  = C.JOINT_STEREO
	DUAL_CHANNEL  = C.DUAL_CHANNEL /* LAME doesn't supports this! */
	MONO          = C.MONO
	NOT_SET       = C.NOT_SET
	MAX_INDICATOR = C.MAX_INDICATOR
	BIT_DEPTH     = 16
)

type Encoder struct {
	handle    Handle
	remainder []byte
	closed    bool
}

func Init() *Encoder {
	handle := C.lame_init()
	encoder := &Encoder{handle, make([]byte, 0), false}
	runtime.SetFinalizer(encoder, finalize)
	return encoder
}

func (e *Encoder) SetNumChannels(num int) {
	C.lame_set_num_channels(e.handle, C.int(num))
}

func (e *Encoder) SetInSamplerate(sampleRate int) {
	C.lame_set_in_samplerate(e.handle, C.int(sampleRate))
}

func (e *Encoder) SetBitrate(bitRate int) {
	C.lame_set_brate(e.handle, C.int(bitRate))
}

func (e *Encoder) SetMode(mode C.MPEG_mode) {
	C.lame_set_mode(e.handle, mode)
}

func (e *Encoder) SetQuality(quality int) {
	C.lame_set_quality(e.handle, C.int(quality))
}

func (e *Encoder) InitParams() int {
	retcode := C.lame_init_params(e.handle)
	return int(retcode)
}

func (e *Encoder) NumChannels() int {
	n := C.lame_get_num_channels(e.handle)
	return int(n)
}

func (e *Encoder) Bitrate() int {
	br := C.lame_get_brate(e.handle)
	return int(br)
}

func (e *Encoder) Mode() int {
	m := C.lame_get_mode(e.handle)
	return int(m)
}

func (e *Encoder) Quality() int {
	q := C.lame_get_quality(e.handle)
	return int(q)
}

func (e *Encoder) InSamplerate() int {
	sr := C.lame_get_in_samplerate(e.handle)
	return int(sr)
}

func (e *Encoder) Encode(buf []byte) []byte {

	if len(e.remainder) > 0 {
		buf = append(e.remainder, buf...)
	}

	if len(buf) == 0 {
		return make([]byte, 0)
	}

	blockAlign := BIT_DEPTH / 8 * e.NumChannels()

	remainBytes := len(buf) % blockAlign
	if remainBytes > 0 {
		e.remainder = buf[len(buf)-remainBytes:]
		buf = buf[0 : len(buf)-remainBytes]
	} else {
		e.remainder = make([]byte, 0)
	}

	numSamples := len(buf) / blockAlign
	estimatedSize := int(1.25*float64(numSamples) + 7200)
	out := make([]byte, estimatedSize)

	cBuf := (*C.short)(unsafe.Pointer(&buf[0]))
	cOut := (*C.uchar)(unsafe.Pointer(&out[0]))

	bytesOut := C.int(C.lame_encode_buffer_interleaved(
		e.handle,
		cBuf,
		C.int(numSamples),
		cOut,
		C.int(estimatedSize),
	))
	return out[0:bytesOut]

}

func (e *Encoder) Flush() []byte {
	estimatedSize := 7200
	out := make([]byte, estimatedSize)
	cOut := (*C.uchar)(unsafe.Pointer(&out[0]))
	bytesOut := C.int(C.lame_encode_flush(
		e.handle,
		cOut,
		C.int(estimatedSize),
	))

	return out[0:bytesOut]
}

func (e *Encoder) Close() {
	if e.closed {
		return
	}
	C.lame_close(e.handle)
	e.closed = true
}

func finalize(e *Encoder) {
	e.Close()
}
