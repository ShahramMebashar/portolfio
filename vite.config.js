import { defineConfig } from "vite";

export default defineConfig({
  publicDir: "./public/static",
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
        assetFileNames: "css/[name].[hash].css",
      },
    },
  },
});
