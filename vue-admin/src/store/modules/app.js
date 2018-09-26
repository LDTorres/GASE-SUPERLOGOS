const app = {
  sidemenu:
  [
    {
      title: 'Ubicaciones',
      routes: [{
        icon: 'outlined_flag',
        title: 'Paises',
        route: 'countries'
      },
      {
        icon: 'location_on',
        title: 'Locaciones',
        route: 'locations'
      },
      {
        icon: 'bookmark_border',
        title: 'Sectores',
        route: 'sectors'
      },
      {
        icon: 'local_activity',
        title: 'Actividades',
        route: 'activities'
      }]
    },
    {
      title: 'Servicios',
      routes: [{
        icon: 'subtitles',
        title: 'Servicios',
        route: 'services'
      },
      {
        icon: 'store',
        title: 'Ordenes',
        route: 'orders'
      },
      {
        icon: 'money',
        title: 'Monedas',
        route: 'currencies'
      },
      {
        icon: 'payment',
        title: 'Pasarelas',
        route: 'gateways'
      },
      {
        icon: 'card_giftcard',
        title: 'Cupones',
        route: 'coupons'
      }]
    },
    {
      title: 'Clientes',
      routes: [{
        icon: 'face',
        title: 'Clientes',
        route: 'clients'
      },
      {
        icon: 'assignment',
        title: 'Briefs',
        route: 'briefs'
      },
      {
        icon: 'work',
        title: 'Portafolios',
        route: 'portfolios'
      }]
    },
    {
      title: 'Administradores',
      routes: [{
        icon: 'security',
        title: 'Usuarios',
        route: 'users'
      }]
    }
    /* ,{
      icon: 'mail_outline',
      title: 'Plantillas de Emails',
      route: 'mails'
    } */
  ]
}

export default {
  namespaced: true,
  state: {sidemenu: app.sidemenu},
  mutations: { },
  actions: { },
  getters: { }
}
