import {Demo} from "./pages/Demo";
import {TopMenuBar} from "./pages/TopMenuBar";
import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import {createTheme, styled, ThemeProvider} from "@mui/material/styles";
import {Box} from "@mui/system";
import {HashRouter, Route, Routes} from "react-router-dom";

const theme = createTheme();

const MyBox = styled(Box)`
  background: red;
  height: 100vh;
  display: flow-root;
`;

export const App = (): JSX.Element => {
  return (
    <ThemeProvider theme={theme}>
      <MyBox>
        <TopMenuBar />

        <Container component="main" maxWidth="xs">
          <CssBaseline />

          <HashRouter basename={"/"}>
            <Routes>
              <Route path="/" element={<Demo />} />
            </Routes>
          </HashRouter>
        </Container>
      </MyBox>
    </ThemeProvider>
  );
};
