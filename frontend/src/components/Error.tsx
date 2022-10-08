import {FallbackProps} from "react-error-boundary";

export type ErrorProps = FallbackProps;

export const Error = ({error}: ErrorProps): JSX.Element => {
  return <div>{error.message}</div>;
};
