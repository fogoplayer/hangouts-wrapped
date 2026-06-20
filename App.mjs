import "./libs/pwaupdate.js";
import { css, html, LitElement } from "./libs/lit-all@2.7.6.js";
import globalCss from "./global-styles/global.css.mjs";

import Home from "./pages/Home.mjs";
import { ReportsPage } from "./pages/ReportsPage.mjs";
import {
  getApplicationPhase,
  getReports,
} from "./services/analysis/analysis.mjs";

/** @typedef {import("./services/analysis/analysis.js").ReportData} ReportData */

// Add global styles to head for resets and fonts
const style = document.createElement("style");
style.textContent = globalCss.cssText;
document.head.appendChild(style);

export default class App extends LitElement {
  static properties = {
    currentPage: { type: Object, state: true },
    applicationPhase: { type: String, state: true },
  };

  constructor() {
    super();
    this.setBaseUrl();
    this.createRoute("/", Home);
    this.createRoute("/reports", ReportsPage);
    page.start();
  }

  setBaseUrl() {
    const pathFraments = location.pathname.split("/");
    page.base("/hangouts-wrapped");
  }

  connectedCallback() {
    super.connectedCallback();

    this.reports = getReports();
    /** @type {ReportData[]} */ this.results = [];

    const phaseState = getApplicationPhase();
    this.applicationPhase = phaseState.value;
    phaseState.onChange(() => {
      this.applicationPhase = getApplicationPhase().value;
    });
  }

  /** @param {Map<string, boolean>} changedProperties  */
  updated(changedProperties) {
    if (changedProperties.get("applicationPhase")) {
      switch (this.applicationPhase) {
        case "Ingesting":
          this.navigate("/");
          break;

        case "WaitingForReport":
          this.navigate("/reports");
          break;
      }
    }
  }

  /**
   * Creates a route for the given pattern and associates it with a custom web component.
   *
   * @param {string} pattern the URL pattern to match for the route.
   * @param {new (context: Context) => LitElement} component the component class to be instantiated when the route is activated.
   * @param {string?} title the title to display in the URL bar
   * @returns {void}
   */
  createRoute(pattern, component, title = "hangouts-wrapped") {
    page(pattern, (context) => {
      this.currentPage = new component(context);
    });
  }

  render() {
    return [
      html`<a href="/">
        <mwc-top-app-bar-fixed centerTitle>
          <div slot="title">Hangouts Wrapped</div>
        </mwc-top-app-bar-fixed>
      </a>`,
      this.currentPage,
    ];
  }

  /** @param {string} path  */
  navigate(path) {
    if (path === window.location.pathname) return;
    if ("/hangouts-wrapped" + path === window.location.pathname) return;
    page.show(path);
  }

  static styles = [
    globalCss,
    css`
      :host {
        font-family: var(--sans);
      }

      :host {
        display: flex;
        flex-flow: column nowrap;
        align-items: stretch;

        height: 100dvh;
      }
    `,
  ];
}

customElements.define("app-", App);
