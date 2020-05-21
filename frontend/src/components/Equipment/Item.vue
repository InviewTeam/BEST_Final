<template>
<div class="col-md-3">
       <div class="card">
         <img :src="image" :alt="name" class="card-img-top" width="200" height="200">
         <div class="card-body">
            <h4 class="card-title">{{ name }}</h4>
            <div class="card-text">{{description}}</div>
            <div class="card-text">${{ price | dollars }}</div>
            <div class="row justify-content-end">
              <button class="btn btn-primary" @click="addToCart(invId)">Add to cart</button>
            </div>
         </div>
       </div>
</div>
</template>

<script>
export default {
  name: 'Item',
  props: ['invId', 'image', 'name', 'description', 'price'],
  filters: {
    dollars: num => `${num / 100}`,
  },
  methods: {
    addToCart(invId) {
      this.$store.dispatch('addToCart', invId);
    },
  },
};
</script>

<style scoped>

</style>
