import { Jimp } from 'jimp';
import { promises as fs } from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const inDir = path.join(__dirname, '../../../shared/img/');
const outDir = path.join(__dirname, './out/');

async function resizeImg(filename: string) {
  if (filename === '.DS_Store') {
    return;
  }

  const fullPath = path.join(inDir, filename);
  const outPath = path.join(outDir, filename);

  try {
    const image = await Jimp.read(fullPath);
    image.resize({ w: image.bitmap.width / 2, h: image.bitmap.height / 2 });
    await image.write(`${outPath}.jpeg`);
  } catch (err) {
    console.log('error with: ', filename);
    throw err;
  }
}

async function main() {
  await fs.rm(outDir, { recursive: true, force: true });
  await fs.mkdir(outDir, { recursive: true });

  const files = await fs.readdir(inDir);

  // Sequential
  console.time("Node sequential image resize")
  for (const f of files) {
    await resizeImg(f);
  }
  console.timeEnd("Node sequential image resize")

  console.time("Node parallel image resize")
  await Promise.all(files.map(f => resizeImg(f)));

  console.timeEnd("Node parallel image resize")
}

main().catch(console.error);
