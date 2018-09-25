<template lang="html">
  <v-container fluid fill-height class="login">
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4 v-show="!showRegister">
        <v-card class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>Inicio de Sesion</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form>
              <v-text-field prepend-icon="mail" name="email" label="Email" type="text" v-model="login.email"></v-text-field>
              <v-text-field id="password" prepend-icon="lock" name="password" label="Password" type="password" v-model="login.password"></v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" v-on:click="showRegister = !showRegister">Registro</v-btn>
            <v-spacer></v-spacer>
            <v-btn color="primary" v-on:click="singIn">Login</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
      <v-flex xs12 sm8 md4 v-show="showRegister">
        <v-card class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>Registro</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form>
              <v-text-field prepend-icon="person-r" name="username-r" label="Nombre" type="text" v-model="register.username"></v-text-field>
              <v-text-field prepend-icon="mail-r" name="email-r" label="Email" type="text" v-model="register.email"></v-text-field>              <v-text-field id="password-r" prepend-icon="lock" name="password-r" label="Password" type="password" v-model="register.password"></v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" v-on:click="showRegister = !showRegister">Inicio de Sesion</v-btn>
            <v-spacer></v-spacer>
            <v-btn color="primary" v-on:click="singUp">Regitro</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script lang="js">
  export default {
    name: 'login',
    props: [],
    created () {
      if (this.verifyUser) {
        this.$router.push('/home')
      }
    },
    mounted () {

    },
    data () {
      return {
        login: {
          email: '',
          password: ''
        },
        register: {
          username: '',
          email: '',
          password: ''
        },
        showRegister: false
      }
    },
    methods: {
      singIn () {
        if (this.login.email !== '' && this.login.password !== '') {
          let params = {
            item: this.login,
            message: 'Has iniciado sesion'
          }
          this.$store.dispatch('login', params)
        }
      },
      singUp () {
        if (this.register.username !== '' && this.register.password !== '' && this.register.email !== '') {
          let params = {
            item: this.register,
            message: 'Te has registrado con exito'
          }
          this.$store.dispatch('register', params)
        }
      }
    },
    computed: {
      verifyUser () {
        if (this.$store.state.token !== null && this.$store.state.token !== undefined && this.$store.state.token !== '') {
          return true
        }
        return false
      }
    }
}
</script>

<style scoped lang="css">
[text-center] {
  text-align: center;
}
</style>
