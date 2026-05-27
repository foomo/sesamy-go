import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'sesamy-go',
  description: 'Server-side tag management system for Go — GA4 GTag + Measurement Protocol v2.',
  lang: 'en-US',
  cleanUrls: true,
  lastUpdated: true,
  appearance: 'dark',
  ignoreDeadLinks: false,
  base: '/sesamy-go/',
  sitemap: {
    hostname: 'https://foomo.github.io/sesamy-go',
  },
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    logo: '/logo.png',
    outline: [2, 4],
    nav: [
      { text: 'Guide', link: '/guide/what-is-sesamy', activeMatch: '/guide/' },
      { text: 'Reference', link: '/reference/packages', activeMatch: '/reference/' }
    ],
    sidebar: {
      '/guide/': [
        {
          text: 'Introduction',
          items: [
            { text: 'What is sesamy-go?', link: '/guide/what-is-sesamy' },
            { text: 'Getting started', link: '/guide/getting-started' },
            { text: 'Core concepts', link: '/guide/concepts' },
            { text: 'Pairing with sesamy-cli', link: '/guide/sesamy-cli' },
          ],
        },
        {
          text: 'Building a collect server',
          items: [
            { text: 'Server-side Collect', link: '/guide/server-collect' },
            { text: 'Events', link: '/guide/events' },
            { text: 'Sending events from Go', link: '/guide/client' },
          ],
        },
        {
          text: 'Integrations',
          items: [
            { text: 'Providers', link: '/guide/providers' },
            { text: 'Grafana Loki', link: '/guide/loki' },
          ],
        },
        {
          text: 'Contributing',
          collapsed: true,
          items: [
            { text: 'Guideline', link: '/CONTRIBUTING' },
            { text: 'Code of conduct', link: '/CODE_OF_CONDUCT' },
            { text: 'Security guidelines', link: '/SECURITY' },
          ],
        },
      ],
      '/reference/': [
        {
          text: 'Reference',
          items: [
            { text: 'Package overview', link: '/reference/packages' },
            { text: 'Encoding (GTag & MPv2)', link: '/reference/encoding' },
            { text: 'HTTP handlers & middleware', link: '/reference/http' },
            { text: 'Collect', link: '/reference/collect' },
            { text: 'Client', link: '/reference/client' },
          ],
        },
        {
          text: 'Providers',
          items: [
            { text: 'Cookiebot', link: '/reference/providers/cookiebot' },
            { text: 'Emarsys', link: '/reference/providers/emarsys' },
            { text: 'Tracify', link: '/reference/providers/tracify' },
          ],
        },
      ],
    },
    socialLinks: [
      { icon: 'github', link: 'https://github.com/foomo/sesamy-go' },
    ],
    editLink: {
      pattern: 'https://github.com/foomo/sesamy-go/edit/main/docs/:path',
    },
    search: {
      provider: 'local',
    },
    footer: {
      message: 'Made with ♥ <a href="https://www.foomo.org">foomo</a> by <a href="https://www.bestbytes.com">bestbytes</a>',
    },
  },
  markdown: {
    // https://github.com/vuejs/vitepress/discussions/3724
    theme: {
      light: 'catppuccin-latte',
      dark: 'catppuccin-frappe',
    },
  },
  head: [
    ['meta', { name: 'theme-color', content: '#ffffff' }],
    ['link', { rel: 'icon', href: '/logo.png' }],
    ['meta', { name: 'author', content: 'foomo by bestbytes' }],
    // OpenGraph
    ['meta', { property: 'og:title', content: 'foomo/sesamy-go' }],
    [
      'meta',
      {
        property: 'og:image',
        content: 'https://github.com/foomo/sesamy-go/blob/main/docs/public/banner.png?raw=true',
      },
    ],
    [
      'meta',
      {
        property: 'og:description',
        content: 'Server-side tag management SDK for Go — GA4 GTag + Measurement Protocol v2.',
      },
    ],
    ['meta', { name: 'twitter:card', content: 'summary_large_image' }],
    [
      'meta',
      {
        name: 'twitter:image',
        content: 'https://github.com/foomo/sesamy-go/blob/main/docs/public/banner.png?raw=true',
      },
    ],
    [
      'meta',
      { name: 'viewport', content: 'width=device-width, initial-scale=1.0, viewport-fit=cover' },
    ],
  ],
})
