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

export default class Home extends LitElement {
  static properties = {
    progress: { type: Object, state: true, default: undefined },
    applicationPhase: { type: String, state: true },
    showReports: { type: String, state: true, default: false },
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
                html`<button @click=${() => runReport(i)}>
                  ${description}
                </button>`
            )
          : ""}
      </div>`;
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
