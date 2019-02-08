<template class="sectors">
<div>
  <v-toolbar flat color="white" class="sticky-nav">
    <v-toolbar-title hidden-md-and-down class="text-capitalize">SAFETY PAY</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn color="success" class="mb-2" @click="confirmPayments()" v-if="list.length">CONFIRMAR TODOS LOS PAGOS</v-btn>
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
          <td>{{ props.item.CreationDateTime.split('T')[0] }}</td>
          <td >{{ props.item.OperationID }}</td>
          <td style="cursor:pointer;" @click="$router.push('/orders?q=' + props.item.MerchantOrderID)">{{ props.item.MerchantOrderID }}</td>
          <td >{{ props.item.OperationStatus == '102' ? 'Pagado' : 'Confirmado' }}</td>
          
          <td class="justify-center layout">
            <v-icon
              title
              class="mr-2" color="primary"
              @click="confirm([props.item])"
            >
              check
            </v-icon>
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
    name: 'safetypay',
    props: ['search'],
    created () {
      this.$store.dispatch('safetypay/getAll')
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
      confirm (payments) {
        this.$store.dispatch('safetypay/confirm', payments)
      },
      confirmPayments () {
        this.confirm(this.list)
      }
    },
    computed: {
      headers () {
        return this.$store.state[this.viewName].struct
      },
      list () {
        return this.$store.getters.getAll('safetypay')
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