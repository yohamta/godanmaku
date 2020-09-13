GOOS=js GOARCH=wasm go build -o wasm/public/godanmaku.wasm github.com/yohamta/godanmaku
cd wasm/ && firebase deploy
