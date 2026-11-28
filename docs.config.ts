const repositoryUrl = 'https://github.com/puremixlab/purekit';
const siteUrl = process.env.PUREKIT_DOCS_URL || 'http://localhost:5173';

export default {
  atomDirs: [{ dir: 'catalog' }],
  homePage: './docs/home/home.tsx',
  publicDocs: ['docs/getting-started.mdx', 'docs/native-desktop.mdx'],
  description: 'UI components for native Go apps, drawn on the GPU with MyGo.',
  favicons: {
    icon: '/favicon-32x32.png',
  },
  siteUrl,
  themeConfig: {
    apiHeader: {
      github: repositoryUrl,
    },
    metadata: {
      openGraph: {
        image: new URL('/og.png', siteUrl).href,
      },
    },
    navItems: [
      { href: '/getting-started', icon: 'rocket', label: 'Getting Started' },
      { href: '/native-desktop', icon: 'monitor', label: 'Native desktop' },
    ],
    prefersColor: 'auto',
    socialLinks: [{ href: repositoryUrl, icon: 'github', label: 'GitHub' }],
  },
  title: 'PureKit',
};
