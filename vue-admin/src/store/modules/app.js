const app = {
  sidemenu:
  [
    {
      icon: 'local_activity',
      title: 'Activities',
      route: 'activities'
    },
    {
      icon: 'work',
      title: 'Briefs',
      route: 'briefs'
    },
    {
      icon: 'shopping_cart',
      title: 'Carts',
      route: 'carts'
    },
    {
      icon: 'face',
      title: 'Clients',
      route: 'clients'
    },
    {
      icon: 'outlined_flag',
      title: 'Countries',
      route: 'countries'
    },
    {
      icon: 'card_giftcard',
      title: 'Coupons',
      route: 'coupons'
    },
    {
      icon: 'money',
      title: 'Currencies',
      route: 'currencies'
    },
    {
      icon: 'payment',
      title: 'Gateways',
      route: 'gateways'
    },
    {
      icon: 'bookmark_border',
      title: 'Locations',
      route: 'locations'
    },
    {
      icon: 'mail_outline',
      title: 'Mails',
      route: 'mails'
    },
    {
      icon: 'description',
      title: 'Orders',
      route: 'orders'
    },
    {
      icon: 'payment',
      title: 'Payments Methods',
      route: 'payments-methods'
    },
    {
      icon: 'bookmark_border',
      title: 'Sectors',
      route: 'sectors'
    },
    {
      icon: 'web_asset',
      title: 'Service Forms',
      route: 'service-forms'
    },
    {
      icon: 'subtitles',
      title: 'Services',
      route: 'services'
    },
    {
      icon: 'account_circle',
      title: 'Users',
      route: 'users'
    }]
}

export default {
  namespaced: true,
  state: {sidemenu: app.sidemenu},
  mutations: { },
  actions: { },
  getters: { }
}
