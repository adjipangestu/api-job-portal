<template>
  <div id="app">
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <div class="navbar-brand" href="#">
        <router-link to="/">Golang</router-link>
      </div>

      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <a class="nav-link">
              <router-link to="/home">Home</router-link>
            </a>
          </li>
          <li class="nav-item active" v-if="isAuthenticated()">
            <a class="nav-link">
              <router-link to="/user">User</router-link>
            </a>
          </li>
          <li class="nav-item active" v-if="!isAuthenticated()">
            <a class="nav-link">
              <router-link to="/login">Login</router-link>
            </a>
          </li>
          <li v-if="isAuthenticated()">
            <a class="nav-link">
              <router-link to="logout" @click.native="logout()">Logout</router-link>
            </a>
          </li>
        </ul>
      </div>
    </nav>
    <div class="container">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import auth from "./auth";

export default {
  data() {
    return {};
  },

  methods: {
    logout() {
      auth.logout(this);
    },
    isAuthenticated() {
      return auth.isAuthenticated();
    }
  }
};
</script>