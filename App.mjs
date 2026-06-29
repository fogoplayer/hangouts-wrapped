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
    openHelpDialog: { type: Boolean, state: true },
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
    /** @type {boolean} */ this.openHelpDialog;
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
            <mwc-icon-button
              icon="help"
              slot="actionItems"
              @click=${(e) => {
                e.preventDefault();
                this.openHelpDialog = true;
              }}
            ></mwc-icon-button>
          </mwc-top-app-bar-fixed>
        </a>
        <!-- TODO move into its own component -->
        <mwc-dialog
          id="dialog"
          aria-label="How to use"
          .open=${this.openHelpDialog}
          @closed=${() => (this.openHelpDialog = false)}
        >
          <h2>How to use</h2>
          <ol>
            <li>
              Create an export of your Google Chat (Hangouts) data using
              <a href="http://takeout.google.com">Google Takeout</a>.
              <details>
                <summary>Google Takeout Instructions</summary>
                <ol>
                  <li>Deselect all products</li>
                  <li>Select only "Google Chat"</li>
                  <li>Scroll to the bottom and click "next step"</li>
                  <li>
                    Choose your desired settings. I usually go for
                    <ul>
                      <li>
                        <span class="setting-name">Transfer to:</span> Send
                        download link via email
                      </li>
                      <li>
                        <span class="setting-name">Frequency:</span> Export once
                      </li>
                      <li><span class="setting-name">File type:</span> .zip</li>
                      <li><span class="setting-name">File size:</span> 50GB</li>
                    </ul>
                  </li>
                  <li>
                    Click "create export"
                    <aside>
                      It will probably take about 24 hours for the export to be
                      ready.
                    </aside>
                  </li>
                </ol>
              </details>
            </li>
            <li>
              Unzip the export. If it came in multiple files, combine them into
              one folder, so that it's structured like this:
              <pre>
Takeout/
└─ Google Chat/
   ├─ Groups/
   │  ├─ Space AAAA.../
   │  └─ DM AAAA..../
   └─ Users/
</pre
              >
            </li>
            <li>
              Select your Takeout directory, and enjoy your Hangouts Wrapped!
            </li>
          </ol>
          <mwc-button slot="primaryAction" dialogAction="close">
            Got It
          </mwc-button>
        </mwc-dialog> `,
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

      .setting-name {
        font-weight: bold;
      }

      pre {
        font-size: 0.75em;
        line-height: 1.25em;
      }
    `,
  ];
}

customElements.define("app-", App);
