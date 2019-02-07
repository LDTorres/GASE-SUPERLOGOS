<template class="sectors">
<div>
  <v-toolbar flat color="white" class="sticky-nav">
    <v-toolbar-title hidden-md-and-down class="text-capitalize">LEADS</v-toolbar-title>
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
          <td>{{ props.item.id }}</td>
          <td >{{ props.item.email }}</td>
          <td >{{ props.item.phone }}</td>
          <td >{{ props.item.schedule }}</td>
          <td >{{ props.item.medium }}</td>
          <td >{{ props.item.countries.name }}</td>
          <td class="justify-center layout">
            <!-- <v-icon
              title
              class="mr-2" color="primary"
              @click="makeUser(props.item)"
            >
              people
            </v-icon> -->
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
          <v-layout row wrap>
            <v-flex xs5 text-xs-left>
              <p><b>Mensaje:</b> {{props.item.message}}</p>
            </v-flex>
            <v-flex xs5 text-xs-left>
              <p><b>Campaña:</b> {{props.item.campaign}}</p>
              <p><b>Medio:</b> {{props.item.medium}}</p>
              <p><b>Fuente:</b> {{props.item.source}}</p>
            </v-flex>
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
    name: 'sectors',
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
        editedIndex: -1
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
      makeUser (user) {
        console.log(user)
      }
    },
    computed: {
      headers () {
        return this.$store.state[this.viewName].struct
      },
      list () {
        return this.$store.getters.getAll('leads')
      },
      pages () {
        if (this.pagination.rowsPerPage == null ||
          this.pagination.totalItems == null
        ) return 0

        return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
      },
      viewName () {
        return this.$route.name
      }
    }
}
</script>