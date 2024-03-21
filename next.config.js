/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: "/:path*",
        destination: "https://client.warpcast.com/:path*",
      },
    ];
  },
};

module.exports = nextConfig;
