<template class="activities">
<div>
    <v-toolbar flat color="white">
      <v-toolbar-title class="text-capitalize">{{ viewName }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-dialog v-model="dialog" max-width="500px">
        <v-btn slot="activator" color="primary" dark class="mb-2">New {{ viewName }}</v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">{{ formTitle }}</span>
          </v-card-title>

          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 sm6 md12>
                  <v-text-field v-model="editedItem.name" label="Name"></v-text-field>
                </v-flex>
                <v-flex xs12 sm6 md12>
                  <v-text-field v-model="editedItem.description" label="Description"></v-text-field>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" flat @click.native="close">Cancel</v-btn>
            <v-btn color="blue darken-1" flat @click.native="save">Save</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-toolbar> 
  <v-container fluid>
    <v-card>
      <v-card-title>
        <v-layout row wrap>
          <v-flex offset-xs8 xs4>
              <v-text-field
                v-model="search"
                append-icon="search"
                label="Search"
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
      no-data-text="No results"
      :search="search"
    >
      <template slot="items" slot-scope="props">
        <td>{{ props.item.id }}</td>
        <td >{{ props.item.name }}</td>
        <td >{{ props.item.description }}</td>
        <td >{{ props.item.sector.name }}</td>
        <td class="justify-center layout px-0">
          <v-icon
            title
            class="mr-2" color="primary"
            @click="editItem(props.item)"
          >
            edit
          </v-icon>
          <v-icon
            title
            @click="deleteItem(props.item)" color="error"
          >
            delete
          </v-icon>
        </td>
      </template>
    </v-data-table>
    </v-card>
  </v-container>
</div>
</template>

<script lang="js">
  export default {
    name: 'activities',
    props: [],
    created () {
      this.$store.dispatch(this.viewName + '/getAll')
    },
    mounted () {
    },
    data () {
      return {
        headers: [
          { text: 'Id', align: 'left', sortable: true, value: 'id' },
          { text: 'Name', value: 'name' },
          { text: 'Description', value: 'description' },
          { text: 'Sector', value: 'sector' },
          {text: 'Actions', align: 'center', value: ''}
        ],
        search: '',
        pagination: {},
        dialog: false,
        editedIndex: -1,
        editedItem: {
          name: '',
          description: 0
        },
        defaultItem: {
          name: '',
          description: 0
        }
      }
    },
    methods: {
      editItem (item) {
        this.editedIndex = this.list.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialog = true
      },

      deleteItem (item) {
        const index = this.list.indexOf(item)
        confirm('Are you sure you want to delete this item?') && this.list.splice(index, 1)
      },

      close () {
        this.dialog = false
        setTimeout(() => {
          this.editedItem = Object.assign({}, this.defaultItem)
          this.editedIndex = -1
        }, 300)
      },

      save () {
        if (this.editedIndex > -1) {
          Object.assign(this.list[this.editedIndex], this.editedItem)
        } else {
          this.list.push(this.editedItem)
        }
        this.close()
      }
    },
    watch: {
      dialog (val) {
        val || this.close()
      }
    },
    computed: {
      list () {
        return this.$store.state[this.viewName].all
      },
      pages () {
        if (this.pagination.rowsPerPage == null ||
          this.pagination.totalItems == null
        ) return 0

        return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
      },
      formTitle () {
        var title = this.editedIndex === -1 ? 'New ' : 'Edit '
        return title + this.viewName
      },
      viewName () {
        return this.$route.name
      }
    }
}
</script>

<style scoped lang="css">
.activities {
}
</style>
