import axios from 'axios'
import config from './../devConfig.js'

var headers = {}
const token = localStorage.getItem('bazam-token')

var baseURL = process.env.NODE_ENV === 'development' ? config.dev.serverURL : config.prod.serverURL

if (token !== null && token !== undefined && token !== '') {
  headers['Authorization'] = token
}

var instance = axios.create({
  baseURL: baseURL,
  headers: headers
})
export default instance
