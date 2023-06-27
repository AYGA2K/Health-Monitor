<template>
  <div>
    <q-layout view="lHh Lpr lFf">
      <q-page-container>
        <q-page class="flex flex-center bg-grey-2">
          <q-card class="q-pa-md shadow-2 my_card" bordered>
            <q-card-section class="text-center">
              <div class="text-grey-9 text-h5 text-weight-bold">Sign up</div>
              <div class="text-grey-8">Create a new account</div>
            </q-card-section>
            <q-card-section>
              <q-input dense outlined filled v-model="username" label="Username"></q-input>
              <q-input dense outlined filled class="q-mt-md" v-model="email" label="Email Address"></q-input>
              <q-input dense outlined filled class="q-mt-md" v-model="password" :type="showPassword ? 'text' : 'password'"
                label="Password">
                <template v-slot:append>
                  <q-icon :name="showPassword ? 'visibility_off' : 'visibility'" class="cursor-pointer"
                    @click="togglePasswordVisibility"></q-icon>
                </template>
              </q-input>
              <q-input dense outlined filled class="q-mt-md" v-model="confirmPassword"
                :type="showPassword ? 'text' : 'password'" label="Confirm Password"></q-input>
            </q-card-section>
            <q-card-section>
              <q-btn @click="signup" style="border-radius: 8px;" color="dark" rounded size="md" label="Sign up" no-caps
                class="full-width"></q-btn>
            </q-card-section>
            <q-card-section class="text-center q-pt-none">
              <div class="text-grey-8">Already have an account?
                <span @click="goToSignIn" class="text-dark cursor-pointer  text-weight-bold"
                  style="text-decoration: none">Sign in.</span>
              </div>
            </q-card-section>
          </q-card>
        </q-page>
      </q-page-container>
    </q-layout>
  </div>
</template>

<script lang="ts">
import { ip } from 'src/ip_adress';
export default {
  data() {
    return {
      email: '' as string,
      username: '' as string,
      password: '' as string,
      confirmPassword: '' as string,
      showPassword: false,
    }
  },
  methods: {
    togglePasswordVisibility() {
      this.showPassword = !this.showPassword;
    },
    signup() {
      if (this.email && this.password && this.username) {


        if (this.password == this.confirmPassword) {
          const payload = {
            name: this.username,
            password: this.password,
            email: this.email,
          };

          this.$axios
            .post('http://' + ip + ':3000/api/signup', payload)
            .then(response => {
              console.log(response)
              // Handle success response
              this.$q.notify({
                type: 'positive',
                message: 'Signed up successfully'

              })
              this.goToSignIn()
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
            message: ' Passwords do not match. Please try again. '
          })
        }
      }
      else {
        this.$q.notify({
          type: 'negative',
          message: 'Please fill in all the required fields.'
        })

      }
    },
    goToSignIn() {
      this.$router.push("/sign-in");
    },
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

