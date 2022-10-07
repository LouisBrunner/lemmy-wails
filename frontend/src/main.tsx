import React from "react";
import {createRoot} from "react-dom/client";

type mainProps = {
  app: () => JSX.Element;
};

export const main = ({app}: mainProps): void => {
  const App = app;
  const container = document.getElementById("root");

  // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
  const root = createRoot(container!);

  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  );
};
