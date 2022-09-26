import {useState} from 'react';
import {Greet} from "../wailsjs/go/bindings/bindings";
import Box from '@mui/material/Box';
import CssBaseline from '@mui/material/CssBaseline';
import {createTheme, ThemeProvider} from '@mui/material/styles';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';

const theme = createTheme();

export const App = () => {
  const [resultText, setResultText] = useState("Please enter your name below 👇");
  const [name, setName] = useState('');
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  const greet = async (e: Event) => {
    e.preventDefault()
    const greeting = await Greet(name)
    updateResultText(greeting);
  }

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Typography component="h1" variant="h5">
            Greetings
          </Typography>
          <Box component="form" onSubmit={greet} noValidate sx={{mt: 1}}>
            <Typography>
              {resultText}
            </Typography>
            <TextField
              margin="normal"
              required
              fullWidth
              id="name"
              label="Name"
              name="name"
              autoFocus
              onChange={updateName}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{mt: 3, mb: 2}}
            >
              Greet
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  )
}
