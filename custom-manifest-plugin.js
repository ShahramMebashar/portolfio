import fs from 'fs';
import path from 'path';

export default function customManifestPlugin() {
  return {
    name: 'vite-plugin-custom-manifest',
    generateBundle(options, bundle) {
      const manifest = {};

      for (const [key, value] of Object.entries(bundle)) {
        if (value.type === 'chunk' || value.type === 'asset') {
          const entry = {
            file: path.relative(options.dir, value.fileName),
            src: key
          };

          if (value.css && value.css.length > 0) {
            entry.css = value.css.map(cssFile => path.relative(options.dir, cssFile));
          }

          manifest[key] = entry;
        }
      }

      fs.writeFileSync(path.resolve(options.dir, 'manifest.json'), JSON.stringify(manifest, null, 2));
    }
  };
}
