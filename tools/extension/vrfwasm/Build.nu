 GOOS=js GOARCH=wasm go build -trimpath -gcflags="-e" -ldflags="-s -w" -o vrf.wasm
 wasm-opt --enable-bulk-memory -Oz -o ../chrome/vrf.wasm vrf.wasm
 rm vrf.wasm