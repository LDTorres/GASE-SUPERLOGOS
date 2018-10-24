<template class="services">
<div>
  <v-toolbar flat color="white">
      <v-toolbar-title hidden-md-and-down class="text-capitalize">{{ viewNameESP }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn :to="'/trashed?m='+viewName" color="error" flat class="mb-2">PAPELERA</v-btn>
      <v-dialog v-model="dialog" max-width="500px">
        <v-btn slot="activator" color="primary" flat class="mb-2">Nuevos {{ viewNameESP }}</v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">{{ formTitle }}</span>
          </v-card-title>

          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12>
                  <v-text-field type="text" name="Nombre" v-validate="'required'" v-model="editedItem.name" label="Nombre"></v-text-field>                   
                  <span v-show="errors.has('Nombre')">{{ errors.first('Nombre') }}</span>
                </v-flex>
                <v-flex xs12>
                  <v-text-field type="number" name="Porcentage" v-validate="'required|numeric|max:2'" v-model="editedItem.percentage" label="Porcentage"></v-text-field>
                  <span v-show="errors.has('Porcentage')">{{ errors.first('Porcentage') }}</span>
                </v-flex>

                <!-- Si es nuevo -->
                <v-flex v-if="editedIndex === -1" xs3 v-for="(currency) in currencies" :key="currency.id">
                  <v-text-field type="number" :name="currency.iso" v-validate="'required|numeric|max:6'" v-model="currency.price" :label="currency.iso"></v-text-field>
                  <span v-show="errors.has(currency.iso)">{{ errors.first(currency.iso) }}</span>
                </v-flex>

                <!-- Si es editado -->
                <v-flex v-if="editedIndex !== -1" xs3 v-for="(price) in editedItem.prices" :key="price.id">
                  <v-text-field type="text" v-validate="'required|max:6'" v-model="price.value" :label="price.currency.iso"></v-text-field>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="error" flat  @click.native="close">Cancelar</v-btn>
            <v-btn color="primary" flat  @click.native="save" :disabled="errors.count() > 0">Guardar</v-btn>
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
        <td>{{ props.item.id }}</td>
        <td >{{ props.item.name }}</td>
        <td >{{ props.item.slug }}</td>
        <td >{{ props.item.percentage }}</td>
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
    name: 'services',
    props: ['search'],
    created () {
      this.$store.dispatch('getAll', {state: this.viewName})
      this.$store.dispatch('getAll', {state: 'currencies'})
    },
    mounted () {
    },
    data () {
      return {
        selectErrors: [],
        pagination: {},
        dialog: false,
        editedIndex: -1,
        viewNameESP: 'Servicios'
      }
    },
    methods: {
      editItem (item) {
        this.editedIndex = this.list.indexOf(item)
        this.editedItem = Object.assign({}, item)
        this.dialog = true
      },
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
      },
      save () {
        this.$validator.validate().then(result => {
          if (!result) {
            alert('Llene los campos correctamente.')
          } else {
            this.editedItem.percentage = parseFloat(this.editedItem.percentage)

            let params = {
              state: this.viewName,
              item: this.editedItem
            }

            if (this.editedIndex > -1) {
              for (let i = 0; i < this.editedItem.prices.length; i++) {
                const price = this.editedItem.prices[i]
                price.value = parseFloat(price.value)
              }
              this.$store.dispatch('updateOne', params)
            } else {
              this.editedItem.prices = []

              this.currencies.forEach(currency => {
                let price = {
                  currency: {
                    id: currency.id
                  },
                  value: parseFloat(currency.price)
                }

                this.editedItem.prices.push(price)
              })

              params.editedItem = this.editedItem

              this.$store.dispatch('create', params)
            }
            this.close()
          }
        })
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
        return this.$store.getters.getAll('services')
      },
      currencies () {
        return this.$store.getters.getAll('currencies')
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
      formTitle () {
        var title = this.editedIndex === -1 ? 'Nuevo ' : 'Editar '
        return title + this.viewNameESP
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