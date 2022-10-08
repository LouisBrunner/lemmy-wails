import {useFrameless, useFramelessProps} from "@lemmy/helpers/useFrameless";

export type TopMenuBarProps = useFramelessProps;

export const TopMenuBar = ({getEnvironment}: TopMenuBarProps): JSX.Element | null => {
  const frameless = useFrameless({getEnvironment});

  if (!frameless) {
    return null;
  }
  return <div className="p-[30px] h-full wails-drag" />;
};
