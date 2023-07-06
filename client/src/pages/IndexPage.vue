<template>
  <q-input borderless class="q-mx-xl q-my-sm" dense debounce="300" v-model="filter" placeholder="Search">
    <template v-slot:append>
      <q-icon name="search" />
    </template>
  </q-input>
  <q-table flat bordered title="Clients" :rows="clients" :columns="columns" row-key="name" :filter="filter">
    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td v-for="column in columns" :key="column.name" :props="props">
          <template v-if="column.name !== 'actions'">
            {{ props.row[column.field] }}
          </template>
          <template v-else>
            <q-btn size="sm" round class="q-mx-sm" @click="deleteClient(props.row.id)" color="red-14" icon="delete" />

          </template>
        </q-td>
      </q-tr>
    </template>
  </q-table>
</template>

<script lang="ts">
import { ip } from "src/ip_adress"
export default {
  data() {
    return {
      columns: [
        {
          name: 'mac',
          required: true,
          label: 'Mac',
          align: 'left',
          field: 'mac',
          sortable: true
        },
        { name: 'name', align: 'center', label: 'Name', field: 'name', sortable: true },
        { name: 'email', label: 'Email', field: 'email', sortable: true },
        { name: 'address', label: 'Address', field: 'address' },
        { name: 'actions', label: 'Actions' }
      ],
      clients: [],
      filter: '',
    };
  },
  mounted() {
    this.fetchClients();
  },
  methods: {
    fetchClients() {

      this.$axios.get('http://' + ip + ':3000/api/users').then(response => this.clients = response.data
      ).catch(err => console.log(err));
      console.log(this.clients)
    },
    deleteClient(id: number) {
      this.$axios.delete('http://' + ip + ':3000/api/user' + id).then(response => {
        console.log(response);
        this.fetchClients();
        this.$q.notify({
          type: 'positive',
          message: 'Deleted successfully'
        })
      }).catch(err => {
        console.log(err)
      })
    },
    editClient(id: number) {
      console.log(id)
      this.$router.push(`/edit-client/${id}`);
    },


    viewClient(id: number) {
      console.log(id)
      this.$router.push(`/view-client/${id}`);
    }
  }
};
</script>
