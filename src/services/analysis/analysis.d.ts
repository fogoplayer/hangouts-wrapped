import { documentJsonString as documentJsonType } from "../JsonDocumenter.mjs";

declare global {
  interface Window {
    MyNamespace: any;
    documentJson: typeof documentJsonType; // used in analysis.mjs to expose JS function to Go
    showWasmDirectoryPicker: () => void;
    getIngestStats(): Record<string, number>;
  }
}
