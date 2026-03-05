import { LitElement, html, css } from "../libs/lit-all@2.7.6.js";
import globalCss from "../global-styles/global.css.mjs";
import "../services/analysis/analysis.mjs";

export default class Home extends LitElement {
  static properties = {
    progress: { type: Number, state: true },
  };

  constructor() {
    super();
  }

  async selectFile() {
    const fileHandle = await showOpenFilePicker({});
  }

  async selectDirectory() {
    const directoryHandle = await showDirectoryPicker();
    debugger;
  }

  render() {
    return html`<header><h1>hangouts-wrapped</h1></header>
      <main>Welcome to my app!</main>
      <button @click=${this.selectDirectory}>Select Directory</button>
      <button @click=${this.selectFile}>Select file</button>`;
  }

  static styles = [globalCss, css``];
}

customElements.define("home-", Home);
