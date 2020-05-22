<template>
 <div class="container">
    <h1>Applications</h1>
    <b-button class="btn btn-success" v-b-modal.application-modal>
      <b-icon icon="plus" aria-hidden="true" ></b-icon> New Application
    </b-button>
  <div class="tab-content custom-scrollbar">
      <table class="table table-hover">
          <thead>
            <tr>
              <th scope="col">ID</th>
              <th scope="col">Title</th>
              <th scope="col">Description</th>
              <th scope="col">Status</th>
              <th scope="col">Executor</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="inf in app" :key="inf.id">
              <td>{{ inf.id }}</td>
              <td>{{  inf.title }}</td>
              <td>{{ inf.description }}</td>
              <td>{{inf.status}}</td>
              <td>{{inf.executor}}</td>
              <td>
                <b-button variant="danger" @click="onDeleteEvent(inf)">
                  <b-icon icon="trash" aria-hidden="true">
                  </b-icon>Delete
                </b-button>
              </td>
            </tr>
          </tbody>
        </table>
    </div>
   <b-modal ref="ApplicationModal"
         id="application-modal"
         title="New Application"
         hide-footer
         hide-backdrop
         centered>
        <b-form @submit="Submit" @reset="Reset" class="w-100">
        <b-form-group id="form-title-group"
                    label="Title:"
                    label-for="form-title-input">
          <b-form-input id="form-title-input"
                        type="text"
                        v-model="Application.title"
                        required
                        placeholder="Enter title">
          </b-form-input>
        </b-form-group>
        <b-form-group id="form-description-group"
                      label="Description:"
                      label-for="form-description-input">
            <b-form-textarea id="form-description-input"
                          v-model="Application.description"
                          required
                          placeholder="Enter description">
            </b-form-textarea>
        </b-form-group>
        <b-button type="submit" variant="primary">Submit</b-button>
        <b-button type="reset" variant="danger">Reset</b-button>
      </b-form>
      </b-modal>
    </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ApplicationsList',
  data() {
    return {
      Application: {
        title: '',
        description: '',
      },
      app: [{
        id: 1,
        title: 'Проблема с подключением',
        description: 'Не могу подключиться к базе данных компании',
        status: 'В процессе',
        executor: 'Иван Иванов',
      }],
    };
  },
  methods: {
    getApplications() {
      const path = 'http://localhost:1984/get/applications';
      axios.get(path, {
        headers: { Authorization: `Bearer ${this.$store.state.token}` },
      }).then((resp) => {
        this.applications = resp.data;
      }).catch(err => console.log(err));
    },
    InitForm() {
      this.Application.title = '';
      this.Application.description = '';
    },
    Reset(evt) {
      evt.preventDefault();
      this.$refs.ApplicationModal.hide();
      this.initForm();
    },
    Submit(evt) {
      evt.preventDefault();
      this.$refs.ApplicationModal.hide();
      const payload = {
        title: this.Application.title,
        description: this.Application.description,
      }
      this.sendApplication(payload);
      this.initForm();
    },
    sendApplication(payload) {
      const path = 'http://localhost:1984/add/applications';
      axios.post(path, payload, {
        headers: { Authorization: `Bearer ${this.$store.state.token}` },
      })
        .then(() => {
          this.getApplications();
        }).catch((err) => {
          console.log(err);
          this.getApplications();
        });
    },
    removeApplication(appID) {
      const path = `http://localhost:1984/delete/application/${appID}`;
      axios.delete(path, {
        headers: { Authorization: `Bearer ${this.$store.state.token}` },
      })
        .then(() => {
          this.getApplications();
        })
        .catch((error) => {
        // eslint-отключение следующей строки
          console.error(error);
          this.getApplications();
        });
    },
    onDeleteEvent(application) {
      this.removeApplication(application.id);
    },
  },
  created() {
    this.InitForm();

  },
};
</script>

<style scoped>

</style>
