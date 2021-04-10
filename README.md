# deno-deploy-scale-image

* image scaling app built for Deno Deploy using Go WebAssembly.
  - this app fetches image from [syumai's GitHub Repo](https://github.com/syumai/images).

## Requirements

* tinygo
* Deno
* [deployctl](https://deno.com/deploy/docs/deployctl)

## Usage

* Visit URL: https://scale-image.deno.dev/image?path=${imagePath}&width=${widthNumber}
  - Example: https://scale-image.deno.dev/image?path=landscape.jpg&width=200

available images

* syumai.png
* landscape.jpg
* ramen.jpg

## Development

### Run app locally

```
make run
```

### Scale image

show scaled image on browser

`http://0.0.0.0:8080/image?path=landscape.jpg&width=200`

## Author

syumai

## License

MIT

