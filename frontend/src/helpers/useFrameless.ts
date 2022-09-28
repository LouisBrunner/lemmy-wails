import {Environment} from "@wailsjs/runtime/runtime";
import {useEffect, useState} from "react";

export const useFrameless = (): boolean => {
  const [frameless, setFrameless] = useState(false);

  useEffect(() => {
    // eslint-disable-next-line @typescript-eslint/no-floating-promises
    (async (): Promise<void> => {
      const env = await Environment();
      setFrameless(env.platform == "darwin");
    })();
  }, [setFrameless]);

  return frameless;
};
