<template>
  <div id="app">
    <div class="container my-5">
      <div class="row mb-3">
        <div class="col-md-9">
          <h1>Choose equipment</h1>
        </div>
        <div class="col-md-3">
          <EquipmentCart/>
        </div>
      </div>
    <div class="row">
      <Item
        v-for="item in getEquipments"
        :key="item.invId"
        :invId="item.invId"
        :name="item.name"
        :image="item.image"
        :description="item.description"
        :price="item.price" />
      </div>
    </div>
  </div>
</template>

<script>
import Item from './Item';
import EquipmentCart from './EquipmentCart';

export default {
  name: 'EquipmentsList',
  computed: {
    getEquipments() { return this.$store.getters.getEquipments; },
  },
  components: {
    Item,
    EquipmentCart,
  },

};
</script>

<style scoped>

</style>
