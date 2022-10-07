import { MainWindow } from "lemmy-wails/components/MainWindow";
import { createTheme, styled } from "@mui/material/styles";
import { Box } from "@mui/system";
import { Demo } from "pages/Demo";

const theme = createTheme();

const MyBox = styled(Box)`
  background: red;
  height: 100vh;
  display: flow-root;
`;

export const App = (): JSX.Element => {
  return (
    <MainWindow
      theme={theme}
      wrapper={MyBox}
      routers={[
        {
          path: "/",
          element: <Demo />,
        },
      ]}
      getEnvironment={Environment}
    />
  );
};
