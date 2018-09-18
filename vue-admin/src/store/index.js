import Vue from 'vue'
import Vuex from 'vuex'
import { mutations, STORAGE_KEY } from './mutations'
import actions from './actions'

/**
 * Modules
 */

import activities from './modules/activities'
import briefs from './modules/briefs'
import carts from './modules/carts'
import clients from './modules/clients'
import countries from './modules/countries'
import coupons from './modules/coupons'
import currencies from './modules/currencies'
import gateways from './modules/gateways'
import locations from './modules/locations'
import mails from './modules/mails'
import orders from './modules/orders'
import paymentsMethods from './modules/payments_methods'
import sectors from './modules/sectors'
import servicesForms from './modules/service_forms'
import services from './modules/services'
import users from './modules/users'

import app from './modules/app'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    todos: JSON.parse(window.localStorage.getItem(STORAGE_KEY) || '[]')
  },
  actions,
  mutations,
  modules: {
    activities: activities,
    briefs: briefs,
    carts: carts,
    clients: clients,
    countries: countries,
    coupons: coupons,
    currencies: currencies,
    gateways: gateways,
    locations: locations,
    mails: mails,
    orders: orders,
    paymentsMethods: paymentsMethods,
    sectors: sectors,
    servicesForms: servicesForms,
    services: services,
    users: users,
    app: app
  }
})
