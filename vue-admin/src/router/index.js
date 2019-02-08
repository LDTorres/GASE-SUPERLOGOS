import Vue from 'vue'
import Router from 'vue-router'

/**
 * Components
 */

import analytics from '@/components/analytics'
import login from '@/components/login'
import orders from '@/components/orders'
import activities from '@/components/activities'
import briefs from '@/components/briefs'
import clients from '@/components/clients'
import countries from '@/components/countries'
import coupons from '@/components/coupons'
import currencies from '@/components/currencies'
import gateways from '@/components/gateways'
import locations from '@/components/locations'
import sectors from '@/components/sectors'
import services from '@/components/services'
import users from '@/components/users'
import portfolios from '@/components/portfolios'
import trashed from '@/components/trashed'
import projects from '@/components/projects'
import leads from '@/components/leads'
import safetypay from '@/components/safetypay'

Vue.use(Router)

var instance = new Router({
  routes: [
    {
      path: '/analytics',
      name: 'analytics',
      component: analytics
    },
    {
      path: '/',
      name: 'login',
      component: login
    },
    {
      path: '/orders',
      name: 'orders',
      component: orders,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/projects',
      name: 'projects',
      component: projects,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/activities',
      name: 'activities',
      component: activities,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/briefs',
      name: 'briefs',
      component: briefs,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/clients',
      name: 'clients',
      component: clients,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/countries',
      name: 'countries',
      component: countries,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/coupons',
      name: 'coupons',
      component: coupons,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/currencies',
      name: 'currencies',
      component: currencies,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/gateways',
      name: 'gateways',
      component: gateways,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/locations',
      name: 'locations',
      component: locations,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/sectors',
      name: 'sectors',
      component: sectors,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/services',
      name: 'services',
      component: services,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/users',
      name: 'users',
      component: users,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/safetypay',
      name: 'safetypay',
      component: safetypay,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/leads',
      name: 'leads',
      component: leads,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/portfolios',
      name: 'portfolios',
      component: portfolios,
      props: (route) => ({ search: route.query.q })
    },
    {
      path: '/trashed',
      name: 'trashed',
      component: trashed,
      props: (route) => ({ model: route.query.m, search: route.query.q })
    },
    {
      path: '*',
      redirect: '/orders'
    }
  ]
})

instance.beforeEach((to, from, next) => {
  let token = localStorage.getItem('bazam-token')
  if (token) {
    this.name = to.params.name
    next()
  } else {
    next()
  }
})

export default instance
