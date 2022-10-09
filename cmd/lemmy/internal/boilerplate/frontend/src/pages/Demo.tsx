import {Greet} from "@wailsjs/go/bindings/bindings";
import {ChangeEvent, FormEvent, useCallback, useState} from "react";
import {Button, Input} from "react-daisyui";

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
    <div className="mt-4 flex flex-col items-center">
      <h1 className="text-3xl my-4">Greetings</h1>
      <form className="mt-1" onSubmit={greet} noValidate>
        <p>{resultText}</p>
        <div className="form-control w-full max-w-xs">
          <label className="label" htmlFor="name">
            <span className="label-text">Name</span>
          </label>
          <Input id="name" name="name" required onChange={updateName} />
        </div>
        <Button type="submit" fullWidth className="mt-3 mb-2">
          Greet
        </Button>
      </form>
    </div>
  );
};
