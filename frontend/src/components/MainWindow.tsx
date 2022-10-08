import {Error} from "@lemmy/components/Error";
import {TopMenuBar, TopMenuBarProps} from "@lemmy/components/TopMenuBar";
import {Theme, ThemeProps} from "react-daisyui";
import {ErrorBoundary, ErrorBoundaryProps} from "react-error-boundary";
import {createHashRouter, RouteObject, RouterProvider} from "react-router-dom";

export type WrapperProps = {
  children: JSX.Element[];
};

export type Wrapper = (props: WrapperProps) => JSX.Element;

export type MainWindowProps = {
  theme?: ThemeProps["dataTheme"];
  routes: RouteObject[];
  Wrapper?: Wrapper | string;
  ErrorFallback?: ErrorBoundaryProps["FallbackComponent"];
} & TopMenuBarProps;

export const MainWindow = ({theme, routes, Wrapper = "div", ErrorFallback, getEnvironment}: MainWindowProps): JSX.Element => {
  const router = createHashRouter(routes);

  return (
    <ErrorBoundary FallbackComponent={ErrorFallback ?? Error}>
      <Theme dataTheme={theme}>
        <Wrapper>
          <TopMenuBar getEnvironment={getEnvironment} />

          <ErrorBoundary FallbackComponent={ErrorFallback ?? Error}>
            <RouterProvider router={router} />
          </ErrorBoundary>
        </Wrapper>
      </Theme>
    </ErrorBoundary>
  );
};
