import tailwindConfig from "./tailwind.config";
import autoprefixer from "autoprefixer";
import tailwindcss, {Config} from "tailwindcss";

export default {
  plugins: [tailwindcss(tailwindConfig as unknown as Config), autoprefixer],
};
