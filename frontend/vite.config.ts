import postcssConfig from "./postcss.config";
import react from "@vitejs/plugin-react";
import {resolve} from "path";
import {defineConfig} from "vite";
import tsconfigPaths from "vite-tsconfig-paths";

// https://vitejs.dev/config/
export default defineConfig({
  root: __dirname,
  build: {
    emptyOutDir: true,
    outDir: resolve(__dirname, "user", "dist"),
  },
  css: {
    postcss: postcssConfig,
  },
  server: {
    open: true,
    fs: {
      strict: false,
    },
  },
  optimizeDeps: {
    include: ["react-dom/client"],
  },
  plugins: [
    react(),
    tsconfigPaths({
      projects: [resolve(__dirname, "tsconfig.vite.json")],
    }),
  ],
});
