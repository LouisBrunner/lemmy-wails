import Container from "@mui/material/Container"
import CssBaseline from "@mui/material/CssBaseline"
import {createTheme, ThemeProvider} from "@mui/material/styles"
import {HashRouter, Route, Routes} from "react-router-dom"
import {Demo} from "./Demo"

const theme = createTheme()

export const App = (): JSX.Element => {
  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />

        <HashRouter basename={"/"}>
          <Routes>
            <Route path="/" element={<Demo />} />
          </Routes>
        </HashRouter>,
      </Container>
    </ThemeProvider>
  )
}
