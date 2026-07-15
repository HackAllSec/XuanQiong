import { reactive } from 'vue'
import api from './api'

export const brandingState = reactive({
  loaded: false,
  siteName: '玄穹漏洞库平台',
  frontendTitle: '玄穹漏洞库平台',
  adminTitle: '玄穹后台管理系统',
  logoUrl: '/avatar.svg',
  faviconUrl: '',
  footerText: 'Copyright © 2024. Hack All Sec rights reserved.',
  helpUrl: 'https://github.com/HackAllSec/XuanQiong',
  suggestUrl: 'https://github.com/HackAllSec/XuanQiong/issues',
})

function updateFavicon(url: string) {
  if (!url) {
    return
  }
  let link = document.querySelector("link[rel*='icon']") as HTMLLinkElement | null
  if (!link) {
    link = document.createElement('link')
    link.rel = 'icon'
    document.head.appendChild(link)
  }
  link.href = url
}

function applyFrontendBranding() {
  document.title = brandingState.frontendTitle || brandingState.siteName
  updateFavicon(brandingState.faviconUrl)
}

export async function loadBranding(force = false) {
  if (brandingState.loaded && !force) {
    applyFrontendBranding()
    return brandingState
  }
  try {
    const response = await api.get('/api/v1/getbrandconfig')
    if (response.data.code === 1) {
      const data = response.data.data
      Object.assign(brandingState, {
        siteName: data.site_name || brandingState.siteName,
        frontendTitle: data.frontend_title || brandingState.frontendTitle,
        adminTitle: data.admin_title || brandingState.adminTitle,
        logoUrl: data.logo_url || brandingState.logoUrl,
        faviconUrl: data.favicon_url || brandingState.faviconUrl,
        footerText: data.footer_text || brandingState.footerText,
        helpUrl: data.help_url || brandingState.helpUrl,
        suggestUrl: data.suggest_url || brandingState.suggestUrl,
        loaded: true,
      })
    }
  } catch (error) {
    console.error(error)
  }
  applyFrontendBranding()
  return brandingState
}
