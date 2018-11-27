<template class="portfolios">
<div>
  <v-toolbar flat color="white" class="sticky-nav">
      <v-toolbar-title hidden-md-and-down class="text-capitalize">{{ viewNameESP }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn :to="'/trashed?m='+viewName" color="error" flat class="mb-2">PAPELERA</v-btn>
      <v-dialog v-model="dialog" max-width="700px">
        <v-btn slot="activator" color="primary" flat class="mb-2">Nuevos {{ viewNameESP }}</v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">{{ formTitle }}</span>
          </v-card-title>

          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 md6>
                  <v-layout wrap>
                    <v-flex xs12>
                      <v-text-field type="text" name="Nombre" v-validate="'required'" v-model="editedItem.name" label="Nombre del Portafolio"></v-text-field>
                      <span v-show="errors.has('Nombre')">{{ errors.first('Nombre') }}</span>
                    </v-flex>
                    <v-flex xs12>
                      <v-text-field type="text" name="Cliente" v-validate="'required'" v-model="editedItem.client" label="Cliente"></v-text-field>
                      <span v-show="errors.has('Cliente')">{{ errors.first('Cliente') }}</span>
                    </v-flex>
                    <v-flex xs12>
                      <v-text-field type="number" name="Prioridad" v-validate="'required|numeric|max:2'" v-model.number="editedItem.priority" label="Prioridad"></v-text-field>
                      <span v-show="errors.has('Prioridad')">{{ errors.first('Prioridad') }}</span>
                    </v-flex>
                  </v-layout>
                </v-flex>
                <v-flex xs12 md6>
                  <v-layout wrap>
                    <v-flex xs12>
                      <v-select
                        v-model="editedItem.location"
                        :items="locations"
                        item-text="name"
                        :error-messages="selectErrors"
                        return-object
                        label="Pais"
                        required
                        name="Pais" 
                        v-validate="'required'"
                      ></v-select>
                      <span v-show="errors.has('Pais')">{{ errors.first('Pais') }}</span>
                    </v-flex>
                    <v-flex xs12>
                      <v-select
                        v-model="editedItem.service"
                        :items="services"
                        item-text="name"
                        :error-messages="selectErrors"
                        return-object
                        label="Servicio"
                        required
                        name="Servicio" 
                        v-validate="'required'"
                      ></v-select>
                      <span v-show="errors.has('Servicio')">{{ errors.first('Servicio') }}</span>
                    </v-flex>
                    <v-flex xs12>
                      <v-select
                        v-model="editedItem.activity"
                        :items="activities"
                        item-text="name"
                        :error-messages="selectErrors"
                        return-object
                        label="Sector o Actividades"
                        required
                        name="Actividad" 
                        v-validate="'required'"
                      ></v-select>
                      <span v-show="errors.has('Actividad')">{{ errors.first('Actividad') }}</span>
                    </v-flex>
                  </v-layout>
                </v-flex>
                <v-flex xs6>
                  <v-textarea
                    name="Descripcion"
                    label="Descripción"
                    v-model="editedItem.description"
                  ></v-textarea>
                </v-flex>
                <v-flex xs6 v-if="editedIndex == -1">
                  <p >Cargar Imagenes Jpg o Png:</p>
                  <input type="file" v-validate="'required|size:15000|mimes:image/*'" name="Imagenes" v-on:change="fileSelected" multiple>
                  <br>
                  <br>
                  <span v-show="errors.has('Imagenes')">{{ errors.first('Imagenes') }}</span>
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
        <tr @click="props.expanded = !props.expanded">
          <td>{{ props.item.id }}</td>
          <td >{{ props.item.name }}</td>
          <td >{{ props.item.description }}</td>
          <td >{{ props.item.client }}</td>
          <td >{{ props.item.location.name }}</td>
          <td >{{ props.item.service.name }}</td>
          <td >{{ props.item.activity.name }}</td>
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
        </tr>
      </template>
      <template slot="expand" slot-scope="props">
        <v-container grid-list-md text-xs-center>
          <v-layout row wrap>
            <v-flex
               v-for="(img, i) in props.item.images" :key="img.name + i"
               xs3
            >
              <v-card>
                <v-img
                  :src="urlHosting + img.url"
                  height="200px"
                ></v-img>
                <v-card-actions>
                  <v-flex>
                      <v-text-field type="number" name="Prioridad" v-model.number="img.priority" label="Prioridad"></v-text-field>
                  </v-flex>
                  <v-btn icon @click="imagePriority(img, i)">
                    <v-icon>save</v-icon>
                  </v-btn>
                  <v-spacer></v-spacer>
                  <v-btn icon @click="deleteImage(img, props.item, i)">
                    <v-icon>delete</v-icon>
                  </v-btn>
                </v-card-actions>
              </v-card>
            </v-flex>

            <!-- prueba -->
            <v-flex xs3>
              <v-card>
                <v-img @click="triggerFileButton"
                  :src="imagePreviewUrl"
                  height="200px" :class="{ 'without-image' : imagePreviewUrl == ''}"
                ></v-img>
                <input type="file" class="input-file-preview" name="Imagenes" v-on:change="fileSelectedAddNew">
                <v-card-actions>
                  <v-flex>
                      <v-text-field type="number" name="prioridadnueva" v-validate="'max:2'" :value="'1'" v-model="priorityImage" label="Prioridad"></v-text-field>
                  </v-flex>
                  <v-spacer></v-spacer>
                  <v-btn icon :disabled="errors.first('prioridadnueva')" @click="uploadImage(props.item)">
                      <v-icon>save</v-icon>
                  </v-btn>
                </v-card-actions>
              </v-card>
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
    name: 'portfolios',
    props: ['search'],
    created () {
      this.$store.dispatch('getAll', {state: this.viewName})
      this.$store.dispatch('getAll', {state: 'locations'})
      this.$store.dispatch('getAll', {state: 'services'})
      this.$store.dispatch('getAll', {state: 'activities'})
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
        viewNameESP: 'Portafolios',
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
              this.$store.dispatch('portfolios/updateOne', params)
            } else {
              this.$store.dispatch('portfolios/create', params)
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
        return this.$store.getters.getAll('portfolios')
      },
      locations () {
        return this.$store.getters.getAll('locations')
      },
      services () {
        return this.$store.getters.getAll('services')
      },
      activities () {
        return this.$store.getters.getAll('activities')
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

<style scoped>
  input.input-file-preview {
    position: absolute;
    width: 0;
    left: 0;
    top: 0;
    visibility: hidden;
  }

  .without-image{
    background: url('../assets/add-circular-button.png');
    background-repeat: no-repeat;
    background-position: 50%;
  }
</style>
