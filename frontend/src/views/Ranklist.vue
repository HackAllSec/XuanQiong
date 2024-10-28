<template>
    <el-empty v-if = "data == ''" :description="t('app.webui.empty')" />
    <div v-else style="width: 70%; height: 90vh; margin: auto; display: flex; gap: 5%;">
        <el-card style="width: 50%; font-weight: bold; font-size: 20px;" shadow="always" :header="t('app.webui.annualranking')">
            <div v-for="(item, index) in data.annual" :key="index" style="font-weight: lighter; font-size: 14px; display: flex; align-items: center; padding: 8px;">
                <el-avatar :src="item.avatar" size="large"></el-avatar>
                <span style="margin-left: 10px;">{{ item.username }}</span>
                <span style="margin-left: auto;">{{ item.ranking }} {{ t('app.webui.score') }}</span>
            </div>
        </el-card>
        <el-card style="width: 50%; font-weight: bold; font-size: 20px;" shadow="always" :header="t('app.webui.quarterlyranking')">
            <div v-for="(item, index) in data.quarterly" :key="index" style="font-weight: lighter; font-size: 14px; display: flex; align-items: center; padding: 8px;">
                <el-avatar :src="item.avatar" size="large"></el-avatar>
                <span style="margin-left: 10px;">{{ item.username }}</span>
                <span style="margin-left: auto;">{{ item.ranking }} {{ t('app.webui.score') }}</span>
            </div>
        </el-card>
        <el-card style="width: 50%; font-weight: bold; font-size: 20px;" shadow="always" :header="t('app.webui.monthlyranking')">
            <div v-for="(item, index) in data.monthly" :key="index" style="font-weight: lighter; font-size: 14px; display: flex; align-items: center; padding: 8px;">
                <el-avatar :src="item.avatar" size="large"></el-avatar>
                <span style="margin-left: 10px;">{{ item.username }}</span>
                <span style="margin-left: auto;">{{ item.ranking }} {{ t('app.webui.score') }}</span>
            </div>
        </el-card>
    </div>
</template>
<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n';
import api from '../api'
const { t } = useI18n()
const data = ref('')
onMounted(getRankingList)
async function getRankingList() {
    const res = await api.get('/api/v1/usertop')
    const rankings = res.data;

        // 遍历每个排名列表并处理空的 avatar
        Object.keys(rankings).forEach((period) => {
            rankings[period].forEach((user) => {
                user.avatar = '/download/file?id=' + user.avatar || '/avatar.svg';
            });
        });
    data.value = rankings
}
console.log(data.value)
</script>
<style scoped>
    
</style>