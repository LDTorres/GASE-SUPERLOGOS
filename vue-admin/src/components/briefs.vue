<template class="briefs">
<div>
  <v-toolbar flat color="white" class="sticky-nav">
    <v-toolbar-title hidden-md-and-down class="text-capitalize">{{ viewNameESP }}</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn :to="'/trashed?m='+viewName" color="error" flat class="mb-2">PAPELERA</v-btn>
  </v-toolbar>
  <v-container fluid>
    <v-card>
      <v-card-title>
        <v-layout row wrap>
          <v-flex offset-xs8 xs4>
              <v-text-field
                v-model="search"
                append-icon="search"
                label="Buscar:"
                single-line
                hide-details
              ></v-text-field>
          </v-flex>
        </v-layout>
      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="list"
        class="elevation-1"
        no-data-text="No hubo resultados"
        :search="search"
      >
      <template slot="items" slot-scope="props">
        <tr @click="props.expanded = !props.expanded">
          <td>{{ props.item.cookie }}</td>
          <td>{{ props.item.data.service.name }}</td>
          <td>{{ props.item.data.information.company.value }}</td>
          <td>{{ props.item.data.information.phone.value }}</td>
          <td>{{ props.item.data.information.email.value }}</td>
          <td class="justify-center layout px-0">
            <v-icon
              title
              @click="deleteItem(props.item)" color="error"
            >
              delete
            </v-icon>
          </td>
        </tr>
      </template>
      <template slot="expand" slot-scope="props">
        <v-container grid-list-md text-xs-center>
          <!-- props.item -->
          <v-layout row wrap>
              <!-- {{props.item}} -->
              <table>
                <thead>
                  <tr>
                    <th colspan="4">
                      <b>Datos del cliente</b>
                    </th>
                    <th colspan="4">
                      <b>Informacion del brief</b>
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td colspan="4" class="text-xs-left">
                      <tr>
                        <td class="text-xs-left"><b>Nombre: </b> {{props.item.client.name}}</td>
                      </tr>
                      <tr>
                        <td class="text-xs-left"><b>Email: </b> {{props.item.client.email}}</td>
                      </tr>
                      <tr>
                        <td class="text-xs-left"><b>Compañia: </b> {{props.item.client.company}}</td>
                      </tr>
                      <tr>
                        <td class="text-xs-left"><b>Teléfono: </b> {{props.item.client.phone}}</td>
                      </tr>
                    </td>
                    <td colspan="4">
                      <tr>
                        <td class="text-xs-left"><b>Servicio:</b> {{props.item.data.service.name}} </td>
                      </tr>
                      <tr v-if="item.value != [] && item.value != '' && item.value" v-for="(item, i) in props.item.data.information" :key="i">
                        <td class="text-xs-left" v-if="!Array.isArray(item.value)">
                          <b>{{item.label}}: </b> {{item.value}}
                        </td>
                        <td class="text-xs-left" v-if="Array.isArray(item.value)">
                          <b>{{item.label}}:</b> 
                          <span v-for="(subitem, oi) in item.value" :key="oi">{{subitem}}</span>
                        </td>
                      </tr>
                    </td>
                  </tr>
                  <tr>
                    <th colspan="7">Diseños</th>
                  </tr>
                  <tr> 
                    <td v-if="props.item.data.designs.length">
                      <span v-for="design in props.item.data.designs" :key="design">  
                        <img width="200px" style="margin-right:40px;" :src="'http://liderlogos.com/images/brief/designs/' + design + '.jpg'">
                      </span>
                    </td>
                    <td v-if="!props.item.data.designs.length">No seleccionó diseños</td>
                  </tr>
                  <tr>
                    <th colspan="4">Colores: </th>
                    <th colspan="3">Estilos</th>
                  </tr>
                  <tr>
                    <td colspan="4" v-if="props.item.data.colors.length">
                      <span v-for="color in props.item.data.colors" :key="color" class="mr-2">
                        <img width="30px" :src="'http://liderlogos.com/images/brief/colors/' + color + '.png'">
                      </span>
                    </td>
                    <td colspan="4" v-if="!props.item.data.colors.length">No seleccionó colores</td>

                    <td colspan="3">
                      <div class="mb-3 style-range" v-for="(style, i) in props.item.data.styles" :key="i">
                        <b>{{espStyles[i][0]}}</b>
                        <b class="mx-3">{{style}}</b>
                        <b>{{espStyles[i][1]}}</b> 
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
          </v-layout>
        </v-container>
      </template>
    </v-data-table>
    </v-card>
  </v-container>
</div>
</template>

<script lang="js">
  export default {
    name: 'briefs',
    props: ['search'],
    created () {
      this.$store.dispatch('getAll', {state: this.viewName})
    },
    mounted () {
    },
    data () {
      return {
        selectErrors: [],
        pagination: {},
        dialog: false,
        editedIndex: -1,
        viewNameESP: 'Briefs',
        espStyles: {
          'clasic_modern': ['Clásico', 'Moderno'],
          'mature_youthful': ['Adulto', 'Juvenil'],
          'abstract_literal': ['Abstracto', 'Literal'],
          'geometric_organic': ['Geométrico', 'Orgánico'],
          'feminine_masculine': ['Femenino', 'Masculino'],
          'playful_sophisticated': ['Juguetón', 'Sofisticado'],
          'economical_luxurious': ['Económico', 'Lujoso']
        }
      }
    },
    methods: {
      deleteItem (item) {
        item.index = this.list.indexOf(item)
        let params = {
          state: this.viewName,
          item: item
        }

        confirm('Esta seguro que desea eliminar este elemento?') && this.$store.dispatch('deleteOne', params)
      },
      close () {
        this.dialog = false
        setTimeout(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        }, 300)
      }
    },
    watch: {
      dialog (val) {
        val || this.close()
      }
    },
    computed: {
      headers () {
        return this.$store.state[this.viewName].struct
      },
      list () {
        return this.$store.getters.getAll('briefs')
      },
      pages () {
        if (this.pagination.rowsPerPage == null ||
          this.pagination.totalItems == null
        ) return 0

        return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
      },
      viewName () {
        return this.$route.name
      },
      defaultItem () {
        return this.$store.state[this.viewName].defaultItem
      },
      editedItem: {
        get () {
          return this.$store.state[this.viewName].editedItem
        },
        set (value) {
          this.$store.state[this.viewName].editedItem = value
          return this.$store.state[this.viewName].editedItem
        }
      }
    }
}
</script>

<style scoped>

.theme--light.v-table th {
    color: rgb(0, 0, 0);
}

.style-range{ display:flex; justify-content: space-around; }
.style-range > :first-child, .style-range > :last-child { width: 100px; }

</style>
