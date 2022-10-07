import {styled} from "@mui/system";
import {useFrameless, useFramelessProps} from "helpers/useFrameless";

const DragBar = styled("div")`
  --wails-draggable: drag;
  width: 100%;
  height: 30px;
`;

export type TopMenuBarProps = useFramelessProps;

export const TopMenuBar = ({getEnvironment}: TopMenuBarProps): JSX.Element | null => {
  const frameless = useFrameless({getEnvironment});

  if (!frameless) {
    return null;
  }
  return <DragBar />;
};
