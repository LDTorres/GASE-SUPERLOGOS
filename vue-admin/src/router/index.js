import Vue from 'vue'
import Router from 'vue-router'

/**
 * Components
 */

import home from '@/components/home'
import orders from '@/components/orders'
import activities from '@/components/activities'
import briefs from '@/components/briefs'
import carts from '@/components/carts'
import clients from '@/components/clients'
import countries from '@/components/countries'
import coupons from '@/components/coupons'
import currencies from '@/components/currencies'
import gateways from '@/components/gateways'
import locations from '@/components/locations'
import mails from '@/components/mails'
import paymentsMethods from '@/components/payments_methods'
import sectors from '@/components/sectors'
import serviceForms from '@/components/service_forms'
import services from '@/components/services'
import users from '@/components/users'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: home
    },
    {
      path: '/orders',
      name: 'orders',
      component: orders
    },
    {
      path: '/activities',
      name: 'activities',
      component: activities
    },
    {
      path: '/briefs',
      name: 'briefs',
      component: briefs
    },
    {
      path: '/carts',
      name: 'carts',
      component: carts
    },
    {
      path: '/clients',
      name: 'clients',
      component: clients
    },
    {
      path: '/countries',
      name: 'countries',
      component: countries
    },
    {
      path: '/coupons',
      name: 'coupons',
      component: coupons
    },
    {
      path: '/currencies',
      name: 'currencies',
      component: currencies
    },
    {
      path: '/gateways',
      name: 'gateways',
      component: gateways
    },
    {
      path: '/locations',
      name: 'locations',
      component: locations
    },
    {
      path: '/mails',
      name: 'mails',
      component: mails
    },
    {
      path: '/payments-methods',
      name: 'payments-methods',
      component: paymentsMethods
    },
    {
      path: '/sectors',
      name: 'sectors',
      component: sectors
    },
    {
      path: '/service-forms',
      name: 'service-forms',
      component: serviceForms
    },
    {
      path: '/services',
      name: 'services',
      component: services
    },
    {
      path: '/users',
      name: 'users',
      component: users
    }
  ]
})
