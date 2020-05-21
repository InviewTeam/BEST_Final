import Vue from 'vue';
import Router from 'vue-router';
import EquipmentsList from '@/components/Equipment/EquipmentsList';


Vue.use(Router);

export default new Router({
  routes: [
    { name: 'Equipment',
      path: '/equipment',
      component: EquipmentsList,
    },
  ],
  mode: 'history',
});
