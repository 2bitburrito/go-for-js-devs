import { readdir } from "fs/promises";
import { Worker } from "worker_threads";
import * as path from "path";

const dirPath = "code/shared/audio/";

function runWorker(filePath: string, fileName: string): Promise<void> {
  return new Promise((resolve, reject) => {
    const worker = new Worker("./code/ts/src/audio-metadata/worker.ts", {
      workerData: { filePath, fileName },
    });

    worker.on("message", (msg) => {
      if (msg === "done") resolve();
    });

    worker.on("error", reject);
    worker.on("exit", (code) => {
      if (code !== 0) reject(new Error(`Worker stopped with code ${code}`));
    });
  });
}

async function main() {
  const entries = await readdir(dirPath);
  const start = Date.now();

  // Run workers in parallel
  await Promise.all(
    entries.map((entry) => runWorker(path.join(dirPath, entry), entry))
  );

  console.log("Time taken with workers:", Date.now() - start, "ms");
}

main().catch(console.error);
