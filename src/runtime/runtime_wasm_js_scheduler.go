//go:build wasm && !wasi && !wasm_freestanding && !scheduler.none
// +build wasm,!wasi,!wasm_freestanding,!scheduler.none

package runtime

//export resume
func resume() {
	go func() {
		handleEvent()
	}()

	if wasmNested {
		minSched()
		return
	}

	wasmNested = true
	scheduler()
	wasmNested = false
}

//export go_scheduler
func go_scheduler() {
	if wasmNested {
		minSched()
		return
	}

	wasmNested = true
	scheduler()
	wasmNested = false
}
