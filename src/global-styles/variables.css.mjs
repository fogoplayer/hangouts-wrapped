import { css } from "../libs/lit-all@2.7.6.js";

export default css`
  :host {
    /* Theme colors */
    --theme-primary: #109d59;
    --theme-secondary: darkgray;
    --theme-tertiary: gray --border: 1px solid var(--theme-primary);
    --theme-border: #e4e4e4;
    --box-shadow: rgba(0, 0, 0, 0.5) 0px 1px 1px 0px;
    --theme-background: #eceff4;

    /* Fonts */
    --serif: serif;
    --sans:
      "Roboto", -apple-system, "BlinkMacSystemFont", "Segoe UI", "Lato",
      "Helvetica", "Arial", sans-serif;

    /* Units */
    --limited-width: 40rem;

    /* MDC */
    --mdc-theme-primary: var(--theme-primary);
    --mdc-theme-background: var(--theme-background);
    --mdc-icon-font: "Material Symbols Sharp";
  }
`;
