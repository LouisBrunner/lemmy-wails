import tailwindConfig from "./tailwind.config";
import autoprefixer from "autoprefixer";
import tailwindcss from "tailwindcss";

export default {
  plugins: [tailwindcss(tailwindConfig), autoprefixer],
};
