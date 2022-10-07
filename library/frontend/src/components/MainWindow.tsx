import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import {Theme, ThemeProvider} from "@mui/material/styles";
import {TopMenuBar, TopMenuBarProps} from "components/TopMenuBar";
import {createHashRouter, RouteObject, RouterProvider} from "react-router-dom";

type WrapperProps = {
  children: JSX.Element[];
};

type Wrapper = (props: WrapperProps) => JSX.Element;

type MainWindowProps = {
  theme: Theme;
  routes: RouteObject[];
  wrapper: Wrapper;
} & TopMenuBarProps;

export const MainWindow = ({theme, routes, wrapper, getEnvironment}: MainWindowProps): JSX.Element => {
  const router = createHashRouter(routes);
  const Wrapper = wrapper;

  return (
    <ThemeProvider theme={theme}>
      <Wrapper>
        <TopMenuBar getEnvironment={getEnvironment} />

        <Container component="main">
          <CssBaseline />

          <RouterProvider router={router} />
        </Container>
      </Wrapper>
    </ThemeProvider>
  );
};
