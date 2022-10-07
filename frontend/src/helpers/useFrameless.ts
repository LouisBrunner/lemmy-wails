import {useEffect, useState} from "react";

export type Environment = {
  platform: string;
};

export type useFramelessProps = {
  getEnvironment: () => Promise<Environment>;
};

export const useFrameless = ({getEnvironment}: useFramelessProps): boolean => {
  const [frameless, setFrameless] = useState(false);

  useEffect(() => {
    // eslint-disable-next-line @typescript-eslint/no-floating-promises
    (async (): Promise<void> => {
      const env = await getEnvironment();
      setFrameless(env.platform == "darwin");
    })();
  }, [setFrameless]);

  return frameless;
};
