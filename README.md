# farcaster-api-proxy

This is a [Next.js](https://nextjs.org/) project using the App Router that proxies Farcaster API endpoints.

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

## API Routes

All routes proxy to corresponding `client.farcaster.xyz/v2/` endpoints:

- `/[username]` → `user-by-username?username=${username}`
- `/[username]/[hash]` → `user-thread-casts?castHashPrefix=${hash}&username=${username}&limit=5`
- `/fids/[fid]` → `user?fid=${fid}`
- `/search/casts` → `search-casts?q=${q}&limit=${limit}`
- `/search/users` → `search-users?q=${q}&excludeSelf=${excludeSelf}&limit=${limit}&includeDirectCastAbility=${includeDirectCastAbility}`
- `/search/summary` → `search-summary?q=${q}&maxChannels=${maxChannels}&maxUsers=${maxUsers}&maxMiniApps=${maxMiniApps}&maxTokens=${maxTokens}&addFollowersYouKnowContext=${addFollowersYouKnowContext}`

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.
