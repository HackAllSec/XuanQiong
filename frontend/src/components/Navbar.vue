<template>
    <el-menu class="site-navbar" mode="horizontal" :ellipsis="false">
      <el-menu-item index="0" class="brand-menu-item" @click="GotoIndex">
        <img class="brand-logo" :src="branding.logoUrl" alt="logo" />
        <span class="brand-name">{{ branding.siteName }}</span>
      </el-menu-item>
      <el-menu-item index="1" @click="GotoIndex">{{ t('app.webui.home') }}</el-menu-item>
      <el-menu-item index="2" @click="SubmitVuln">{{ t('app.webui.submitvuln') }}</el-menu-item>
      <el-menu-item index="3" @click="Ranklist">{{ t('app.webui.rankinglist') }}</el-menu-item>
      <el-sub-menu index="4">
        <template #title>{{ t('app.webui.about') }}</template>
        <el-menu-item index="4-1" @click="Help">{{ t('app.webui.help') }}</el-menu-item>
        <el-menu-item index="4-2" @click="Suggest">{{ t('app.webui.suggest') }}</el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="5">
        <template #title>
          <svg viewBox="0 0 24 24" width="1.2em" height="1.2em" data-v-f414ea64="" class="el-icon">
            <path fill="currentColor" d="m18.5 10l4.4 11h-2.155l-1.201-3h-4.09l-1.199 3h-2.154L16.5 10h2zM10 2v2h6v2h-1.968a18.222 18.222 0 0 1-3.62 6.301a14.864 14.864 0 0 0 2.336 1.707l-.751 1.878A17.015 17.015 0 0 1 9 13.725a16.676 16.676 0 0 1-6.201 3.548l-.536-1.929a14.7 14.7 0 0 0 5.327-3.042A18.078 18.078 0 0 1 4.767 8h2.24A16.032 16.032 0 0 0 9 10.877a16.165 16.165 0 0 0 2.91-4.876L2 6V4h6V2h2zm7.5 10.885L16.253 16h2.492L17.5 12.885z"></path>
          </svg>
          <span>{{ t('app.webui.language') }}</span>
        </template>
        <el-menu-item index="5-1" @click="changelanguage('zh-CN')">简体中文</el-menu-item>
        <el-menu-item index="5-2" @click="changelanguage('en-US')">English</el-menu-item>
      </el-sub-menu>
      <el-menu-item index="6" @click="Searchvuln">{{ t('app.webui.search') }}</el-menu-item>
      <div class="navbar-account">
        <div v-if="!username" @click="Login" style="display: flex; align-items: center;">
          <el-icon style="margin-right: 10px;"><UserFilled /></el-icon>
          <span>{{ t('app.webui.login') }}</span>
        </div>
        <div v-else>
          <el-dropdown style="color: #fff;" trigger="click" @command="handleCommand">
            <span style="display: flex; align-items: center;"><el-avatar style="margin-right: 10px;" :src="avatar"></el-avatar>{{ username }}</span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command=1>{{ t('app.webui.myprofile') }}</el-dropdown-item>
              <el-dropdown-item command=2>{{ t('app.webui.mysubmitvuln')}}</el-dropdown-item>
              <el-dropdown-item command=3>{{ t('app.webui.messages') }}</el-dropdown-item>
              <el-dropdown-item command=4>{{ t('app.webui.apikeymanager') }}</el-dropdown-item>
              <el-dropdown-item command=5>{{ t('app.webui.modifypassword') }}</el-dropdown-item>
              <el-dropdown-item command=6>{{ t('app.webui.logout') }}</el-dropdown-item>
            </el-dropdown-menu>
          </template>
          </el-dropdown>
        </div>
      </div>
    </el-menu>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted  } from 'vue'
  import { UserFilled } from '@element-plus/icons-vue'
  import { useRouter } from 'vue-router'
  import { useI18n } from 'vue-i18n';
  import api from '../api'
  import { clearAuthSession } from '../auth'
  import { brandingState as branding, loadBranding } from '../branding'

  const router = useRouter()
  const { t, locale } = useI18n();
  const username = sessionStorage.getItem('username')
  const avatar = sessionStorage.getItem('avatar')

  onMounted(() => {
    loadBranding()
  })

  const changelanguage = (language) => {
    locale.value = language;
    localStorage.setItem('selectedLanguage', language);
  };
  function handleCommand(command: number) {
    if (command == 1) {
      router.push('/profile');
    }
    if (command == 2) {
      router.push('/myvulns');
    }
    if (command == 3) {
      router.push('/messages');
    }
    if (command == 4) {
      router.push('/apikeys');
    }
    if (command == 5) {
      router.push('/modifypwd');
    }
    if (command == 6) {
      logout();
    }
  }

  function GotoIndex() {
    router.push('/');
  }
  function SubmitVuln() {
    router.push('/submit');
  }
  function Ranklist() {
    router.push('/ranklist');
  }
  function Help() {
    window.open(branding.helpUrl);
  }
  function Suggest() {
    window.open(branding.suggestUrl);
  }
  function Searchvuln() {
    router.push('/search');
  }
  function Login() {
    router.push('/login');
  }
  async function logout() {
    try {
      const token = sessionStorage.getItem('token')
        const response = await api.get('/api/v1/logout')
        clearAuthSession()
        location.reload();
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
  }
  </script>
  
  <style scoped>
  .site-navbar {
    align-items: center;
    width: 100%;
    --el-menu-bg-color: #383737;
    --el-menu-text-color: #fff;
    padding-inline: 1%;
}
  .brand-menu-item {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-right: 20px;
    padding-inline: 12px;
  }
  .brand-logo {
    width: 44px;
    height: 44px;
    object-fit: contain;
  }
  .brand-name {
    color: #fff;
    font-weight: bold;
    font-size: 18px;
    white-space: nowrap;
  }
  .navbar-account {
    margin-left: auto;
    cursor: pointer;
    color: var(--el-menu-text-color);
  }
  </style>
  
