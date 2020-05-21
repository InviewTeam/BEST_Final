<template>
      <div class="container">
      <div class="row">
      <div class="col-sm-12">
       <button type="button"
                class="btn  btn-dark btn-sm"
                v-b-modal.login-modal
       >Sign in</button>
        <button type="button"
                class="btn btn-dark btn-sm"
                v-b-modal.register-modal
       >Sign up</button>
      </div>
      </div>

      <!-- Login -->
      <b-modal ref="LoginModal"
             id="login-modal"
             title="Login" hide-footer
              hide-backdrop
              centered>
      <b-form @submit="loginSubmit" @reset="loginReset" class="w-100">
      <b-form-group id="form-username-group"
                    label="Username:"
                    label-for="form-username-input">
          <b-form-input id="form-username-input"
                        type="text"
                        v-model="Login.username"
                        required
                        placeholder="Enter username">
          </b-form-input>
        </b-form-group>
        <b-form-group id="form-password-group"
                      label="Password:"
                      label-for="form-password-input">
            <b-form-input id="form-password-input"
                          type="password"
                          v-model="Login.password"
                          required
                          placeholder="Enter password">
            </b-form-input>
        </b-form-group>
        <b-button type="submit" variant="primary">Submit</b-button>
        <b-button type="reset" variant="danger">Reset</b-button>
      </b-form>
    </b-modal>
        <!--Register-->
         <b-modal ref="RegisterModal"
             id="register-modal"
             title="Register" hide-footer
              hide-backdrop
              centered>
      <b-form @submit="registerSubmit" @reset="registerReset" class="w-100">
      <b-form-group id="form-username-register-group"
                    label="Username:"
                    label-for="form-username-register-input">
          <b-form-input id="form-username-register-input"
                        type="text"
                        v-model="Register.username"
                        required
                        placeholder="Enter username">
          </b-form-input>
        </b-form-group>
        <b-form-group id="form-password-register-group"
                      label="Password:"
                      label-for="form-password-register-input">
            <b-form-input id="form-password-register-input"
                          type="password"
                          v-model="Register.password"
                          required
                          placeholder="Enter password">
            </b-form-input>
        </b-form-group>
        <b-button type="submit" variant="primary">Submit</b-button>
        <b-button type="reset" variant="danger">Reset</b-button>
      </b-form>
    </b-modal>
      </div>
</template>

<script>
export default {
  data() {
    return {
      projects: [],
      Login: {
        username: null,
        password: null,
      },
      Register: {
        username: null,
        password: null,
      },
    };
  },
  methods: {
    initForm() {
      this.Login.username = '';
      this.Login.password = '';
    },
    loginSubmit(evt) {
      evt.preventDefault();
      this.$refs.LoginModal.hide();
      const username = this.Login.username;
      const password = this.Login.password;
      this.$store.dispatch('login', { username, password })
        .then(() => console.log('ok'))
        .catch(err => console.log(err));
      this.initForm();
    },
    registerSubmit(evt) {
      evt.preventDefault();
      this.$refs.RegisterModal.hide();
      const payload = {
        username: this.Register.username,
        password: this.Register.password,
      };
      this.$store.dispatch('register', payload)
        .then(() => this.$router.push('/'))
        .catch(err => console.log(err));
      this.initForm();
    },
    registerReset(evt) {
      evt.preventDefault();
      this.$refs.RegisterModal.hide();
      this.initForm();
    },
    loginReset(evt) {
      evt.preventDefault();
      this.$refs.LoginModal.hide();
      this.initForm();
    },
  },
  created() {
    this.initForm();
  },
};
</script>

<style scoped>

</style>
