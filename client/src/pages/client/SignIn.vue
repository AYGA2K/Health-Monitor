<template>
  <div>
    <q-layout view="lHh Lpr lFf">
      <q-page-container>
        <q-page class="flex flex-center bg-grey-2">
          <q-card class="q-pa-md shadow-2 my_card" bordered>
            <q-card-section class="text-center">
              <div class="text-grey-9 text-h5 text-weight-bold">Sign in</div>
              <div class="text-grey-8">Sign in below to access your account</div>
            </q-card-section>
            <q-card-section>
              <q-input dense outlined v-model="email" label="Email Address"></q-input>
              <q-input dense outlined class="q-mt-md" v-model="password" type="password" label="Password"></q-input>
            </q-card-section>
            <q-card-section>
              <q-btn style="
  border-radius: 8px;" color="dark" rounded size="md" @click="signin" label="Sign in" no-caps
                class="full-width"></q-btn>
            </q-card-section>
            <q-card-section class="text-center q-pt-none">
              <div class="text-grey-8">Don't have an account yet?
                <span @click="SignUp" class="cursor-pointer text-dark text-weight-bold" style="text-decoration: none">Sign
                  up.</span>
              </div>
            </q-card-section>

          </q-card>
        </q-page>
      </q-page-container>
    </q-layout>
  </div>
</template>
<script lang="ts">
import { LogInState } from 'src/router/logIn';
import { ip } from 'src/ip_adress';
export default {
  data() {
    return {
      email: '' as string,
      password: '' as string
    }
  },
  methods: {
    signin() {
      if (this.email && this.password) {



        const payload = {
          password: this.password,
          email: this.email,
        };

        this.$axios
          .post('http://' + ip + ':3000/api/login', payload)
          .then(response => {
            console.log(response)
            // Handle success response
            this.$q.notify({
              type: 'positive',
              message: 'Signed in successfully'

            })
            localStorage.setItem("access_token", response.data.access_token);
            localStorage.setItem("refresh_token", response.data.refresh_token);

            LogInState.isLogedIn = true

            this.$router.push("/")
          })
          .catch(error => {
            this.$q.notify({
              type: 'negative',
              message: error

            })
            // Handle error response
          });


      }
      else {
        this.$q.notify({
          type: 'negative',
          message: 'Please fill in all the required fields.'
        })

      }
    },

    SignUp() {
      this.$router.push("/sign-up")
    }

  },

}
</script>
<style>
.my_card {
  width: 25rem;
  border-radius: 8px;
  box-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1);
}
</style>
