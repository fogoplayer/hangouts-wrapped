import globalCss from "../global-styles/global.css.mjs";
import "../libs/chart@4.5.0.umd.js";
import {
  createRef,
  html,
  LitElement,
  ref,
  css,
} from "../libs/lit-all@2.7.6.js";

export class ChartComponent extends LitElement {
  static properties = {
    config: { type: Object },
    outputDataPointsCount: { type: Number },
    orientation: { type: String },
  };

  /** @type {number} */
  outputDataPointsCount = 0;
  /** @type {import("../services/analysis/analysis.js").ReportData | undefined} */
  config = undefined;
  /** @type {import("../libs/lit-all@2.7.6.js").Ref<HTMLCanvasElement>} */
  canvasRef = createRef();
  /** @type {"vertical" | "horizontal"} */
  orientation = "vertical";

  firstUpdated() {
    this.updateChart();
  }

  /** @param {import("../libs/lit-all@2.7.6.js").PropertyValues} changedProperties  */
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

      if (this.config.type.toLowerCase() === "line") {
        this.orientation = "horizontal";
      } else {
        this.orientation = "vertical";
      }

      this.outputDataPointsCount = this.config.data.labels?.length ?? 0;
      // I don't fully understand why this is is needed in addition to the height, but it is
      if (this.orientation === "vertical") {
        this.canvasRef.value.style.height = this.outputDataPointsCount + "em";
      } else {
        this.canvasRef.value.style.height = "30em"; // TODO better solution here
        this.canvasRef.value.style.width = this.outputDataPointsCount + "em";
      }

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
      : this.orientation === "vertical"
        ? // TODO make responsive
          // TODO overflow:auto wrapper around fixed size chart
          html`<canvas
            ${ref(this.canvasRef)}
            height=${this.outputDataPointsCount * 32}
          ></canvas>`
        : html`<canvas
            ${ref(this.canvasRef)}
            width=${this.outputDataPointsCount * 32}
          ></canvas>`;
  }

  static styles = [
    globalCss,
    css`
      table tr:nth-child(even) {
        background-color: var(--theme-background);
      }
    `,
  ];
}

customElements.define("chart-", ChartComponent);
