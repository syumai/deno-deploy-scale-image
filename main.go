//go:generate sh -c "tinygo build -opt=s -o main.wasm -target wasm ./ && cat main.wasm | deno run https://denopkg.com/syumai/binpack/mod.ts > mainwasm.ts && rm main.wasm"
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"syscall/js"

	"github.com/syumai/denoio"
	"golang.org/x/image/draw"
)

var global = js.Global()

func main() {
	scaleImageCallback := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			panic("two args must be given")
		}
		// convert Deno.Reader to Go's io.Reader.
		// - Deno.Reader must be used in Promise in Go side.
		r := denoio.NewReader(args[0])
		width := args[1].Int()
		if width > 2048 {
			panic("width larger than 2048 is not allowed")
		}

		// create callback for Promise
		var cb js.Func
		cb = js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
			resolve := args[0]
			go func() {
				defer cb.Release()
				rd, err := scaleImage(r, width)
				if err != nil {
					panic(err)
				}
				// convert Go's image reader (io.Reader) to Deno.Reader
				result := denoio.NewJSReader(rd)
				resolve.Invoke(result)
			}()
			return js.Undefined()
		})
		return newPromise(cb)
	})
	global.Set("scaleImage", scaleImageCallback)
	select {}
}

func scaleImage(rd io.Reader, width int) (io.Reader, error) {
	m, t, err := image.Decode(rd)
	if err != nil {
		return nil, err
	}
	b := m.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, width, int(float64(b.Dy())/float64(b.Dx())*float64(width))))
	draw.BiLinear.Scale(dst, dst.Bounds(), m, m.Bounds(), draw.Over, nil)

	var buf bytes.Buffer
	switch t {
	case "jpeg":
		if err := jpeg.Encode(&buf, dst, &jpeg.Options{Quality: 100}); err != nil {
			return nil, err
		}
	case "gif":
		if err := gif.Encode(&buf, dst, nil); err != nil {
			return nil, err
		}
	case "png":
		if err := png.Encode(&buf, dst); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown format: %s", t)
	}
	return &buf, nil
}

func newPromise(fn js.Func) js.Value {
	p := global.Get("Promise")
	return p.New(fn)
}
