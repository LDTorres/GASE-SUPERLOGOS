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
      <v-btn @click="logout()" flat class="mb-2">
        <v-icon>supervised_user_circle</v-icon>
        <h4 class="ml-2">SALIR</h4>
      </v-btn>
    </v-toolbar>
    <v-content v-bind:class="{ 'no-padding': !verifyUser }">
      <router-view/>
    </v-content>
    <section class="loader http">
      <div class="loaders">
        <div class="loader">
          <h3 class="display-1" style="margin-bottom: 20px">Cargango...</h3>
          <div class="loader-inner ball-pulse">
            <div></div>
            <div></div>
            <div></div>
          </div>
        </div>
      </div>
    </section>
  </v-app>
</template>

<script>
export default {
  created () {
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
      this.$router.push('/')
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
      if (this.$store.state.token) {
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

section.loader {
    position: absolute;
    height: 100vh;
    width: 100vw;
    display: none;
    z-index: 999;
}

.loaders {
    height: 100vh;
    width: 100vw;
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    display: -webkit-box;
    display: -ms-flexbox;
    display: flex;
    -webkit-box-flex: 0;
    -ms-flex: 0 1 auto;
    flex: 0 1 auto;
    -webkit-box-orient: horizontal;
    -webkit-box-direction: normal;
    -ms-flex-direction: row;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    background: #ffffffd4;
}
.loaders .loader {
  box-sizing: border-box;
  display: flex;
  flex: 0 1 auto;
  flex-direction: column;
  flex-grow: 1;
  flex-shrink: 0;
  flex-basis: 90%;
  max-width: 90%;
  height: 450px;
  align-items: center;
  justify-content: center;
}

@-webkit-keyframes scale {
  0% {
    -webkit-transform: scale(1);
    transform: scale(1);
    opacity: 1;
  }
  45% {
    -webkit-transform: scale(0.1);
    transform: scale(0.1);
    opacity: 0.7;
  }
  80% {
    -webkit-transform: scale(1);
    transform: scale(1);
    opacity: 1;
  }
}
@keyframes scale {
  0% {
    -webkit-transform: scale(1);
    transform: scale(1);
    opacity: 1;
  }
  45% {
    -webkit-transform: scale(0.1);
    transform: scale(0.1);
    opacity: 0.7;
  }
  80% {
    -webkit-transform: scale(1);
    transform: scale(1);
    opacity: 1;
  }
}

.ball-pulse > div:nth-child(1) {
  -webkit-animation: scale 0.75s -0.24s infinite cubic-bezier(0.2, 0.68, 0.18, 1.08);
  animation: scale 0.75s -0.24s infinite cubic-bezier(0.2, 0.68, 0.18, 1.08);
}

.ball-pulse > div:nth-child(2) {
  -webkit-animation: scale 0.75s -0.12s infinite cubic-bezier(0.2, 0.68, 0.18, 1.08);
  animation: scale 0.75s -0.12s infinite cubic-bezier(0.2, 0.68, 0.18, 1.08);
}

.ball-pulse > div:nth-child(3) {
  -webkit-animation: scale 0.75s 0s infinite cubic-bezier(0.2, 0.68, 0.18, 1.08);
  animation: scale 0.75s 0s infinite cubic-bezier(0.2, 0.68, 0.18, 1.08);
}

.ball-pulse > div {
  background-color: #10a6ef;
  width: 25px;
  height: 25px;
  border-radius: 100%;
  margin: 2px;
  -webkit-animation-fill-mode: both;
  animation-fill-mode: both;
  display: inline-block;
}
</style>


