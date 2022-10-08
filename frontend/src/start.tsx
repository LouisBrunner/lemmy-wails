import React from "react";
import {createRoot} from "react-dom/client";

export type startProps = {
  App: () => JSX.Element;
};

export const start = ({App}: startProps): void => {
  const container = document.getElementById("root");

  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const root = createRoot(container!);

  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  );
};
