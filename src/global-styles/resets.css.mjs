import { css } from "../libs/lit-all@2.7.6.js";

export default css`
  :where(*) {
    box-sizing: border-box;
    margin: 0;
    background-color: inherit;
  }

  @media screen and (max-width: 62em) {
    :root {
      font-size: 1.15em;
    }
  }

  :where(body) {
    padding: 0;
    margin: 0;
    background-color: white;
  }

  :where(.unstyled-ul) {
    padding: unset;
    list-style-type: none;
  }

  :where(.prevent-select) {
    -webkit-user-select: none; /* Safari */
    -ms-user-select: none; /* IE 10 and IE 11 */
    user-select: none; /* Standard syntax */
  }

  :where(button) {
    background-color: transparent;
    border-width: 0;
    font-family: inherit;
    font-size: inherit;
    font-style: inherit;
    font-weight: inherit;
    line-height: inherit;
    text-align: inherit;
    padding: 0;
    min-height: 0;
    max-width: unset;
  }
`;
