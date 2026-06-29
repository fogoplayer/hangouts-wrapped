import { LitElement, html, css } from "../libs/lit-all@2.7.6.js";
import globalCss from "../global-styles/global.css.mjs";
import "../services/analysis/analysis.mjs";
import {
  getApplicationPhase,
  getIngestStats,
  selectDirectoryForAnalysis,
} from "../services/analysis/analysis.mjs";
import { documentJsonFile } from "../services/JsonDocumenter.mjs";
import "../components/Chart.mjs";
import "@material/mwc-button";
/** @typedef {import("../libs/chart@4.5.0.js").ChartConfiguration} ChartConfiguration */
/** @typedef {import("../services/analysis/analysis.js").ReportData} ReportData */

export default class Home extends LitElement {
  static properties = {
    progress: { type: Object, state: true, default: undefined },
    applicationPhase: { type: String, state: true },
    showReports: { type: String, state: true, default: false },
    results: { type: Array, state: true, default: [] },
  };

  /** @type {NodeJS.Timeout | undefined} */
  statsInterval = undefined;

  async selectFile() {
    const fileHandles = await showOpenFilePicker({});
    fileHandles.forEach(documentJsonFile);
  }

  async selectDirectory() {
    selectDirectoryForAnalysis();
  }

  connectedCallback() {
    super.connectedCallback();

    const phaseState = getApplicationPhase();
    phaseState.onChange(() => {
      const applicationPhase = getApplicationPhase().value;

      switch (applicationPhase) {
        case "WaitingForDirectory":
          clearInterval(this.statsInterval);
          break;

        case "Ingesting":
          this.statsInterval = setInterval(() => {
            this.progress = getIngestStats(); // TODO not working
          }, 50);
          break;
      }
    });
  }

  disconnectedCallback() {
    super.connectedCallback();
    clearInterval(this.statsInterval);
  }

  render() {
    return html`
      <mwc-button @click=${this.selectDirectory} raised
        >Select Directory</mwc-button
      >
      <!-- <mwc-button @click=${this
        .selectFile} raised>Select file</mwc-button> -->
      <output>${this.progress?.toString()}</output>
    `;
  }

  static styles = [
    globalCss,
    css`
      :host {
        flex: 1;

        display: flex;
        flex-flow: column nowrap;
        align-items: center;
        justify-content: center;
      }

      output {
        margin-block: 0.5em;
        white-space: pre-wrap;
        text-align: center;
      }
    `,
  ];
}

customElements.define("home-", Home);
