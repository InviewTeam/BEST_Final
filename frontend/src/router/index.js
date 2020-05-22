import Vue from 'vue';
import Router from 'vue-router';
import EquipmentsList from '@/components/Equipment/EquipmentsList';
import ApplicationsList from '@/components/Application/ApplicationsList';
import store from '../store';


Vue.use(Router);

const router = new Router({
  routes: [
    {
      name: 'Equipment',
      path: '/equipment',
      component: EquipmentsList,
      meta: {
        requiresAuth: true,
      },
    },
    {
      name: 'Applications',
      path: '/application',
      component: ApplicationsList,
      meta: {
        requiresAuth: true,
      },
    },
  ],
  mode: 'history',
});


router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.isLoggedIn) {
      next();
      return;
    }
    next('/');
  } else {
    next();
  }
});

export default router;
