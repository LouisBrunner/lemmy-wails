import daisyui from "daisyui";

export default {
  content: [
    "./index.html",
    "./src/**/*.tsx",
    "../../src/**/*.tsx",
    "node_modules/daisyui/dist/**/*.js",
    "node_modules/react-daisyui/dist/**/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [daisyui],
};
