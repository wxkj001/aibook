<script setup lang="ts">
import { type MenuOption, NIcon,darkTheme } from 'naive-ui';
import { type Component, h, ref } from 'vue';
import { RouterLink } from 'vue-router'
import { SunnyOutline,MoonOutline,LogoGithub,ChatbubblesOutline as ChatbubblesIcon, ColorPaletteOutline as ColorPaletteIcon, CogOutline as CogIcon, PersonOutline as MyIcon } from '@vicons/ionicons5'
import { useThemeStore } from '@/stores/theme';
const renderIcon = (icon: Component) => {
    return () => h(NIcon, null, { default: () => h(icon) })
}
// const menuOptions: MenuOption[] = 
const menuOptions =ref<MenuOption[]>([
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    name: "homeChat"
                }
            },
            { default: () => '对话' }
        ),
        key: 'chat',
        icon: renderIcon(ChatbubblesIcon)
    },
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    name: "homeColor"
                }
            },
            { default: () => '绘画' }
        ),
        key: 'color',
        icon: renderIcon(ColorPaletteIcon)
    }
])
menuOptions.value.push(
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    name: "homeif",
                    params: {
                        href:"https://github.com",
                    }
                },
                
            },
            { default: () => '自定义' }
        ),
        key: 'zidingyi',
        icon: renderIcon(ColorPaletteIcon)
    }
)
const menuUserOptions = ref<MenuOption[]>([
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    name: "homeMy"
                }
            },
            { default: () => '我' }
        ),
        key: 'my',
        icon: renderIcon(MyIcon)
    },
    {
        label: () => h(
            RouterLink,
            {
                to: {
                    name: "homeSetting"
                }
            },
            { default: () => '设置' }
        ),
        key: 'setting',
        icon: renderIcon(CogIcon),

    },
])
const collapsed = ref(false)
const bottomValue = ref()
const mValue = ref()
const theme=useThemeStore()
const handleUpdateValue=(value: boolean)=>{
    theme.$patch({
        isDark: value
    })
}
const MeUpdateValue=(value: number)=>{
    bottomValue.value=0
    mValue.value=value
}

const MeUpdateValueBottom=(value: number)=>{
    bottomValue.value=value
    mValue.value=0
}
</script>
<template>
    <n-layout>
        <n-layout-header class="header">
            <n-grid :x-gap="5" :cols="24">
                <n-grid-item :offset="1">
                    LOGO
                </n-grid-item>
                <n-grid-item :span="1" :offset="20">
                    <n-switch @update:value="handleUpdateValue" :checked-value="true" :unchecked-value="false" size="large">
                        <template #checked-icon>
                            <n-icon :component="MoonOutline" />
                        </template>
                        <template #unchecked-icon>
                            <n-icon :component="SunnyOutline" />
                        </template>
                    </n-switch>
                </n-grid-item>
                <n-grid-item :span="1">
                    <n-icon :size="26" :component="LogoGithub" />
                </n-grid-item>
            </n-grid>
        </n-layout-header>
        <n-layout has-sider>
            <n-layout-sider bordered collapse-mode="width" :collapsed-width="64" :width="140" :collapsed="collapsed"
                show-trigger @collapse="collapsed = true" @expand="collapsed = false">
                <n-menu :value="mValue" @update:value="MeUpdateValue" :collapsed="collapsed" :collapsed-width="64" :collapsed-icon-size="22" :options="menuOptions" />
                <n-menu :value="bottomValue" @update:value="MeUpdateValueBottom" class="bottomMenu"  :collapsed="collapsed" :collapsed-width="64" :collapsed-icon-size="22"
                    :options="menuUserOptions" />
            </n-layout-sider>
            <n-layout-content content-style="padding: 15px;height: 95vh">
                <RouterView />
            </n-layout-content>
        </n-layout>

    </n-layout>
</template>

<style scoped>
.bottomMenu {
    position: absolute;
    bottom: 0;
    width: 100%;
}
.header{
    height: 5vh;
    padding: 5px;
}
</style>