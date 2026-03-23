import { documentJsonString as documentJsonType } from "../JsonDocumenter.mjs";

declare global {
  interface Window {
    MyNamespace: any;
    documentJson: typeof documentJsonType; // used in analysis.mjs to expose JS function to Go
    showWasmDirectoryPicker(): void;
    // It is VERY important that the return value is readonly
    // See IngestStatsType.String() for more details
    getIngestStats(): Readonly<Record<string, number>>;
    getApplicationPhase(): Readonly<{
      value: string; // actually an enum, but since it's readonly I think it's okay to be more general until we have codegen that can define TS enums from Go enums
      onChange(callback: () => void): void;
    }>;
    getReportsList(): string[];
    runReport(reportEnum: number): unknown;
  }
}
