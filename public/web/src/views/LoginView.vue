<script setup lang="ts">
import router from '../router';
import { useUserStore } from '../stores/user';

const user = useUserStore()
const login=()=>{
    console.log(1);
    user.$patch({
        user:{
            name:"wxkj",
            token:"123456"
        }
    })
    router.push({name:"home"})
}
</script>
<template>
    <n-grid :x-gap="12" :y-gap="12" :cols="4" layout-shift-disabled>
        <n-gi></n-gi>
        <n-gi :span="2">
            <n-card>
                <n-tabs class="card-tabs" default-value="signin" size="large" animated pane-wrapper-style="margin: 0 -4px"
                    pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;">
                    <n-tab-pane name="signin" tab="登录">
                        <n-form>
                            <n-form-item-row label="用户名">
                                <n-input />
                            </n-form-item-row>
                            <n-form-item-row label="密码">
                                <n-input />
                            </n-form-item-row>
                        </n-form>
                        <n-button type="primary" block secondary strong @click="login">
                            登录
                        </n-button>
                    </n-tab-pane>
                    <n-tab-pane name="signup" tab="注册">
                        <n-form>
                            <n-form-item-row label="用户名">
                                <n-input />
                            </n-form-item-row>
                            <n-form-item-row label="密码">
                                <n-input />
                            </n-form-item-row>
                            <n-form-item-row label="重复密码">
                                <n-input />
                            </n-form-item-row>
                        </n-form>
                        <n-button type="primary" block secondary strong>
                            注册
                        </n-button>
                    </n-tab-pane>
                </n-tabs>
            </n-card>
        </n-gi>
        <n-gi></n-gi>
    </n-grid>
</template>