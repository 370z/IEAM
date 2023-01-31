  export default ({ $axios }) => {
    $axios.onRequest(config => {
      config.headers.common['Content-Type'] = 'application/json';
      config.headers.common['Accept'] = 'application/json';
      config.headers.common['Authorization'] = localStorage.getItem('auth._token.local');
    });
  }