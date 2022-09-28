import {styled} from "@mui/system";
import {useFrameless} from "helpers/useFrameless";

const DragBar = styled("div")`
  --wails-draggable: drag;
  width: 100%;
  height: 30px;
`;

export const TopMenuBar = (): JSX.Element | null => {
  const frameless = useFrameless();

  if (!frameless) {
    return null;
  }
  return <DragBar />;
};