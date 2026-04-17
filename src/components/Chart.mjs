import "../libs/chart@4.5.0.umd.js";
import { createRef, html, LitElement, ref } from "../libs/lit-all@2.7.6.js";

export class ChartComponent extends LitElement {
  static properties = {
    config: { type: Object },
    outputLineCount: { type: Number },
  };

  /** @type {number} */
  outputLineCount = 0;
  /** @type {import("../services/analysis/analysis.js").ReportData | undefined} */
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
      if (this.config.type.toLowerCase() === "text") {
        return;
      }

      this.outputLineCount = this.config.data.labels?.length ?? 0;
      // I don't fully understand why this is is needed in addition to the height, but it is
      this.canvasRef.value.style.height = this.outputLineCount + "em";
      (async () => new Chart(this.canvasRef.value, this.config))(); // TODO fix type
    }
  }

  render() {
    return this.config?.type === "text"
      ? // TODO style
        // TODO line-break spaces between slashes
        html`<table>
          ${this.config.data.labels.map(
            (label, i) =>
              html`<tr>
                <th>${label}</th>
                ${this.config?.data.datasets.map(
                  ({ data }) => html`<td>${data[i]}</td>`
                )}
              </tr>`
          )}
        </table>`
      : html`<canvas
          ${ref(this.canvasRef)}
          height=${this.outputLineCount * 32}
        ></canvas>`;
  }
}

customElements.define("chart-", ChartComponent);
