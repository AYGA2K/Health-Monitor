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
            <q-btn size="sm" round color="blue-8" @click="editClient(props.row.id)" icon="edit" />
            <q-btn size="sm" round class="q-mx-sm" color="positive" @click="viewClient(props.row.id)" icon="visibility" />

          </template>
        </q-td>
      </q-tr>
    </template>
  </q-table>
</template>

<script lang="ts">
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
            { name: 'nom', align: 'center', label: 'Nom', field: 'nom', sortable: true },
        { name: 'prenom', label: 'Prenom', field: 'prenom', sortable: true },
        { name: 'tel', label: 'Tel', field: 'tel' },
        { name: 'adresse', label: 'Adresse', field: 'adresse' },
        { name: 'mail', label: 'Mail', field: 'mail' },
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
      this.$axios.get('http://154.49.137.28:8080/listClients').then(response => this.clients = response.data).catch(err => console.log(err));
    },
    deleteClient(id: number) {
      this.$axios.delete('http://154.49.137.28:8080/deleteClient/' + id).then(response => {
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
