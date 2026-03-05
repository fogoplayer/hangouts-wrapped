import "../../libs/go-wasm-runtime.js";

const go = new Go();
WebAssembly.instantiateStreaming(
  fetch("/services/analysis/analysis.wasm"),
  go.importObject
).then((result) => {
  go.run(result.instance);
});

/** @param {Event} event */
export function handleZipFile(event) {
  debugger;
}
