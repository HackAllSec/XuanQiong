<template>
  <div>
    <el-container style="width: 100%;">
      <el-header style="background-color: #393939; display: flex; justify-content: flex-end; align-items: center; padding: 10px;">
          <div style="display: flex; justify-content: center; align-items: center; color: #fff;">
            <img src="../assets/vue.svg" fit="contain" style="margin-left: 1vw; height: 50px; width: 50px;">
            <span style="margin-left: 1vw;">玄穹后台管理系统</span>
          </div>
          <div style="cursor: pointer; margin-left: auto;">
            <el-dropdown style="color: #fff;" trigger="click" @command="handleCommand">
              <span style="display: flex; align-items: center;"><el-avatar style="margin-right: 10px;" :src="avatar" /> {{ username }}</span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item command=1>{{ t('app.webui.myprofile') }}</el-dropdown-item>
                  <el-dropdown-item command=2>{{ t('app.webui.alterpwd') }}</el-dropdown-item>
                  <el-dropdown-item command=3>{{ t('app.webui.logout') }}</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
      </el-header>
      <el-container style="width: 100%;">
        <el-menu
        :collapse="isCollapse"
        class="el-menu-vertical-demo"
        @select="handleMenuClick">
        <el-menu-item index="1" style="margin-top: 2vh;">
          <el-icon><Compass /></el-icon>
          <span>{{ t('app.webui.dashboard') }}</span>
        </el-menu-item>
        <el-menu-item index="2">
          <el-icon><setting /></el-icon>
          <span>{{ t('app.webui.systemsetting') }}</span>
        </el-menu-item>
        <el-menu-item index="3">
          <el-icon><User /></el-icon>
          <span>{{ t('app.webui.usermanager') }}</span>
        </el-menu-item>
        <el-sub-menu index="4">
          <template #title>
            <el-icon><svg t="1728359731855" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1480" width="200" height="200"><path d="M940 512H792V412c76.8 0 139-62.2 139-139 0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8 0 34.8-28.2 63-63 63H232c-34.8 0-63-28.2-63-63 0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8 0 76.8 62.2 139 139 139v100H84c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h148v96c0 6.5 0.2 13 0.7 19.3C164.1 728.6 116 796.7 116 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-44.2 23.9-82.9 59.6-103.7 6 17.2 13.6 33.6 22.7 49 24.3 41.5 59 76.2 100.5 100.5S460.5 960 512 960s99.8-13.9 141.3-38.2c41.5-24.3 76.2-59 100.5-100.5 9.1-15.5 16.7-31.9 22.7-49C812.1 793.1 836 831.8 836 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-79.3-48.1-147.4-116.7-176.7 0.4-6.4 0.7-12.8 0.7-19.3v-96h148c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zM716 680c0 36.8-9.7 72-27.8 102.9-17.7 30.3-43 55.6-73.3 73.3-20.1 11.8-42 20-64.9 24.3V484c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v396.5c-22.9-4.3-44.8-12.5-64.9-24.3-30.3-17.7-55.6-43-73.3-73.3C317.7 752 308 716.8 308 680V412h408v268z" p-id="1481" fill="#515151"></path><path d="M304 280h56c4.4 0 8-3.6 8-8 0-28.3 5.9-53.2 17.1-73.5 10.6-19.4 26-34.8 45.4-45.4C450.9 142 475.7 136 504 136h16c28.3 0 53.2 5.9 73.5 17.1 19.4 10.6 34.8 26 45.4 45.4C650 218.9 656 243.7 656 272c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-40-8.8-76.7-25.9-108.1-17.2-31.5-42.5-56.8-74-74C596.7 72.8 560 64 520 64h-16c-40 0-76.7 8.8-108.1 25.9-31.5 17.2-56.8 42.5-74 74C304.8 195.3 296 232 296 272c0 4.4 3.6 8 8 8z" p-id="1482" fill="#515151"></path></svg></el-icon>
            <span>{{ t('app.webui.vulnmanager') }}</span>
          </template>
          <el-menu-item index="4-1">
            <el-icon><Place /></el-icon>
            <span>{{ t('app.webui.vulntype') }}</span>
          </el-menu-item>
          <el-menu-item index="4-2">
            <el-icon><Place /></el-icon>
            <span>{{ t('app.webui.allvuln') }}</span>
          </el-menu-item>
          <el-menu-item index="4-3">
            <el-icon><Place /></el-icon>
            <span>{{ t('app.webui.unaudit') }}</span>
          </el-menu-item>
          <el-menu-item index="4-4">
            <el-icon><Place /></el-icon>
            <span>{{ t('app.webui.audited') }}</span>
          </el-menu-item>
        </el-sub-menu>
        <el-menu-item index="5">
          <el-icon><TrophyBase /></el-icon>
          <span>{{ t('app.webui.scorerule') }}</span>
        </el-menu-item>
        <!--
        <el-menu-item index="6">
          <el-icon><svg t="1728366011089" class="icon" viewBox="0 0 1161 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4589" width="200" height="200"><path d="M258.205538 354.461538c10.436923 0 18.924308 8.467692 18.924308 18.944v624.836924c0 10.456615-8.467692 18.944-18.924308 18.944H177.289846a18.944 18.944 0 0 1-18.924308-18.944v-624.836924c0-10.476308 8.467692-18.944 18.924308-18.944h80.915692z m158.365539 120.497231c10.436923 0 18.924308 8.467692 18.924308 18.924308v504.359385c0 10.456615-8.467692 18.944-18.924308 18.944h-80.915692a18.944 18.944 0 0 1-18.924308-18.944V493.883077c0-10.436923 8.467692-18.924308 18.924308-18.924308h80.915692z m-316.731077 0c10.456615 0 18.924308 8.467692 18.924308 18.924308v504.359385c0 10.456615-8.467692 18.944-18.924308 18.944H18.944A18.944 18.944 0 0 1 0 998.242462V493.883077c0-10.436923 8.467692-18.924308 18.944-18.924308H99.84z m139.421538-82.628923H196.214154v586.988308h43.047384V392.329846z m158.365539 120.477539h-43.027692V979.298462h43.027692V512.787692z m-316.731077 0H37.868308V979.298462h43.027692V512.787692zM1063.384615 0a98.461538 98.461538 0 0 1 98.461539 98.461538v748.307693a98.461538 98.461538 0 0 1-98.461539 98.461538H492.307692V433.230769a39.384615 39.384615 0 0 0-36.430769-39.286154L452.923077 393.846154h-98.461539v-59.076923a39.384615 39.384615 0 0 0-36.430769-39.286154L315.076923 295.384615h-98.461538V98.461538a98.461538 98.461538 0 0 1 98.461538-98.461538h748.307692z m-142.867692 520.664615c-44.898462 0-78.769231 14.572308-102.4 44.504616-21.267692 26.387692-31.507692 60.258462-31.507692 102.006154 0 42.535385 9.846154 76.406154 30.326154 101.612307 23.236923 29.144615 57.895385 44.110769 103.975384 44.11077 30.326154 0 56.32-8.664615 77.981539-25.993847 23.236923-18.510769 37.809231-44.110769 44.110769-76.8h-41.747692l-1.476923 5.395693c-5.710769 19.377231-15.064615 34.126769-28.061539 44.228923-13.390769 10.24-30.72 15.36-51.2 15.36-31.507692 0-54.744615-9.846154-69.710769-29.538462-14.178462-18.510769-21.267692-44.504615-21.267692-78.375384 0-32.689231 7.089231-58.683077 21.661538-77.587693 15.36-20.873846 38.203077-31.113846 68.529231-31.113846 20.48 0 37.021538 4.332308 50.412307 13.390769 13.390769 9.058462 22.449231 23.236923 27.175385 42.141539h41.747692l-0.984615-5.592616c-5.060923-25.796923-17.329231-46.710154-36.430769-62.148923-21.267692-17.329231-48.443077-25.6-81.132308-25.6z m-30.956308-413.538461c-42.929231 0-76.8 13.784615-101.218461 42.141538-23.630769 26.781538-35.052308 61.44-35.052308 104.369231 0 42.535385 11.421538 77.193846 35.052308 103.975385 24.418462 27.569231 58.289231 41.747692 101.218461 41.747692 42.535385 0 76.406154-13.784615 101.218462-41.353846 23.630769-26.781538 35.446154-61.44 35.446154-104.369231 0-43.323077-11.815385-78.375385-35.446154-104.763077-24.812308-27.963077-58.683077-41.747692-101.218462-41.747692z m-292.489846 5.513846H480.492308V393.846154h42.92923v-109.489231h72.861539l5.395692-0.059077c63.763692-1.516308 95.822769-30.247385 95.822769-86.193231 0-57.107692-33.476923-85.464615-100.430769-85.464615z m292.470154 32.295385c29.932308 0 53.169231 9.452308 69.316923 29.144615 15.753846 19.298462 24.024615 45.686154 24.024616 79.556923 0 33.476923-8.270769 59.864615-24.024616 78.769231-16.147692 19.298462-39.384615 29.144615-69.316923 29.144615s-53.169231-10.24-69.710769-29.932307c-15.753846-19.298462-23.630769-45.292308-23.630769-77.981539 0-33.083077 7.876923-59.076923 23.630769-78.375385 16.541538-20.48 39.778462-30.326154 69.710769-30.326153z m-296.014769 4.332307c20.873846 0 36.233846 3.938462 46.08 11.815385 9.846154 7.089231 14.966154 19.692308 14.966154 37.021538 0 17.329231-5.12 29.932308-14.572308 37.809231-9.846154 7.876923-25.206154 11.815385-46.473846 11.815385h-70.104616v-98.461539z" p-id="4590" fill="#515151"></path></svg></el-icon>
          <span>{{ t('app.webui.pocmanager') }}</span>
        </el-menu-item>
        <el-menu-item index="7">
          <el-icon><svg t="1727429727880" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="13778" width="200" height="200"><path d="M676.408889 303.407407l-101.262222-175.976296c-132.437333-40.201481-114.801778 130.085926-203.207111 168.846222L303.407407 246.708148C347.477333 160.274963 387.337481 83.057778 429.283556 0h206.582518L758.518519 235.747556 676.408889 303.407407zM127.393185 493.037037l80.327111 61.477926c-28.899556 61.060741-55.48563 117.380741-82.223407 173.018074 29.430519 138.657185 149.238519 47.824593 215.836444 103.917037V910.222222H94.966519L0 735.156148 127.393185 493.037037zM1024 735.004444L925.999407 910.222222H682.666667v-79.568592c68.342519-48.279704 171.956148 24.045037 212.574814-97.583408l-83.133629-184.395852L895.089778 493.037037 1024 735.004444zM0 387.640889L101.793185 189.62963 303.407407 422.874074l-12.818963 32.237037L0 387.640889zM1024 358.362074L739.745185 417.185185 720.592593 392.836741 919.210667 189.62963 1024 358.362074zM417.185185 1024l80.402963-263.661037 26.699852-1.820444L606.814815 1024h-189.62963z" fill="#515151" p-id="13779"></path><path d="M606.814815 455.111111h-189.62963v-75.851852h189.62963v75.851852zM682.666667 568.888889H341.333333v-75.851852h341.333334v75.851852zM796.444444 682.666667H227.555556v-75.851852h568.888888v75.851852z" fill="#515151" p-id="13780"></path></svg></el-icon>
          <span>{{ t('app.webui.expmanager') }}</span>
        </el-menu-item>
        <el-menu-item index="8">
          <el-icon><Histogram /></el-icon>
          <span>{{ t('app.webui.statistics') }}</span>
        </el-menu-item>
        -->
        <el-sub-menu index="6">
          <template #title>
            <svg viewBox="0 0 24 24" width="1.2em" height="1.2em" data-v-f414ea64="" class="el-icon">
              <path fill="currentColor" d="m18.5 10l4.4 11h-2.155l-1.201-3h-4.09l-1.199 3h-2.154L16.5 10h2zM10 2v2h6v2h-1.968a18.222 18.222 0 0 1-3.62 6.301a14.864 14.864 0 0 0 2.336 1.707l-.751 1.878A17.015 17.015 0 0 1 9 13.725a16.676 16.676 0 0 1-6.201 3.548l-.536-1.929a14.7 14.7 0 0 0 5.327-3.042A18.078 18.078 0 0 1 4.767 8h2.24A16.032 16.032 0 0 0 9 10.877a16.165 16.165 0 0 0 2.91-4.876L2 6V4h6V2h2zm7.5 10.885L16.253 16h2.492L17.5 12.885z"></path>
            </svg>
            <span>{{ t('app.webui.language') }}</span>
          </template>
          <el-menu-item index="11-1" @click="changelanguage('zh-CN')">简体中文</el-menu-item>
          <el-menu-item index="11-2" @click="changelanguage('en-US')">English</el-menu-item>
        </el-sub-menu>
        <div style="display: flex; margin: auto; cursor: pointer;">
          <el-icon v-if="!isCollapse" @click="handleCollapse" style="margin-left: auto;" size="25" color="#008B8B"><Fold /></el-icon>
          <el-icon v-else @click="handleCollapse" style="margin-left: auto;" size="25" color="#008B8B"><Expand /></el-icon>
        </div>
        </el-menu>
        <el-main>
          <div v-if="enumid === '1'">
            <Dashboard />
          </div>
          <div v-else-if="enumid === '2'">
            <System />
          </div>
          <div v-else-if="enumid === '3'">
            <Userlist />
          </div>
          <div v-else-if="enumid === '4-1'">
            <VulnType />
          </div>
          <div v-else-if="enumid === '4-2'">
            <Vulnlist />
          </div>
          <div v-else-if="enumid === '4-3'">
            <Unaudited />
          </div>
          <div v-else-if="enumid === '4-4'">
            <Audited />
          </div>
          <div v-else-if="enumid === '5'">
            <ScoreRule />
          </div>
          <div v-else-if="showprofile">
            <Profile />
          </div>
          <div v-else-if="changepwd">
            <Modifypasswd />
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import { User, Compass, TrophyBase, Histogram, Menu as IconMenu, Place, Setting, Fold, Expand } from '@element-plus/icons-vue'
import { useI18n } from 'vue-i18n';
import { onMounted } from 'vue';
import { useRoute } from 'vue-router'
import { ref } from 'vue'
import Dashboard from '../pages/Dashboard.vue';
import Userlist from '../pages/Userlist.vue';
import Vulnlist from '../pages/Vulnlist.vue';
import Unaudited from '../pages/Unaudited.vue';
import Audited from '../pages/Audited.vue';
import System from '../pages/System.vue';
import ScoreRule from '../pages/ScoreRule.vue';
import VulnType from '../pages/VulnType.vue';
import Profile from '../pages/Profile.vue';
import Modifypasswd from '../pages/Modifypasswd.vue';
import api from '../api'

const { t, locale } = useI18n();
const showprofile = ref(false)
const changepwd = ref(false)
const changelanguage = (language) => {
    locale.value = language;
    localStorage.setItem('selectedLanguage', language);
    enumid.value = '1'
};
const username = sessionStorage.getItem('username');
const avatar = sessionStorage.getItem('avatar');
const enumid = ref('1')
const isCollapse = ref(false)
const token = sessionStorage.getItem('token')

function handleCollapse() {
  isCollapse.value = !isCollapse.value
}
function handleMenuClick(index: string) {
      enumid.value = index
}

const router = useRoute()
onMounted(performAction);
function performAction() {
  if (!token) {
    router.push('/login')
  }
}

function handleCommand(command: number) {
    if (command == 1) {
      enumid.value = ''
      changepwd.value = false
      showprofile.value = true
    }
    if (command == 2) {
      enumid.value = ''
      showprofile.value = false
      changepwd.value = true
    }
    if (command == 3) {
      logout();
    }
  }
  async function logout() {
    try {
        const config = {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        };
        const response = await api.get('/api/v1/logout', config)
        sessionStorage.removeItem('username');
        sessionStorage.removeItem('token');
        sessionStorage.removeItem('avatar');
        location.reload();
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
  }
</script>

<style scoped>

.el-table {
  width: 80%;
  margin: auto;
  margin-top: 50px;
}
:global(h2#card-usage ~ .example .example-showcase) {
  background-color: var(--el-fill-color) !important;
}

.title {
  display: inline-flex;
  font-weight: bold;
  font-size: 16px;
  align-items: center
}
.statistic-card {
  height: 80%;
  padding: 20px;
  border-radius: 4px;
  background-color: var(--el-bg-color-overlay);
}

.statistic-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  font-size: 12px;
  color: var(--el-text-color-regular);
  margin-top: 16px;
}

.statistic-footer .footer-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.statistic-footer .footer-item span:last-child {
  display: inline-flex;
  align-items: center;
  margin-left: 4px;
}

.green {
  color: var(--el-color-success);
}
.red {
  color: var(--el-color-error);
}
.status {
  display: flex;
  gap: 4px;
}

.image-slot {
  display: flex;
  justify-content: center;
  width: 100px;
  height: 100px;
  background-color: #eee;
}
</style>