import "../../libs/go-wasm-runtime.js";
import { documentJsonString as documentJsonFunc } from "../JsonDocumenter.mjs";

const go = new Go();
await WebAssembly.instantiateStreaming(
  fetch("services/analysis/analysis.wasm"),
  go.importObject
).then((result) => {
  go.run(result.instance);
});

// TODO wasmexport
// https://go.dev/blog/wasmexport

export function selectDirectoryForAnalysis() {
  window.showWasmDirectoryPicker();
}

export function getIngestStats() {
  return window.getIngestStats();
}

export function getApplicationPhase() {
  return window.getApplicationPhase();
}

/** @returns {string[]} */
export function getReports() {
  return window.getReportsList();
}

/** @param {number} reportEnum */
export function runReport(reportEnum) {
  return window.runReport(reportEnum);
}

export function getStableChatsList() {
  return window.getStableChatsList();
}

/** @param {...number} indexes  */
export function setChatsFilter(...indexes) {
  window.setChatFilter(...indexes);
}

window.documentJson = documentJsonFunc;
