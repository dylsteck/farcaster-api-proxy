import Cors from "cors";
import { NextRequest, NextResponse } from "next/server";

// Initializing the cors middleware
// You can read more about the available options here: https://github.com/expressjs/cors#configuration-options
const cors = Cors({
  methods: ["POST", "GET", "HEAD"],
});

// Helper method to wait for a middleware to execute before continuing
// And to throw an error when an error happens in a middleware
function runMiddleware(req: any, res: any, fn: Function) {
  return new Promise((resolve, reject) => {
    fn(req, res, (result: any) => {
      if (result instanceof Error) {
        return reject(result);
      }

      return resolve(result);
    });
  });
}

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ fid: string }> }
) {
  const { fid } = await params;

  if (!fid) {
    return NextResponse.json({ error: "Missing fid" }, { status: 400 });
  }

  const farcaster = await fetch(
    `https://client.farcaster.xyz/v2/user?fid=${fid}`
  );
  const cast = await farcaster.json();

  return NextResponse.json(cast, {
    headers: {
      "Cache-Control": "s-maxage=3600",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, POST, HEAD",
      "Access-Control-Allow-Headers": "Content-Type",
    },
  });
}

export async function POST(
  request: NextRequest,
  { params }: { params: Promise<{ fid: string }> }
) {
  const { fid } = await params;

  if (!fid) {
    return NextResponse.json({ error: "Missing fid" }, { status: 400 });
  }

  const farcaster = await fetch(
    `https://client.farcaster.xyz/v2/user?fid=${fid}`
  );
  const cast = await farcaster.json();

  return NextResponse.json(cast, {
    headers: {
      "Cache-Control": "s-maxage=3600",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, POST, HEAD",
      "Access-Control-Allow-Headers": "Content-Type",
    },
  });
}

export async function HEAD(
  request: NextRequest,
  { params }: { params: Promise<{ fid: string }> }
) {
  const { fid } = await params;

  if (!fid) {
    return new NextResponse(null, { status: 400 });
  }

  return new NextResponse(null, {
    status: 200,
    headers: {
      "Cache-Control": "s-maxage=3600",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, POST, HEAD",
      "Access-Control-Allow-Headers": "Content-Type",
    },
  });
}
