interface Context<ParamObject = undefined> {
  save(): void;
  handled: boolean;
  canonicalPath: string;
  path: string;
  querystring: string;
  pathname: string;
  state: any;
  title: string;
  params: ParamObject;
}

type Callback = (context: Context, next?: Function) => void;

interface Options {
  click?: boolean;
  popstate?: boolean;
  dispatch?: boolean;
}

interface PageFunction {
  (path: string, callback?: Callback, ...callbacks: Callback[]): void;
  (callback: Callback): void;
  (fromPath: string, toPath: string): void;
  redirect(fromPath: string, toPath: string): void;
  redirect(path: string): void;
  (options: Options): void;
  start(options?: Options): void;
}

declare const page: PageFunction;
