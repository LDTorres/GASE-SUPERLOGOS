<template class="portfolios">
  <div>
    <v-toolbar flat color="white" class="sticky-nav">
      <v-toolbar-title hidden-md-and-down class="text-capitalize">PROYECTOS</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn :to="'/trashed?m='+viewName" color="error" flat class="mb-2">PAPELERA</v-btn>
      <v-dialog v-model="dialog" max-width="700px">
        <v-btn slot="activator" color="primary" flat class="mb-2">Nuevo Proyecto</v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">{{ formTitle }}</span>
          </v-card-title>

          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs6>
                  <v-text-field
                    type="text"
                    name="Nombre"
                    v-validate="'required'"
                    v-model="editedItem.name"
                    label="Nombre del Portafolio"
                  ></v-text-field>
                  <span v-show="errors.has('Nombre')">{{ errors.first('Nombre') }}</span>
                </v-flex>

                <v-flex xs12 md6>
                  <v-layout wrap>
                    <v-flex xs12>
                      <v-select
                        v-model="editedItem.client"
                        :items="clients"
                        item-text="name"
                        :error-messages="selectErrors"
                        return-object
                        label="Cliente"
                        required
                        name="Cliente"
                        v-validate="'required'"
                      ></v-select>
                      <span v-show="errors.has('Cliente')">{{ errors.first('Cliente') }}</span>
                    </v-flex>
                  </v-layout>
                </v-flex>

                <v-flex xs6>
                  <v-text-field
                    type="text"
                    name="notifications_email"
                    v-validate="'required'"
                    v-model="editedItem.notifications_email"
                    label="Email de notificaciones"
                  ></v-text-field>
                  <span v-show="errors.has('notifications_email')">{{ errors.first('notifications_email') }}</span>
                </v-flex>
                  
                <v-flex xs6>
                  <v-text-field
                    type="text"
                    name="agile_id"
                    v-validate="'required'"
                    v-model="editedItem.agile_id"
                    label="Agile ID"
                  ></v-text-field>
                  <span v-show="errors.has('agile_id')">{{ errors.first('agile_id') }}</span>
                </v-flex>

                 <v-flex xs12 v-if="editedIndex == -1">
                  <v-checkbox
                    label="Cerrado"
                    v-model="editedItem.closed"
                    name="Cerrado"
                    v-validate="'required'"
                  ></v-checkbox>
                  <span v-show="errors.has('Cerrado')">{{ errors.first('Cerrado') }}</span>
                </v-flex>
                <v-flex xs12 v-if="editedIndex != -1">
                  <v-checkbox
                    label="Publico"
                    v-model="editedItem.closed"
                    name="Publico"
                    v-validate="'required'"
                  ></v-checkbox>
                  <span v-show="errors.has('Publico')">{{ errors.first('Publico') }}</span>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="error" flat @click.native="close">Cancelar</v-btn>
            <v-btn color="primary" flat @click.native="save" :disabled="errors.count() > 0">Guardar</v-btn>
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
            <tr>
              <td>{{ props.item.id }}</td>
              <td>{{ props.item.name }}</td>
              <td>{{ props.item.client.name }}</td>
              <td>{{ props.item.closed ? 'Cerrado' : 'Abierto' }}</td>
              <td>{{ props.item.notifications_email }}</td>
              <td>{{ props.item.agile_id || 'No colocado'}}</td>
              <td class="justify-center layout px-0">
                <v-icon title class="mr-2" color="success" @click="editItem(props.item)">link</v-icon>
                <v-icon title class="mr-2" color="primary" @click="editItem(props.item)">edit</v-icon>
                <v-icon title @click="deleteItem(props.item)" color="error">delete</v-icon>
              </td>
            </tr>
          </template>
        </v-data-table>
      </v-card>
    </v-container>
  </div>
</template>

<script lang="js">
  export default {
    name: 'portfolios',
    props: ['search'],
    created () {
      this.$store.dispatch('getAll', {state: this.viewName})
      this.$store.dispatch('getAll', {state: 'clients'})
    },
    mounted () {
    },
    data () {
      return {
        // Cambiar dependiendo del enviroment
        urlHosting: process.env.NODE_ENV === 'development' ? 'http://localhost:9090' : 'http://api.liderlogos.com',
        selectErrors: [],
        pagination: {},
        dialog: false,
        editedIndex: -1,
        imagePreviewUrl: '',
        priorityImage: 0,
        selectedImage: ''
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
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      },
      save () {
        this.$validator.validate().then(result => {
          if (!result) {
            alert('Llene los campos correctamente.')
          } else {
            let params = {
              state: this.viewName,
              item: this.editedItem
            }

            if (this.editedIndex > -1) {
              params.edited = true
              this.$store.dispatch('updateOne', params)
            } else {
              this.$store.dispatch('create', params)
            }
            this.imagePreviewUrl = ''
            this.close()
          }
        })
      },
      uploadImage (item) {
        item['files'] = this.selectedImage
        item.priorityImage = this.priorityImage
        let params = {
          state: this.viewName,
          item: item
        }

        this.imagePreviewUrl = ''
        this.$store.dispatch('portfolios/uploadImage', params)
      },
      imagePriority (img) {
        let params = {
          state: this.viewName,
          item: img
        }

        this.$store.dispatch('portfolios/imagePriority', params)
      },
      deleteImage (img, item, i) {
        let params = {
          state: this.viewName,
          item: item,
          indexImage: i
        }

        this.$store.dispatch('portfolios/imageDelete', params)
      },
      triggerFileButton () {
        document.querySelector('.input-file-preview').click()
      },
      fileSelectedAddNew (e) {
        this.previewImage(e.target.files)
        this.selectedImage = e.target.files
      },
      fileSelected (e) {
        this.editedItem['files'] = e.target.files
      },
      previewImage (files) {
        var v = this

        if (files && files[0]) {
          var reader = new FileReader()

          reader.onload = (e) => {
            v.imagePreviewUrl = e.target.result
          }

          reader.readAsDataURL(files[0])
        }
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
        return this.$store.getters.getAll('projects')
      },
      clients () {
        return this.$store.getters.getAll('clients')
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
        return title + 'Proyecto'
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
input.input-file-preview {
  position: absolute;
  width: 0;
  left: 0;
  top: 0;
  visibility: hidden;
}

.without-image {
  background: url("../assets/add-circular-button.png");
  background-repeat: no-repeat;
  background-position: 50%;
}
</style>
