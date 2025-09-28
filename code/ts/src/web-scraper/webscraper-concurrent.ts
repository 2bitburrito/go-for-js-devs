import * as https from "https";
import fs from "fs"

function fetchURL(url: string) {
  return new Promise((resolve, reject) => {
    const start = Date.now();
    https.get(url, (res) => {
      let data = '';
      res.on('data', chunk => data += chunk);
      res.on('end', () => {
        console.log(`Done! ${url} (${Date.now() - start}ms)`);
        resolve(data);
      });
    }).on('error', reject);
  });
}

async function scrapeAll() {
  const filepath = "../../../shared/urls.json"
  const jsonFile = fs.readFileSync(filepath)
  const urls: string[] = await JSON.parse(jsonFile.toString())

  console.time('Node.js Concurrent');

  // This is the "complex" part in Node.js
  const promises = urls.map(url => fetchURL(url));
  await Promise.all(promises);

  console.timeEnd('Node.js Concurrent');
}

scrapeAll();
