import { documentJsonString as documentJsonType } from "../JsonDocumenter.mjs";

declare global {
  interface Window {
    MyNamespace: any;
    documentJson: typeof documentJsonType;
    showWasmDirectoryPicker: () => void;
  }
}
