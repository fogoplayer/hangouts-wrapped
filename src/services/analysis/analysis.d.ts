import { documentJsonString as documentJsonType } from "../JsonDocumenter.mjs";

declare global {
  interface Window {
    MyNamespace: any;
    documentJson: typeof documentJsonType; // used in analysis.mjs to expose JS function to Go
    showWasmDirectoryPicker: () => void;
    // It is VERY important that the return value is readonly
    // See IngestStatsType.String() for more details
    getIngestStats(): Readonly<Record<string, number>>;
  }
}
