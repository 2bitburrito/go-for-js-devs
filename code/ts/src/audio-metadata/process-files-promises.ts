import { promises as fs } from "fs";
import * as path from "path";
import { XMLParser } from "fast-xml-parser";
import process from "process";

interface ChunkHeader {
  id: string;
  size: number;
}

async function main() {
  const dirPath = "../../../shared/audio/";
  const entries = await fs.readdir(dirPath);
  const start = Date.now();

  await Promise.all(
    entries.map(async (entry) => {
      const fullPath = path.join(dirPath, entry);
      await getIxmlMetadata(fullPath, entry);
    })
  );

  console.log("Time taken for Node implementation:", Date.now() - start, "ms");
  console.log("Current Memory Allocation: ", process.memoryUsage().heapTotal / 1024 / 1024, "MB")
}

async function getIxmlMetadata(filepath: string, fileName: string) {
  const fd = await fs.open(filepath, "r");

  try {
    // Skip RIFF header (12 bytes: "RIFF" + fileSize + "WAVE")
    let position = 12;

    while (true) {
      const headerBuf = Buffer.alloc(8);
      const { bytesRead } = await fd.read(headerBuf, 0, 8, position);
      if (bytesRead < 8) {
        console.log("No XML for", filepath);
        break;
      }

      const id = headerBuf.toString("ascii", 0, 4);
      const size = headerBuf.readUInt32LE(4);

      position += 8;

      if (id === "iXML") {
        const dataBuf = Buffer.alloc(size);
        await fd.read(dataBuf, 0, size, position);

        const xml = dataBuf.toString("utf8").replace(/\x00+$/, "");
        await fs.writeFile(`./outputs/${fileName}.xml`, xml, "utf8");
        const ixml = parseIxml(xml);
        if (ixml instanceof Error) {
          console.error("ERROR while parsing iXML for: ", fileName, ixml.message);
        } else {
          const tc = ixml.convertSamplesToTimecode();
          console.log("TIMECODE: ", tc);
        }

        break;
      } else {
        let skip = size;
        if (skip % 2 === 1) skip++;
        position += skip;
      }
    }
  } finally {
    await fd.close();
  }
}

interface TimecodeMetadata {
  TIMECODE_RATE?: string;
  Speed?: number;
  FILE_SAMPLE_RATE?: number;
  TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_LO?: number;
  TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_HI?: number;
}

interface iXML {
  Project?: string;
  Speed?: TimecodeMetadata;
  Scene?: string;
  Take?: string;
}


class IXML {
  Project?: string;
  Speed?: TimecodeMetadata;
  Scene?: string;
  Take?: string;

  constructor(ixml: any) {
    this.Project = ixml.PROJECT;
    this.Scene = ixml.SCENE;
    this.Take = ixml.TAKE;
    this.Speed = {
      TIMECODE_RATE: ixml.SPEED?.TIMECODE_RATE,
      FILE_SAMPLE_RATE: ixml.SPEED?.FILE_SAMPLE_RATE,
      TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_LO:
        ixml.SPEED?.TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_LO,
      TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_HI:
        ixml.SPEED?.TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_HI,
    };
  }

  convertSpeed(): Error | null {
    if (!this.Speed?.TIMECODE_RATE) {
      return new Error("no speed data present");
    }
    const [rate, div] = this.Speed.TIMECODE_RATE.split("/").map(Number);
    if (div === 0) {
      return new Error("couldn't find speed data");
    }
    this.Speed.Speed = rate / div;
    return null;
  }

  convertSamplesToTimecode(): string {
    if (
      !this.Speed?.TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_LO ||
      !this.Speed?.FILE_SAMPLE_RATE ||
      !this.Speed?.Speed
    ) {
      return "00:00:00:00";
    }

    const totalSeconds =
      this.Speed.TIMESTAMP_SAMPLES_SINCE_MIDNIGHT_LO /
      this.Speed.FILE_SAMPLE_RATE;
    const hours = Math.floor(totalSeconds / 3600);
    const minutes = Math.floor(totalSeconds / 60) % 60;
    const seconds = Math.floor(totalSeconds) % 60;
    const secondsFraction = totalSeconds - Math.floor(totalSeconds);
    const frames = Math.floor(secondsFraction * this.Speed.Speed);

    return `${String(hours).padStart(2, "0")}:${String(minutes).padStart(
      2,
      "0"
    )}:${String(seconds).padStart(2, "0")}:${String(frames).padStart(2, "0")}`;
  }
}

function parseIxml(xml: string): IXML | Error {
  const parser = new XMLParser();
  const parsed = parser.parse(xml);
  const ixml = new IXML(parsed.BWFXML);
  const err = ixml.convertSpeed();
  if (err) {
    return err;
  }
  return ixml;
}

try {
  main();
} catch (e) {
  console.error(e)
}
