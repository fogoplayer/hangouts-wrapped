import { LitElement, html, css } from "../libs/lit-all@2.7.6.js";
import globalCss from "../global-styles/global.css.mjs";
import "../services/analysis/analysis.mjs";
import {
  getApplicationPhase,
  getIngestStats,
  getReports,
  runReport,
  selectDirectoryForAnalysis,
} from "../services/analysis/analysis.mjs";
import { documentJsonFile } from "../services/JsonDocumenter.mjs";
import "../components/Chart.mjs";
/** @typedef {import("../libs/chart@4.5.0.js").ChartConfiguration} ChartConfiguration */

export default class Home extends LitElement {
  static properties = {
    progress: { type: Object, state: true, default: undefined },
    applicationPhase: { type: String, state: true },
    showReports: { type: String, state: true, default: false },
    results: { type: Array, state: true, default: [] },
  };

  /** @type {NodeJS.Timeout | undefined} */ // TODO ? is null in this repo, for some reason?
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

    this.reports = getReports();
    /** @type {ChartConfiguration[]} */ this.results = [];

    const phaseState = getApplicationPhase();
    this.applicationPhase = phaseState.value;
    phaseState.onChange(() => {
      this.applicationPhase = getApplicationPhase().value;
    });
  }

  disconnectedCallback() {
    super.connectedCallback();
    clearInterval(this.statsInterval);
  }

  /** @param {Map<string, boolean>} changedProperties  */
  updated(changedProperties) {
    if (changedProperties.get("applicationPhase")) {
      if (this.applicationPhase === "Ingesting") {
        // TODO this currently does nothing because ingesting blocks the main thread. Seek ways around this.
        this.statsInterval = setInterval(() => {
          this.progress = getIngestStats();
        }, 50);
      } else clearInterval(this.statsInterval);

      if (this.applicationPhase === "WaitingForReport") {
        this.showReports = true;
      } else this.showReports = false;
    }
  }

  render() {
    return html`<header><h1>hangouts-wrapped</h1></header>
      <main>Welcome to my app!</main>
      <button @click=${this.selectDirectory}>Select Directory</button>
      <button @click=${this.selectFile}>Select file</button>
      <output>${this.progress?.toString()}</output>
      <div>
        ${this.showReports
          ? this.reports?.map(
              (description, i) =>
                html`<button @click=${() => this.results?.push(runReport(i))}>
                  ${description}
                </button>`
            )
          : ""}
      </div>
      <button @click=${() => (this.results = [])}>Clear</button>
      <!-- <div>${this.results?.map((config) =>
        JSON.stringify(config)
      )}</div> -->
      <div>
        ${this.results?.map(
          (config) => html`<chart- .config=${config}></chart->`
        )}
      </div> `;
  }

  static styles = [
    globalCss,
    css`
      button {
        outline: 1px solid;
      }
    `,
  ];
}

customElements.define("home-", Home);
