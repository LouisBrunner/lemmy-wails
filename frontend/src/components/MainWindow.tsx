import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import {Theme, ThemeProvider} from "@mui/material/styles";
import {Error} from "components/Error";
import {TopMenuBar, TopMenuBarProps} from "components/TopMenuBar";
import {ErrorBoundary, ErrorBoundaryProps} from "react-error-boundary";
import {createHashRouter, RouteObject, RouterProvider} from "react-router-dom";

export type WrapperProps = {
  children: JSX.Element[];
};

export type Wrapper = (props: WrapperProps) => JSX.Element;

export type MainWindowProps = {
  theme: Theme;
  routes: RouteObject[];
  Wrapper: Wrapper;
  ErrorFallback: ErrorBoundaryProps["FallbackComponent"];
} & TopMenuBarProps;

export const MainWindow = ({theme, routes, Wrapper, ErrorFallback, getEnvironment}: MainWindowProps): JSX.Element => {
  const router = createHashRouter(routes);

  return (
    <ThemeProvider theme={theme}>
      <ErrorBoundary FallbackComponent={ErrorFallback ?? Error}>
        <Wrapper>
          <TopMenuBar getEnvironment={getEnvironment} />

          <Container component="main">
            <CssBaseline />

            <RouterProvider router={router} />
          </Container>
        </Wrapper>
      </ErrorBoundary>
    </ThemeProvider>
  );
};
