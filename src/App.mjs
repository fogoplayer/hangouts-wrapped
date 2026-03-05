import "./libs/pwaupdate.js";
import { css, LitElement } from "./libs/lit-all@2.7.6.js";
import globalCss from "./global-styles/global.css.mjs";

import Home from "./pages/Home.mjs";

// Add global styles to head for resets and fonts
const style = document.createElement("style");
style.textContent = globalCss.cssText;
document.head.appendChild(style);

export default class App extends LitElement {
  static properties = {
    currentPage: { type: Object, state: true },
  };

  constructor() {
    super();
    this.createRoute("/", Home);
    page.start();
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
    return this.currentPage;
  }

  static styles = [globalCss, css``];
}

customElements.define("app-", App);
