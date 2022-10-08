import autoprefixer from "autoprefixer";
import {resolve} from "path";
import tailwindcss from "tailwindcss";

export default {
  plugins: [tailwindcss(resolve(__dirname, "tailwind.config.ts")), autoprefixer, require("autoprefixer")],
};
