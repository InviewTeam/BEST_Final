<template>
  <div id="app">
    <b-navbar toggleable="lg" type="dark" variant="info">
      <b-navbar-nav>
        <b-nav-item to="/">Home</b-nav-item>
        <b-nav-item to="/equipments" v-if="isLoggedIn">Equipments</b-nav-item>
      </b-navbar-nav>
      <b-navbar-nav class="ml-auto">
      <auth v-if="!isLoggedIn"/>
       <div class="container" v-else>
      <div class="row">
      <div class="col-sm-12">
        <span class="mr-auto">{{getUsername}}</span>
        <button type="button" class="btn btn-dark btn-sm"  @click="logout">Logout</button>
        </div>
      </div>
       </div>
      </b-navbar-nav>
    </b-navbar>
    <router-view/>
  </div>
</template>

<script>
import axios from 'axios';
import Auth from './components/Auth/Auth';

export default {
  name: 'App',
  components: {
    auth: Auth,
  },
  computed: {
    isLoggedIn() {
      return this.$store.getters.isLoggedIn;
    },
    getUsername() {
      return this.$store.getters.getUsername();
    },
  },
  methods: {
    logout() {
      this.$store.dispatch('logout')
        .then(() => this.$router.push('/'));
    },
    testConnection() {
      axios.get('',
        {
          headers: { Authorization: `Bearer ${this.$store.state.token}` },
        })
        .catch((error) => {
          if (error.response.status === 401) {
            this.$store.dispatch('logout');
          }
        });
    },
  },
  created() {
    this.testConnection();
  },
};
</script>

<style>
</style>
