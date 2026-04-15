import "../libs/chart@4.5.0.umd.js";
import { createRef, html, LitElement, ref } from "../libs/lit-all@2.7.6.js";

export class ChartComponent extends LitElement {
  static properties = {
    config: { type: Object },
    outputLineCount: { type: Number },
  };

  /** @type {number} */
  outputLineCount = 0;
  /** @type {import("../libs/chart@4.5.0.js").ChartConfiguration | undefined} */
  config = undefined;
  /** @type {import("../libs/lit-all@2.7.6.js").Ref<HTMLCanvasElement>} */
  canvasRef = createRef();

  firstUpdated() {
    this.updateChart();
  }

  /** @param {import("lit").PropertyValues} changedProperties  */
  updated(changedProperties) {
    if (changedProperties.has("config")) {
      this.updateChart();
    }
  }

  updateChart() {
    if (this.config && this.canvasRef.value) {
      this.outputLineCount = this.config.data.labels?.length ?? 0;
      // I don't fully understand why this is is needed in addition to the height, but it is
      this.canvasRef.value.style.height = this.outputLineCount + "em";
      console.log(this.config);
      (async () => new Chart(this.canvasRef.value, this.config))(); // TODO fix type
    }
  }

  render() {
    return html`<canvas
      ${ref(this.canvasRef)}
      height=${this.outputLineCount * 32}
    ></canvas>`;
  }
}

customElements.define("chart-", ChartComponent);
