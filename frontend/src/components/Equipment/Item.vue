<template>
<div class="col-md-3">
       <div class="card">
         <img :src="image" :alt="name" class="card-img-top" width="200" height="200">
         <div class="card-body">
            <h4 class="card-title">{{ name }}</h4>
            <div class="card-text">{{description}}</div>
            <div class="card-text">Price: {{ price }} ₽</div>
           <div class="card-text">Quantity: {{count}}</div>
           <div>
               <p>
                <button class="btn btn-success" @click="increment">+</button>
                <button class="btn btn-danger" @click="decrement">-</button>
               </p>
           </div>
           <div>
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
  data() {
    return {
      count: 0,
    };
  },
  methods: {
    addToCart(invId) {
      this.$store.dispatch('addToCart', { id: invId, count: this.count });
      this.count = 0;
    },
    increment() {
      this.count += 1;
    },
    decrement() {
      if (this.count > 0) {
        this.count -= 1;
      }
    },
  },
};
</script>

<style scoped>

</style>
