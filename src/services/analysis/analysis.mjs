import "../../libs/go-wasm-runtime.js";
import { documentJsonString as documentJsonFunc } from "../JsonDocumenter.mjs";

const go = new Go();
await WebAssembly.instantiateStreaming(
  fetch("/services/analysis/analysis.wasm"),
  go.importObject
).then((result) => {
  go.run(result.instance);
});

export function selectDirectoryForAnalysis() {
  window.showWasmDirectoryPicker();
}

export function getIngestStats() {
  return window.getIngestStats();
}

export function getApplicationPhase() {
  return window.getApplicationPhase();
}

window.documentJson = documentJsonFunc;
