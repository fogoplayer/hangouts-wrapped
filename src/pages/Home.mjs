import { LitElement, html, css } from "../libs/lit-all@2.7.6.js";
import globalCss from "../global-styles/global.css.mjs";
import { handleZipFile } from "../services/analysis/analysis.mjs";
import { AsyncFileReader } from "../services/AysncFileReader.mjs";

export default class Home extends LitElement {
  static properties = {
    progress: { type: Number, state: true },
  };

  constructor() {
    super();
  }

  /** @param {Event} event */
  async onFileChange(event) {
    const input = event.currentTarget;
    if (!(input instanceof HTMLInputElement)) {
      return;
    }

    const files = input.files ?? [];
    const reader = new AsyncFileReader();
    reader.addEventListener(
      "progress",
      (e) => (this.progress = (e.loaded / e.total) * 100)
    );
    try {
      const buffer = await reader.readAsArrayBufferAsync(files[0]);
    } catch (error) {
      console.error(error);
    }
    debugger;

    handleZipFile(event);
  }

  render() {
    return html`<header><h1>hangouts-wrapped</h1></header>
      <main>Welcome to my app!</main>
      <input
        type="file"
        name="file"
        id="file-input"
        @change="${this.onFileChange}"
      />
      Progress: ${this.progress}%`;
  }

  static styles = [globalCss, css``];
}

customElements.define("home-", Home);
