<template>
    <el-footer>{{ branding.footerText }}</el-footer>
</template>

<script lang="ts" setup>
import { brandingState as branding } from '../branding'
</script>

<style>
.el-footer {
  background-color: #383737;
  color: #fff;
  text-align: center;
  line-height: 60px;
}
</style>
