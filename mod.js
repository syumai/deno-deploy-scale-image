import mainwasm from "./mainwasm.ts";
import "https://gist.githubusercontent.com/paulirish/5438650/raw/ff67d5657042223f6dc5194017f423a0f932f9fa/performance.now()-polyfill.js";
import { Go } from "./wasm_exec.js";
import { decode } from "https://deno.land/std@0.92.0/encoding/base64.ts";
import { readableStreamFromIterable } from "https://deno.land/std@0.92.0/io/streams.ts";
import { Buffer } from "https://deno.land/std@0.92.0/io/buffer.ts";
import { iterSync } from "https://denopkg.com/lucacasonato/deno_std@add-iter-itersync/io/util.ts";
import { serve } from "https://deno.land/x/sift@0.2.0/mod.ts";

const bytes = decode(mainwasm);

const urlBase = "https://github.com/syumai/images/raw/main/";

const handler = async (req, params) => {
  const reqUrl = new URL(req.url);
  const path = reqUrl.searchParams.get("path");
  if (!path) {
    return new Response("path parameter must be given", {
      status: 400,
    })
  }

  const widthStr = reqUrl.searchParams.get("width");
  if (!widthStr) {
    return new Response("width parameter must be given", {
      status: 400,
    })
  }

  let width;
  try {
    width = parseInt(widthStr, 10);
  } catch {
    return new Response("width parameter format is invalid", {
      status: 400,
    })
  }

  if (width > 256) {
    return new Response("width parameter must be smaller than 256", {
      status: 400,
    })
  }

  const go = new Go();
  const result = await WebAssembly.instantiate(bytes, go.importObject);
  go.run(result.instance);

  const url = new URL(path, urlBase);
  const res = await fetch(url);
  const ab = await res.arrayBuffer();
  const buf = new Buffer(ab);

  const scaled = await scaleImage(buf, width);
  const stream = readableStreamFromIterable(Deno.iterSync(scaled));
  return new Response(stream, {
    status: 200,
    headers: {
      server: "denosr",
      "content-type": res.headers.get("content-type"),
    },
  });
};

serve({
  "/image": handler,
});
