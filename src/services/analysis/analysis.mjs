import "../../libs/go-wasm-runtime.js";

const go = new Go();
await WebAssembly.instantiateStreaming(
  fetch("/services/analysis/analysis.wasm"),
  go.importObject
).then((result) => {
  go.run(result.instance);
});

export function selectDirectoryForAnalysis() {
  ingestDirectory();
}
