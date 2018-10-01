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
              <v-layout wrap>
                <v-flex xs12>
                  <v-text-field v-validate="'required|email'" prepend-icon="mail" name="Correo" label="Correo" type="text" v-model="login.email"></v-text-field>
                  <span v-show="errors.has('Correo')">{{ errors.first('Correo') }}</span>
                </v-flex>
                <v-flex xs12>
                  <v-text-field id="password" v-validate="'required|min:6|alpha_dash'" prepend-icon="lock" name="Contrasena" label="Contraseña" type="password" v-model="login.password"></v-text-field>
                  <span v-show="errors.has('Contrasena')">{{ errors.first('Contrasena') }}</span>
                </v-flex>
              </v-layout>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <!-- <v-btn color="primary" v-on:click="showRegister = !showRegister">Registro</v-btn> -->
            <v-spacer></v-spacer>
            <v-btn color="primary" v-on:click="singIn" :disabled="errors.count() > 0">Inicio de Sesion</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
      <!-- <v-flex xs12 sm8 md4 v-show="showRegister">
        <v-card class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>Registro</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form>
              <v-layout wrap>
                <v-flex xs12>
                  <v-text-field v-validate="'required'" prepend-icon="person-r" name="Nombre-r" label="Nombre de Usuario" type="text" v-model="register.username"></v-text-field>
                  <span v-show="errors.has('Nombre-r')">{{ errors.first('Nombre-r') }}</span>
                </v-flex>
                <v-flex xs12>
                  <v-text-field v-validate="'required|email'" prepend-icon="mail-r" name="Correo-r" label="Correo" type="text" v-model="register.email"></v-text-field>
                  <span v-show="errors.has('Correo-r')">{{ errors.first('Correo-r') }}</span>
                </v-flex>
                <v-flex xs12>              
                  <v-text-field v-validate="'required|min:8|alpha_dash'" id="password-r" prepend-icon="lock" name="Contrasena-r" label="Contraseña" type="password" v-model="register.password"></v-text-field>
                  <span v-show="errors.has('Contrasena-r')">{{ errors.first('Contrasena-r') }}</span>
                </v-flex>
              </v-layout>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" v-on:click="showRegister = !showRegister">Inicio de Sesion</v-btn>
            <v-spacer></v-spacer>
            <v-btn color="primary" v-on:click="singUp" :disabled="errors.count() > 0">Registro</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex> -->
    </v-layout>
  </v-container>
</template>

<script lang="js">
  export default {
    name: 'login',
    props: [],
    created () {
      if (this.verifyUser) {
        this.$router.push('/orders')
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
        this.$validator.validate().then(result => {
          if (!result) {
            alert('Llene los campos correctamente.')
          } else {
            let params = {
              item: this.login,
              message: 'Has iniciado sesion'
            }
            this.$store.dispatch('login', params)
          }
        })
      },
      singUp () {
        this.$validator.validate().then(result => {
          if (!result) {
            return alert('Llene los campos correctamente.')
          } else {
            let params = {
              item: this.register,
              message: 'Te has registrado con exito'
            }
            this.$store.dispatch('register', params)
          }
        })
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
