import { defineConfig } from "vite";
import tailwindcss from "tailwindcss";
import autoprefixer from "autoprefixer";
import customManifestPlugin from "./custom-manifest-plugin";

export default defineConfig({
  publicDir: "./public/",
  plugins: [customManifestPlugin()],
  // css
  css: {
    postcss: {
      plugins: [tailwindcss, autoprefixer],
    },
  },
  server: {
    port: 5173,
    strictPort: true, // Fail if the port is already in use
  },
  build: {
    manifest: true,
    dir: "./web",
    outDir: "./public",
    rollupOptions: {
      input: {
        main: "./web/js/main.js",
      },
      output: {
        entryFileNames: "js/[name].[hash].js",
        chunkFileNames: "js/[name].[hash].js",
        assetFileNames: "css/[name].[hash].[ext]",
      },
    },
  },
});
