import React from "react";
import {createRoot} from "react-dom/client";

export type mainProps = {
  App: () => JSX.Element;
};

export const main = ({App}: mainProps): void => {
  const container = document.getElementById("root");

  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const root = createRoot(container!);

  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  );
};
