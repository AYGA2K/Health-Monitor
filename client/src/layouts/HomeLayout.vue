<template>
  <q-layout view=" hHh lPr fFf">

    <q-header class="">
      <q-toolbar :class="{ 'text-white': $q.dark.isActive, 'text-dark': !$q.dark.isActive }">
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

        <q-toolbar-title>
          The App
        </q-toolbar-title>

      </q-toolbar>

    </q-header>
    <q-drawer class="bg-transparent" v-model="leftDrawerOpen" show-if-above>
      <q-list>
        <q-item-label header>
          Explore
        </q-item-label>

        <EssentialLink v-for="link in essentialLinks" :key="link.title" v-bind="link" />
        <q-item clickable @click="toggleDarkMode()" class="fixed-bottom">
          <q-item-section avatar>
            <q-icon name="dark_mode" />
          </q-item-section>

          <q-item-section>
            <q-item-label>Toggle</q-item-label>
            <q-item-label caption></q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>

    <q-page-container class="page-container">
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import EssentialLink from 'components/EssentialLink.vue';
import { useQuasar } from 'quasar'
const linksList = [
  {
    title: 'Home',
    icon: 'person',
    link: '/'
  },
  {
    title: 'Something',
    icon: 'person_add',
    link: ''
  },
];

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink
  },

  setup() {
    const $q = useQuasar()
    const leftDrawerOpen = ref(false)
    function toggleDarkMode() {


      // get status
      console.log($q.dark.isActive) // true, false

      // get configured status
      console.log($q.dark.mode) // "auto", true, false


      // toggle
      $q.dark.toggle()
    }
    return {
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleDarkMode,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value
      }
    }
  }
});
</script>
<style  lang="scss"></style>
