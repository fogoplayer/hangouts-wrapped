import { LitElement, html, css } from "../libs/lit-all@2.7.6.js";
import globalCss from "../global-styles/global.css.mjs";
import "../services/analysis/analysis.mjs";
import {
  getApplicationPhase,
  getIngestStats,
  selectDirectoryForAnalysis,
} from "../services/analysis/analysis.mjs";
import { documentJsonFile } from "../services/JsonDocumenter.mjs";

export default class Home extends LitElement {
  static properties = {
    progress: { type: Object, state: true, default: undefined },
  };

  constructor() {
    super();
  }

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

    const phaseState = getApplicationPhase();
    phaseState.onChange(() => {
      const secondPhase = getApplicationPhase();
      switch (secondPhase.value) {
        case "Ingesting":
          this.statsInterval = setInterval(() => {
            this.progress = getIngestStats();
          }, 50);
        default:
          clearInterval(this.statsInterval);
      }
    });
  }

  disconnectedCallback() {
    super.connectedCallback();
    clearInterval(this.statsInterval);
  }

  render() {
    return html`<header><h1>hangouts-wrapped</h1></header>
      <main>Welcome to my app!</main>
      <button @click=${this.selectDirectory}>Select Directory</button>
      <button @click=${this.selectFile}>Select file</button>
      <output>${this.progress?.toString()}</output>`;
  }

  static styles = [globalCss, css``];
}

customElements.define("home-", Home);
