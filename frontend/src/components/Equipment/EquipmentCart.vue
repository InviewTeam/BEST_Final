<template>
  <div>
    <button class="btn btn-primary" v-b-modal.cart-modal >Cart {{ numInCart }}</button>
    <b-modal ref="EquipmentCart"
             id="cart-modal"
             title="Equipment cart" hide-footer
              hide-backdrop
              centered>
      <b-form @submit="Submit" @reset="Reset" class="w-100">
        <table class="table">
          <tbody>
            <tr v-for="(item,index) in cart">
              <td>{{ item.name }}</td>
              <td>{{ item.price | dollars }}</td>
              <td>{{ item.count }}</td>
              <td>
                <b-button class="btn btn-sm btn-danger" @click="removeFromCart(index)">&times;</b-button>
              </td>
            </tr>
            <tr>
              <th>Total: </th>
              <th>{{ total | dollars}}</th>
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
import { dollars } from '../../filters';

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
      return this.$store.getters.getCart.map(cartItem => this.$store.getters.getEquipments.find(equipmentItem => cartItem === equipmentItem.invId));
    },
    total() {
      return this.cart.reduce((acc, cur) => acc + cur.price, 0);
    },
  },
  methods: {
    removeFromCart(index) {
      this.$store.dispatch('removeFromCart', index);
    },
    Submit(evt) {
      evt.preventDefault();
      this.$refs.EquipmentCart.hide();
      const payload = this.$store.getters.getCart;
      this.addRequest(payload);
    },
    addRequest(payload) {
      const path = '';
      axios.post(path, payload);
    },
    Reset(evt) {
      evt.preventDefault();
      this.$refs.EquipmentCart.hide();
    },
  },
  filters: {
    dollars,
  },
};
</script>
