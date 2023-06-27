<template>
  <div class="q-px-xl row justify-center q-py-sm">
    <div class=" col-10">
      <q-input v-model="mac" label="MAC Address" class="q-ma-sm" />
      <q-input v-model="nom" label="Last Name" class="q-ma-sm" />
      <q-input v-model="prenom" label="First Name" class="q-ma-sm" />
      <q-input v-model="tel" label="Phone Number" class="q-ma-sm" />
      <q-input v-model="adresse" label="Address" class="q-ma-sm" />
      <q-input v-model="mail" label="Email" class="q-ma-sm" />
      <q-btn color="primary" label="Submit" @click="saveClient" class="q-ma-sm float-right" />

    </div>
  </div>
</template>
<script lang="ts">
import axios from 'src/boot/axios';
import { useQuasar } from 'quasar'
export default {
  data() {

    return {
      mac: '',
      nom: '',
      prenom: '',
      tel: '',
      adresse: '',
      mail: ''
    }
  },
  methods: {
    saveClient() {
      const $q = useQuasar()

      if (!this.mac || !this.nom || !this.prenom) {
        this.$q.notify({
          type: 'negative',
          message: 'Please fill in all required fields.'
        });
        return;
      }
      const payload = {
        mac: this.mac,
        nom: this.nom,
        prenom: this.prenom,
        tel: this.tel,
        adresse: this.adresse,
        mail: this.mail
      };


      this.$axios
        .post('http://154.49.137.28:8080/saveClient', payload)
        .then(response => {
          console.log('Client saved successfully', response.data);
          this.mac = '';
          this.nom = '';
          this.prenom = '';
          this.tel = '';
          this.adresse = '';
          this.mail = '';
          // Handle success response
          this.$q.notify({
            type: 'positive',
            message: 'Added successfully'
          })
        })
        .catch(error => {
          console.error('Error saving client', error);
          // Handle error response
        });
    }

  }
};
</script>
