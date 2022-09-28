import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import {Greet} from "@wailsjs/go/bindings/bindings";
import {ChangeEvent, FormEvent, useCallback, useState} from "react";

export const Demo = (): JSX.Element => {
  const [resultText, setResultText] = useState("Please enter your name below 👇");
  const [name, setName] = useState("");

  const updateName = useCallback(
    (e: ChangeEvent<HTMLInputElement>): void => {
      setName(e.target.value);
    },
    [setName],
  );

  const greet = useCallback(
    async (e: FormEvent<HTMLFormElement>): Promise<void> => {
      e.preventDefault();
      const greeting = await Greet(name);
      setResultText(greeting);
    },
    [name, setResultText],
  );

  return (
    <Box
      sx={{
        marginTop: 8,
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography component="h1" variant="h5">
        Greetings
      </Typography>
      <Box component="form" onSubmit={greet} noValidate sx={{mt: 1}}>
        <Typography>{resultText}</Typography>
        <TextField margin="normal" required fullWidth id="name" label="Name" name="name" autoFocus onChange={updateName} />
        <Button type="submit" fullWidth variant="contained" sx={{mt: 3, mb: 2}}>
          Greet
        </Button>
      </Box>
    </Box>
  );
};
