<template>
  <el-container class="layout">
    <el-header class="layout-header">
      <div class="brand-block">
        <img :src="branding.logoUrl" class="brand-logo" alt="logo">
        <span class="brand-title">{{ branding.adminTitle }}</span>
      </div>
      <el-dropdown class="user-dropdown" trigger="click" @command="handleCommand">
        <span class="user-trigger">
          <el-avatar :src="avatar" />
          {{ username }}
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">{{ t('app.webui.myprofile') }}</el-dropdown-item>
            <el-dropdown-item command="password">{{ t('app.webui.alterpwd') }}</el-dropdown-item>
            <el-dropdown-item command="logout">{{ t('app.webui.logout') }}</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </el-header>
    <el-container>
      <el-aside :width="isCollapse ? '64px' : '240px'" class="layout-aside">
        <el-menu :default-active="activeKey" :collapse="isCollapse" class="menu" @select="handleMenuClick">
          <el-menu-item v-for="item in topLevelMenus" :key="item.key" :index="item.key">
            <el-icon><component :is="item.icon" /></el-icon>
            <span>{{ item.label }}</span>
          </el-menu-item>

          <el-sub-menu v-if="vulnerabilityMenus.length > 0" index="vulnerability">
            <template #title>
              <el-icon><WarningFilled /></el-icon>
              <span>{{ t('app.webui.vulnmanager') }}</span>
            </template>
            <el-menu-item v-for="item in vulnerabilityMenus" :key="item.key" :index="item.key">
              <el-icon><component :is="item.icon" /></el-icon>
              <span>{{ item.label }}</span>
            </el-menu-item>
          </el-sub-menu>

          <el-sub-menu index="language">
            <template #title>
              <el-icon><Operation /></el-icon>
              <span>{{ t('app.webui.language') }}</span>
            </template>
            <el-menu-item index="language-zh" @click="changeLanguage('zh-CN')">简体中文</el-menu-item>
            <el-menu-item index="language-en" @click="changeLanguage('en-US')">English</el-menu-item>
          </el-sub-menu>
        </el-menu>

        <div class="collapse-trigger" @click="toggleCollapse">
          <el-icon v-if="isCollapse"><Expand /></el-icon>
          <el-icon v-else><Fold /></el-icon>
        </div>
      </el-aside>

      <el-main class="layout-main">
        <Profile v-if="currentView === 'profile'" />
        <Modifypasswd v-else-if="currentView === 'password'" />
        <component :is="currentComponent" v-else-if="currentComponent" />
        <el-empty v-else :description="t('app.webui.nopermation')" />
      </el-main>
    </el-container>
  </el-container>
</template>

<script lang="ts" setup>
import {
  Compass,
  Setting,
  User,
  Fold,
  Expand,
  Operation,
  WarningFilled,
  List,
  DocumentChecked,
  Finished,
  Files,
  DataBoard,
  Tickets,
} from '@element-plus/icons-vue'
import { computed, markRaw, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import api from '../api'
import { brandingState as branding, loadBranding } from '../branding'
import { clearAuthSession, getStoredPermissions, hasPermission } from '../auth'
import Dashboard from '../pages/Dashboard.vue'
import Userlist from '../pages/Userlist.vue'
import RoleManager from '../pages/RoleManager.vue'
import Vulnlist from '../pages/Vulnlist.vue'
import Unaudited from '../pages/Unaudited.vue'
import Audited from '../pages/Audited.vue'
import System from '../pages/System.vue'
import ScoreRule from '../pages/ScoreRule.vue'
import VulnType from '../pages/VulnType.vue'
import Profile from '../pages/Profile.vue'
import Modifypasswd from '../pages/Modifypasswd.vue'
import AuditLogs from '../pages/AuditLogs.vue'

const { t, locale } = useI18n()
const router = useRouter()
const isCollapse = ref(false)
const activeKey = ref('')
const currentView = ref<'menu' | 'profile' | 'password'>('menu')

const username = computed(() => sessionStorage.getItem('username') || 'admin')
const avatar = computed(() => sessionStorage.getItem('avatar') || '/avatar.svg')
const forcePasswordChange = computed(() => sessionStorage.getItem('force_password_change') === '1')

const menuComponents: Record<string, any> = {
  dashboard: markRaw(Dashboard),
  system: markRaw(System),
  users: markRaw(Userlist),
  roles: markRaw(RoleManager),
  vulnType: markRaw(VulnType),
  allVuln: markRaw(Vulnlist),
  unaudited: markRaw(Unaudited),
  audited: markRaw(Audited),
  scoreRule: markRaw(ScoreRule),
  auditLogs: markRaw(AuditLogs),
}

const topLevelMenus = computed(() => {
  const items = []
  if (hasPermission('dashboard.read')) {
    items.push({ key: 'dashboard', label: t('app.webui.dashboard'), icon: markRaw(Compass) })
  }
  if (hasPermission('system.config.read')) {
    items.push({ key: 'system', label: t('app.webui.systemsetting'), icon: markRaw(Setting) })
  }
  if (hasPermission('user.read')) {
    items.push({ key: 'users', label: t('app.webui.usermanager'), icon: markRaw(User) })
  }
  if (hasPermission('role.read')) {
    items.push({ key: 'roles', label: t('app.webui.rolemanager'), icon: markRaw(Tickets) })
  }
  if (hasPermission('score.rule.read')) {
    items.push({ key: 'scoreRule', label: t('app.webui.scorerule'), icon: markRaw(DataBoard) })
  }
  if (hasPermission('audit.log.read')) {
    items.push({ key: 'auditLogs', label: t('app.webui.auditlogs'), icon: markRaw(Files) })
  }
  return items
})

const vulnerabilityMenus = computed(() => {
  const items = []
  if (hasPermission('vuln.type.read')) {
    items.push({ key: 'vulnType', label: t('app.webui.vulntype'), icon: markRaw(List) })
  }
  if (hasPermission('vuln.read')) {
    items.push({ key: 'allVuln', label: t('app.webui.allvuln'), icon: markRaw(Files) })
  }
  if (hasPermission('vuln.audit.read')) {
    items.push({ key: 'unaudited', label: t('app.webui.unaudit'), icon: markRaw(DocumentChecked) })
    items.push({ key: 'audited', label: t('app.webui.audited'), icon: markRaw(Finished) })
  }
  return items
})

const currentComponent = computed(() => menuComponents[activeKey.value] || null)

function firstAvailableMenu(): string {
  const allMenus = [...topLevelMenus.value, ...vulnerabilityMenus.value]
  return allMenus.length > 0 ? allMenus[0].key : ''
}

function toggleCollapse() {
  isCollapse.value = !isCollapse.value
}

function ensureMenuSelection() {
  if (forcePasswordChange.value) {
    currentView.value = 'password'
    return
  }
  if (!activeKey.value || !menuComponents[activeKey.value]) {
    activeKey.value = firstAvailableMenu()
  }
}

function handleMenuClick(index: string) {
  if (index.startsWith('language-')) {
    return
  }
  if (forcePasswordChange.value) {
    currentView.value = 'password'
    return
  }
  currentView.value = 'menu'
  activeKey.value = index
}

function changeLanguage(language: string) {
  locale.value = language
  localStorage.setItem('selectedLanguage', language)
}

function handleCommand(command: string) {
  if (command === 'logout') {
    logout()
    return
  }
  if (forcePasswordChange.value && command !== 'password') {
    currentView.value = 'password'
    return
  }
  currentView.value = command === 'profile' ? 'profile' : 'password'
}

async function logout() {
  try {
    await api.get('/api/v1/logout')
  } finally {
    clearAuthSession()
    router.push('/login')
  }
}

onMounted(async () => {
  if (!sessionStorage.getItem('token') || !getStoredPermissions().includes('admin.panel.access')) {
    clearAuthSession()
    router.push('/login')
    return
  }
  await loadBranding()
  ensureMenuSelection()
})
</script>

<style scoped>
.layout {
  min-height: 100vh;
}

.layout-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  background: #2f3440;
  color: #fff;
}

.brand-block {
  display: flex;
  align-items: center;
  gap: 12px;
}

.brand-logo {
  width: 36px;
  height: 36px;
  object-fit: contain;
  border-radius: 6px;
  background: #fff;
}

.brand-title {
  font-size: 18px;
  font-weight: 600;
}

.user-dropdown {
  color: #fff;
}

.user-trigger {
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
}

.layout-aside {
  border-right: 1px solid var(--el-border-color-light);
  background: #fff;
}

.menu {
  min-height: calc(100vh - 60px);
  border-right: none;
}

.collapse-trigger {
  display: flex;
  justify-content: center;
  padding: 12px 0;
  cursor: pointer;
  color: var(--el-color-primary);
}

.layout-main {
  min-height: calc(100vh - 60px);
  background: var(--el-fill-color-light);
}
</style>
