<template>
  <div>
    <button class="btn btn-primary" v-b-modal.cart-modal>Cart {{ numInCart }}</button>
    <b-modal ref="EquipmentCart"
             id="cart-modal"
             title="Equipment cart" hide-footer
              hide-backdrop
              centered>
      <b-form @submit="Submit" @reset="Reset" class="w-100">
        <table class="table">
          <tbody>
            <tr v-for="(item, index) in cart">
              <td>{{ item.name }}</td>
              <td>{{ item.price }} ₽</td>
              <td>{{ item.count }}</td>
              <td>
                <b-button class="btn btn-sm btn-danger" @click="removeFromCart(index)">&times;</b-button>
              </td>
            </tr>
            <tr>
              <th>Total: </th>
              <th></th>
              <th></th>
              <th>{{ total }}</th>
            </tr>
          </tbody>
        </table>
        <b-button type="reset" variant="secondary"  >Keep choosing</b-button>
        <b-button type="submit" variant="primary">Check out</b-button>
      </b-form>
    </b-modal>
</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'EquipmentCart',
  data() {
    return {
      count: 0,
    };
  },
  computed: {
    numInCart() { return this.$store.getters.getCart.length; },
    cart() {
      return this.$store.getters.getCart.map((cartItem) => {
        const info = this.$store.getters.getEquipments.find(equipmentItem => cartItem.id === equipmentItem.invId);
        const count = this.$store.getters.getCart.find(item => item.id === cartItem.id).count;
        info.count = count;
        return info;
      });
    },
    total() {
      return this.cart.reduce((acc, cur) => acc + (cur.price * cur.count), 0);
    },
  },
  methods: {
    removeFromCart(index) {
      this.$store.dispatch('removeFromCart', index);
    },
    Submit(evt) {
      evt.preventDefault();
      this.$refs.EquipmentCart.hide();
      const payload = this.createPayload();
      this.addRequest(payload);
    },
    createPayload() {
      return this.$store.getters.getCart.map((cartItem) => {
        // eslint-disable-next-line max-len
        const title = this.$store.getters.getEquipments.find(equipmentItem => cartItem.id === equipmentItem.invId).name;
        const count = this.$store.getters.getCart.find(item => item.id === cartItem.id).count;
        const info = {
          title,
          count,
        };
        return info;
      });
    },
    addRequest(payload) {
      const path = '';
      console.log(payload);
      axios.post(path, payload);
      this.$store.dispatch('clean');
    },
    Reset(evt) {
      evt.preventDefault();
      this.$refs.EquipmentCart.hide();
    },
  },
};
</script>
