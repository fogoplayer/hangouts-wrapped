import { html, LitElement, css } from "lit"; // TODO fix import
import {
  getReports,
  getStableChatsList,
  runReport,
  setChatsFilter,
} from "../services/analysis/analysis.mjs";
import "@material/mwc-button";
import "@material/mwc-icon-button";
import "@material/mwc-menu";
import "@material/mwc-select";
import globalCss from "../global-styles/global.css.mjs";
// import "@material/mwc-list-item";
/** @typedef {import("../services/analysis/analysis.js").ReportData} ReportData */

export class ReportsPage extends LitElement {
  static properties = {
    showReports: { type: String, state: true, default: false },
    results: { type: Array, state: true, default: [] },
    selectedReport: { type: Number, state: true, default: 0 },
  };

  connectedCallback() {
    super.connectedCallback();
    this.reports = getReports();
    /** @type {(ReportData | string)[]} */ this.results = [];
    /** @type {number} */ this.selectedReport;
  }

  render() {
    return html`
      <output>
        ${this.results?.map((configOrMessage) =>
          typeof configOrMessage === "string"
            ? html`<div class="request">${configOrMessage}</div>`
            : html`<div class="response">
                <chart- .config=${configOrMessage}></chart->
              </div>`
        )}
      </output>

      <form action="">
        <div class="filters-group">
          <mwc-select
            label="Chats"
            naturalMenuWidth="true"
            value=""
            @selected=${(/** @type {{detail:{index: number}}} */ event) => {
              if (event.detail.index === 0) return;

              setChatsFilter(event.detail.index - 1);
            }}
          >
            <mwc-list-item></mwc-list-item>
            ${getStableChatsList()?.map(
              (chatName, i) =>
                html`<mwc-list-item group="a" graphic="icon" value="i">
                  <span> ${chatName} </span>
                </mwc-list-item> `
            )}
          </mwc-select>
          <mwc-select
            label="Report"
            naturalMenuWidth="true"
            @selected=${(/** @type {{detail:{index: number}}} */ event) => {
              this.selectedReport = event.detail.index;
            }}
          >
            <mwc-list-item></mwc-list-item>
            ${this.reports?.map(
              (description, i) =>
                html`<mwc-list-item group="a" graphic="icon" value="i">
                  <span> ${description} </span>
                </mwc-list-item> `
            )}
          </mwc-select>
        </div>
        <mwc-button
          outlined
          label="Clear History"
          @click=${() => (this.results = [])}
        ></mwc-button>
        <mwc-icon-button
          label="Submit"
          icon="send"
          @click=${() => {
            // offset by one for blank default selection
            if (this.selectedReport === 0) return;

            const reportEnum = this.selectedReport - 1;
            const reportDescription = this.reports?.[reportEnum];
            if (reportDescription) {
              this.results?.push(reportDescription);
            }
            this.results?.push();
            const result = runReport(reportEnum);
            this.results?.push(result);
            this.requestUpdate();
          }}
        ></mwc-icon-button>
      </form>
    `;
  }

  static styles = [
    globalCss,
    css`
      :host {
        flex: 1;

        display: flex;
        flex-flow: column nowrap;

        height: 100dvh;
        overflow: hidden;
      }

      output {
        flex: 1;

        display: flex;
        flex-flow: column nowrap;
        gap: 1em;

        padding: 1em;
        background-color: var(--theme-background);
        overflow-y: auto; // TODO make whole page scrollable with sticky footer
        overflow-x: hidden;

        .request {
          flex: 0 0 auto;
          align-self: end;

          margin-left: 10%;
          padding: 1em;
          background-color: #cfd8e1; /* TODO make variables */
          border-radius: 0.5em 0 0.5em 0.5em;
          width: fit-content;
        }

        .response {
          flex: 0 0 auto;
          align-self: stretch;

          margin-right: 10%;
          padding: 1em;
          overflow: hidden;
          background-color: #feffff;
          border-radius: 0 0.5em 0.5em 0.5em;
        }
      }

      form {
        display: flex;
        flex-flow: row nowrap;
        align-items: center;
        gap: 0.5em;

        position: sticky;
        bottom: 0;

        margin: 0;
        padding: 1em;

        border-top: 1px solid var(--theme-border);

        flex: 0;

        .filters-group {
          flex: 1;

          display: grid;
          grid-template-columns: repeat(auto-fit, minmax(10em, 1fr));
          place-items: stretch;
          gap: 0.5em;
        }

        > * {
          flex-shrink: 0;
        }
      }
    `,
  ];
}

customElements.define("reports-page", ReportsPage);
