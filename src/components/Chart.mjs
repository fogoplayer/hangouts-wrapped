import "../libs/chart@4.5.0.umd.js";
import { createRef, html, LitElement, ref } from "../libs/lit-all@2.7.6.js";

export class ChartComponent extends LitElement {
  static properties = {
    config: { type: Object },
  };

  /** @type {import("../libs/chart@4.5.0.js").ChartConfiguration | undefined} */
  config = undefined;
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
    if (this.config) {
      (async () => new Chart(this.canvasRef.value, this.config))();
    }
  }

  render() {
    return html`<canvas ${ref(this.canvasRef)}></canvas>`;
  }
}

customElements.define("chart-", ChartComponent);
