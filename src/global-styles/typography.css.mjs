import { css } from "../libs/lit-all@2.7.6.js";
import materialSymbols from "./fonts/material-symbols/family.css.mjs";

export default css`
  ${materialSymbols}

  html {
    font-family: var(--sans);
  }

  * {
    font: inherit;
  }

  h1 {
    font-family: var(--serif);
    font-size: 1.5em;
    font-weight: 400;
    line-height: 1.3em;
  }

  h2 {
    /* typography */
    font-family: var(--serif);
    font-size: 1.8em;
    font-weight: 400;
    line-height: 1.3em;
  }

  h3 {
    /* typography */
    font-family: var(--serif);
    font-size: 1.3125em;
    font-weight: 400;
    line-height: 1.3em;
  }
`;
