import fs from "fs"

function fetchURL(url: string) {
  return new Promise((resolve, reject) => {
    fetch(url, {
      headers: {
        "Accept": "image/*"
      }
    })
      .then((res) => res.ok ? resolve(true) : reject(res.url))
      .catch(() => reject(false))
  });
}

async function scrapeAll() {
  const filepath = "../../../shared/urls.json"
  const jsonFile = fs.readFileSync(filepath)
  const urls: string[] = await JSON.parse(jsonFile.toString())

  console.time('Node.js Concurrent');

  const promises = urls.map(url => fetchURL(url));
  const res = await Promise.allSettled(promises);

  const successCount = res.filter((r) => r.status == "fulfilled")
  console.log(`Scraped a total of ${successCount.length} URLs of ${urls.length}`)
  console.timeEnd('Node.js Concurrent');

}

scrapeAll();
