/** @typedef {typeof FileReader.prototype.result} FileReaderResult */

export class AsyncFileReader extends FileReader {
  /**
   * @param {Blob} blob
   * @return {Promise<ArrayBuffer>}
   */
  async readAsArrayBufferAsync(blob) {
    this.readAsArrayBuffer(blob);
    return new Promise((res, rej) => {
      this.addEventListener("error", (event) =>
        rej(event.target?.error?.message)
      );
      this.addEventListener("loadend", (event) => {
        debugger;
        if (this.result instanceof ArrayBuffer) {
          res(this.result);
        } else {
          rej("result of wrong type");
        }
      });
    });
  }
}
