<template class="briefs">
<div>
  <v-toolbar flat color="white">
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
          <td>{{ props.item.data.information.names.value }}</td>
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
                    <th colspan="3">
                      <b>Cliente</b>
                    </th>
                    <th colspan="4">
                      <b>Informacion</b>
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td colspan="3" class="text-xs-left">
                      <tr v-for="(item, i) in props.item.client" :key="i">
                        <td class="text-xs-left"><b class="text-capitalize">{{i}}: </b> {{item}}</td>
                      </tr>
                    </td>
                    <td colspan="4">
                      <tr v-for="(item, i) in props.item.data.information" :key="i">
                        <td class="text-xs-left"><b>{{item.label}}: </b> {{item.value}}</td>
                      </tr>
                    </td>
                  </tr>
                  <tr>
                    <th colspan="2">Colores: </th>
                    <th colspan="2">Diseños</th>
                    <th colspan="2">Estilos</th>
                  </tr>
                  <tr>
                    <td colspan="2">{{props.item.data.colors}}</td>
                    <td colspan="2">{{props.item.data.designs}}</td>
                    <td colspan="2">
                      <div class="mb-3" v-for="(style, i) in props.item.data.styles" :key="i">
                        <b class="text-capitalize">{{i}}: </b> {{style}}
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
        viewNameESP: 'Briefs'
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

</style>
