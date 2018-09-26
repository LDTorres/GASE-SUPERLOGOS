<template>
  <v-app>
    <v-navigation-drawer
      persistent
      :mini-variant="miniVariant"
      :clipped="clipped"
      v-model="drawer"
      enable-resize-watcher
      fixed
      app
    >
      <v-list>
        <v-list v-for="(items, i) in sidemenu"
            :key="i">
          <v-subheader>{{items.title}}</v-subheader>
          
          <v-list-tile v-if="items.routes" :to="item.route"
            value="true"
            v-for="(item, i) in items.routes"
            :key="i"
          >
            <v-list-tile-action>
              <v-icon v-html="item.icon"></v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title class="dark--text" v-text="item.title"></v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
        <v-divider></v-divider>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar v-show="verifyUser"
      app
      :clipped-left="clipped"
    >
      <v-btn icon @click.stop="miniVariant = !miniVariant">
        <v-icon v-html="miniVariant ? 'chevron_right' : 'chevron_left'"></v-icon>
      </v-btn>
      <v-toolbar-title v-text="title">
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn @click="logout()" flat class="mb-2">SALIR</v-btn>
    </v-toolbar>
    <v-content v-bind:class="{ 'no-padding': !verifyUser }">
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>
export default {
  beforeCreated () {
    if (!this.verifyUser) {
      this.$router.push('/')
    }
  },
  data () {
    return {
      clipped: false,
      fixed: false,
      miniVariant: false,
      right: true,
      rightDrawer: true,
      title: 'Liderlogo Admin'
    }
  },
  methods: {
    logout () {
      localStorage.clear()
      location.reload()
    }
  },
  computed: {
    sidemenu () {
      return this.$store.state.app.sidemenu
    },
    drawer: {
      set (value) {
        return value
      },
      get () {
        return this.verifyUser
      }
    },
    verifyUser () {
      if (
        this.$store.state.token !== null &&
        this.$store.state.token !== undefined &&
        this.$store.state.token !== ''
      ) {
        return true
      }
      return false
    }
  },
  name: 'App'
}
</script>

<style>
main.no-padding.v-content {
  padding: 0 !important;
}
</style>


